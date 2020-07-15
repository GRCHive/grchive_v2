package session

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/sessions"
	"github.com/jmoiron/sqlx"
	"github.com/lestrrat-go/jwx/jwt"
	"gitlab.com/grchive/grchive-v2/shared/backend"
	gsess "gitlab.com/grchive/grchive-v2/shared/backend/sessions"
	"gitlab.com/grchive/grchive-v2/shared/backend/users"
	"gitlab.com/grchive/grchive-v2/shared/backend/utility"
	"gitlab.com/grchive/grchive-v2/shared/fusionauth"
	"net/http"
	"time"
)

const loginSessionName = "login-session"
const loginStateName = "login-state"
const sessionIdName = "session-id"

var sessionDuration = time.Minute * time.Duration(30)

var ForceOauthLogoutErr = errors.New("Force OAuth Logout")

type LoginSession struct {
	s   *sessions.Session
	w   http.ResponseWriter
	r   *http.Request
	itf *backend.BackendInterface

	session     *gsess.Session
	sessionUser *users.User
}

func (l *LoginSession) GetSessionUser() *users.User {
	return l.sessionUser
}

func (l *LoginSession) Logout(sess *gsess.Session, fa *fusionauth.FusionAuthClient) error {
	// We need to do two things here to force a relogin.
	// 1) Logout the user on the FusionAuth side so that they'll be forced to see the login screen and type in their credentials.
	// 2) Delete our copy of the session in the database. This will trigger the "sess == nil" check on the next go around.
	resp, err := fa.Logout(true, sess.RefreshToken)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New("Failed to force logout.")
	}

	return l.itf.Sessions.WrapDatabaseTx(func(tx *sqlx.Tx) error {
		return l.itf.Sessions.DeleteSession(tx, sess.Id)
	})
}

func (l *LoginSession) LogoutCurrentSession(fa *fusionauth.FusionAuthClient) error {
	return l.Logout(l.session, fa)
}

func (l *LoginSession) ValidateLogin(fa *fusionauth.FusionAuthClient) error {
	sessionId, ok := l.s.Values[sessionIdName]
	if !ok {
		return nil
	}

	sess, err := l.itf.Sessions.GetSessionFromId(sessionId.(string))
	if err != nil {
		return err
	}

	if sess == nil {
		return nil
	}

	tm := time.Now()
	// First, check if the session itself is expired. If that's the case, force the user to re-login - this is
	// not a valid session!
	if tm.After(sess.SessionExpiration) {
		err = l.Logout(sess, fa)
		if err != nil {
			return err
		}
		return ForceOauthLogoutErr
	}

	// Next check if the access token is expired - in that case, use the refresh token
	// to get a new access token. Maintain the same session.
	if tm.After(sess.JwtExpiration) {
		token, err := fa.ExchangeRefreshTokenForAccessToken(sess.RefreshToken)
		if err != nil {
			return err
		}

		validatedToken, err := utility.ValidateJWT(token.Token(), fa)
		if err != nil {
			return err
		}

		sess.JwtExpiration = validatedToken.Expiration().UTC()
		sess.AccessToken = token.Token()
		sess.RefreshToken = token.RefreshToken
	}

	// Bump up the expiration time since the user did something.
	sess.SessionExpiration = time.Now().Add(sessionDuration).UTC()
	err = l.itf.Sessions.WrapDatabaseTx(func(tx *sqlx.Tx) error {
		return l.itf.Sessions.UpdateSession(tx, sess)
	})

	l.session = sess
	l.sessionUser, err = l.itf.Users.GetUserFromId(l.session.UserId)
	return err
}

func (l *LoginSession) GetServerSession() *gsess.Session {
	return l.session
}

func (l *LoginSession) IsLoggedIn() bool {
	return l.session != nil
}

func (l *LoginSession) PopulateLoginState() error {
	// This is the state to pass to FusionAuth use for any login operations - NOT THE CSRF
	// since we don't want to expose this to the client.
	l.s.Values[loginStateName] = uuid.New().String()

	err := l.s.Save(l.r, l.w)
	if err != nil {
		return err
	}

	return nil
}

func (l *LoginSession) GetLoginState() string {
	return l.s.Values[loginStateName].(string)
}

func (l *LoginSession) CreateUserSession(user *users.User, parsedToken jwt.Token, token *fusionauth.AccessToken) error {
	newSession := gsess.Session{
		Id:                uuid.New().String(),
		SessionExpiration: time.Now().Add(sessionDuration).UTC(),
		JwtExpiration:     parsedToken.Expiration().UTC(),
		AccessToken:       token.Token(),
		RefreshToken:      token.RefreshToken,
		UserId:            user.Id,
	}

	err := l.itf.Sessions.WrapDatabaseTx(func(tx *sqlx.Tx) error {
		return l.itf.Sessions.CreateSession(tx, &newSession)
	})

	if err != nil {
		return err
	}

	l.s.Values[sessionIdName] = newSession.Id
	err = l.s.Save(l.r, l.w)
	if err != nil {
		return err
	}

	return nil
}

func (s *SessionStore) GetLoginSession(c *gin.Context) *LoginSession {
	// First check if we exist in the request's context.
	// Only if we don't already exist do we parse it from the session store - this is because we want
	// to reuse the same LoginSession for the same request to keep track of state.
	const contextKey = "loginContextKey"
	ls, ok := c.Get(contextKey)
	if ok {
		return ls.(*LoginSession)
	}

	// We can ignore the error as the error would just tell us if the session couldn't be decoded.
	// That's fine - just create a new one and cause a little inconvenience.
	sess, _ := s.store.Get(c.Request, loginSessionName)
	if sess != nil {
		sess.Options = &sessions.Options{
			MaxAge:   0,
			HttpOnly: true,
			Secure:   s.SecureCookies,
		}
	}

	ret := &LoginSession{
		s:   sess,
		w:   c.Writer,
		r:   c.Request,
		itf: s.itf,
	}

	c.Set(contextKey, ret)
	return ret
}

func (s *SessionStore) PopulateLoginState(c *gin.Context) {
	sess := s.GetLoginSession(c)
	err := sess.PopulateLoginState()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	} else {
		c.Next()
	}
}

func (s *SessionStore) RedirectIfLoggedIn(to string) gin.HandlerFunc {
	return func(c *gin.Context) {
		sess := s.GetLoginSession(c)
		if sess.IsLoggedIn() {
			c.Redirect(http.StatusTemporaryRedirect, to)
		} else {
			c.Next()
		}
	}
}

func (s *SessionStore) RedirectIfLoggedOut(to string) gin.HandlerFunc {
	return func(c *gin.Context) {
		sess := s.GetLoginSession(c)
		if !sess.IsLoggedIn() {
			c.Redirect(http.StatusTemporaryRedirect, to)
		} else {
			c.Next()
		}
	}
}

func (s *SessionStore) ValidateLogin(fa *fusionauth.FusionAuthClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		sess := s.GetLoginSession(c)
		err := sess.ValidateLogin(fa)
		if err != nil {
			if err == ForceOauthLogoutErr {
				c.Redirect(http.StatusTemporaryRedirect, fa.GetOAuthLogoutUrl())
			} else {
				c.AbortWithError(http.StatusInternalServerError, err)
			}
		} else {
			c.Next()
		}
	}
}

// We need to figure out if the user has verified their email so we keep track of this information
// here. If we detect that (our) user hasn't verified their email, then send out a request to FusionAuth
// to update.
func (s *SessionStore) SyncEmailVerification(fa *fusionauth.FusionAuthClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		sess := s.GetLoginSession(c)
		if sess.sessionUser != nil && !sess.sessionUser.EmailVerified {
			resp, faErrs, err := fa.RetrieveUser(sess.sessionUser.FusionAuthUserId)
			// Silently fail if we can't get the user -- assume that that's OK.
			if err == nil && faErrs == nil && resp.User.SecureIdentity.Verified {
				s.itf.Users.WrapDatabaseTx(func(tx *sqlx.Tx) error {
					return s.itf.Users.MarkUserVerified(tx, sess.sessionUser.Id)
				})
			}
		}
		c.Next()
	}
}
