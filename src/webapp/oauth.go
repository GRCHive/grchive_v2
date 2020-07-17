package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"gitlab.com/grchive/grchive-v2/shared/backend/users"
	"gitlab.com/grchive/grchive-v2/shared/backend/utility"
	"gitlab.com/grchive/grchive-v2/shared/gin_middleware/gin_backend_utility"
	"net/http"
)

func (w *WebappApplication) renderLogin(c *gin.Context) {
	session := w.sessionStore.GetLoginSession(c)

	c.Redirect(
		http.StatusTemporaryRedirect,
		w.fusionauth.GetOauthLoginEndpoint(session.GetLoginState()),
	)
}

func (w *WebappApplication) renderLogout(c *gin.Context) {
	session := w.sessionStore.GetLoginSession(c)
	if !session.IsLoggedIn() {
		c.Redirect(
			http.StatusTemporaryRedirect,
			"/",
		)
		return
	}

	err := session.LogoutCurrentSession(w.fusionauth)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, &gin_backend_utility.WebappError{
			Err:     err,
			Context: "Logout user.",
		})
		return
	}

	c.Redirect(
		http.StatusTemporaryRedirect,
		w.fusionauth.GetOAuthLogoutUrl(),
	)
}

func (w *WebappApplication) handleOauthCallback(c *gin.Context) {
	// This function is called when FusionAuth redirects a successful login back to us in the
	// OAuth authentication code flow. Verify the state and use the code to get a token for the user
	// and create an actually logged in session. Any errors here should just redirect user back to
	// the login page.
	onError := func(err error, msg string) {
		if err != nil {
			w.log.Error().Err(err).Msg(msg)
		}
		c.AbortWithError(http.StatusUnauthorized, &gin_backend_utility.WebappError{
			Err:     err,
			Context: msg,
		})
	}

	session := w.sessionStore.GetLoginSession(c)
	if session.GetLoginState() != c.Query("state") {
		onError(nil, "State mismatch.")
		return
	}

	token, err := w.fusionauth.ExchangeOAuthCodeForAccessToken(c.Query("code"))
	if err != nil {
		onError(err, "Failed to retrieve access token.")
		return
	}

	validatedToken, err := utility.ValidateJWT(token.Token(), w.fusionauth)
	if err != nil {
		onError(err, "Failed to validate token.")
		return
	}

	// Now we want to transform this JWT token into a login session.
	// We don't really have a need to use the JWT as we can get the same
	// benefits of having a secure session with less data by using a session ID.
	// First check if we need to track the user in our database as this could be the first time
	// the user is logging in with us (e.g. via SAML).
	user, err := w.backend.itf.Users.GetUserFromFusionAuthId(validatedToken.Subject())
	if err != nil {
		onError(err, "Failed to obtain user from FusionAuth ID.")
		return
	}

	if user == nil {
		user = &users.User{
			FullName:         validatedToken.PrivateClaims()["fullName"].(string),
			Email:            validatedToken.PrivateClaims()["email"].(string),
			FusionAuthUserId: validatedToken.Subject(),
		}

		err = w.backend.itf.WrapDatabaseTx(w.middleware.GetAuditTrailIdWithUserOverride(user, c), func(tx *sqlx.Tx) error {
			return w.backend.itf.Users.CreateUser(tx, user)
		})

		if err != nil {
			onError(err, "Failed to create user.")
			return
		}
	}

	err = session.CreateUserSession(
		user,
		validatedToken,
		token,
	)
	if err != nil {
		onError(err, "Failed to store access token.")
		return
	}

	c.Redirect(http.StatusTemporaryRedirect, "/")
}
