package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (w *WebappApplication) renderLogin(c *gin.Context) {
	session := w.sessionStore.GetLoginSession(c)

	c.Redirect(
		http.StatusTemporaryRedirect,
		w.fusionauth.GetOauthLoginEndpoint(session.GetLoginState()),
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
		c.Redirect(http.StatusTemporaryRedirect, "/login")
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

	err = session.StoreAccessToken(token)
	if err != nil {
		onError(err, "Failed to store access token.")
		return
	}

	c.Redirect(http.StatusTemporaryRedirect, "/")
}
