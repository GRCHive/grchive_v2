package session

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/sessions"
	"github.com/lestrrat-go/jwx/jwt"
	"gitlab.com/grchive/grchive-v2/shared/fusionauth"
	"net/http"
	"time"
)

const loginSessionName = "login-session"
const loginStateName = "login-state"
const accessTokenName = "accesstoken"
const refreshTokenName = "refreshtoken"
const userIdName = "userid"
const expiresName = "expires"
const timeStorageFormat = time.RFC3339

var noAccessTokenErr = errors.New("No access token found.")

type LoginSession struct {
	s *sessions.Session
	w http.ResponseWriter
	r *http.Request

	token        jwt.Token
	jwtValidated bool
}

func (l *LoginSession) ValidateLogin(fa *fusionauth.FusionAuthClient) error {
	rawToken, ok := l.s.Values[accessTokenName]
	if !ok {
		return noAccessTokenErr
	}

	// Don't need to send the token to FusionAuth because we should
	// be OK since we're storing it in a secure cookie so we should be able
	// to detect when the session gets modified.
	var err error
	l.token, err = jwt.ParseString(rawToken.(string))
	if err != nil {
		return err
	}

	err = jwt.Verify(l.token,
		jwt.WithAcceptableSkew(time.Minute*time.Duration(1)),
		jwt.WithIssuer("grchive.com"),
	)

	l.jwtValidated = (err == nil)
	return nil
}

func (l *LoginSession) IsLoggedIn() bool {
	return l.jwtValidated
}

func (l *LoginSession) PopulateLoginState() error {
	// This is the state to use for any login operations - NOT THE CSRF
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

func (l *LoginSession) StoreAccessToken(token *fusionauth.AccessToken) error {
	l.s.Values[accessTokenName] = token.Token()
	l.s.Values[refreshTokenName] = token.RefreshToken
	l.s.Values[expiresName] = time.Now().UTC().Add(time.Second * time.Duration(token.ExpiresIn)).Format(timeStorageFormat)
	l.s.Values[userIdName] = token.UserId

	err := l.s.Save(l.r, l.w)
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
	ret := &LoginSession{
		s: sess,
		w: c.Writer,
		r: c.Request,
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
		if err != nil && err != noAccessTokenErr {
			c.AbortWithError(http.StatusInternalServerError, err)
		} else {
			c.Next()
		}
	}
}
