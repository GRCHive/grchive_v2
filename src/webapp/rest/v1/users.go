package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (w *WebappApplication) apiv1GetCurrentUser(c *gin.Context) {
	session := w.sessionStore.GetLoginSession(c).GetServerSession()
	user, err := w.backend.userMgr.GetUserFromId(session.UserId)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, user)
}
