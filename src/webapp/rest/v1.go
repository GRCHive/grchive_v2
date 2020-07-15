package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (w *WebappApplication) apiv1EnsureAuth(c *gin.Context) {
	// For now just ensure we have a valid session.
	// TODO: #9jjxxu
	// Support API keys for non browser users.
	if !w.sessionStore.GetLoginSession(c).IsLoggedIn() {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	c.Next()
}

func (w *WebappApplication) registerApiv1(r *gin.Engine) {
	apiR := r.Group("/api/v1", w.apiv1EnsureAuth)
	{
		userR := apiR.Group("/users")
		{
			currentR := userR.Group("/current")
			currentR.GET("/", w.apiv1GetCurrentUser)
			currentR.PUT("/", w.apiv1UpdateCurrentUser)
			currentR.GET("/orgs", w.apiv1GetCurrentUserOrgs)
			currentR.POST("/verify", w.apiv1ResendEmailVerification)
		}

		orgsR := apiR.Group("/orgs")
		{
			orgsR.POST("/", w.apiv1CreateNewOrg)
		}
	}
}
