package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (w *WebappApplication) apiv1GetCurrentUser(c *gin.Context) {
	session := w.sessionStore.GetLoginSession(c)
	c.JSON(http.StatusOK, session.GetSessionUser())
}

func (w *WebappApplication) apiv1ResendEmailVerification(c *gin.Context) {
	user := w.sessionStore.GetLoginSession(c).GetSessionUser()
	_, emailErr, err := w.fusionauth.ResendEmailVerification(user.Email)
	if emailErr != nil || err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	} else {
		c.Status(http.StatusNoContent)
	}
}
