package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (w *WebappApplication) apiv1GetCurrentUser(c *gin.Context) {
	session := w.sessionStore.GetLoginSession(c)
	c.JSON(http.StatusOK, session.GetSessionUser())
}
