package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"gitlab.com/grchive/grchive-v2/shared/backend/users"
	"gitlab.com/grchive/grchive-v2/shared/fusionauth"
	"net/http"
)

func (w *WebappApplication) apiv1GetCurrentUser(c *gin.Context) {
	session := w.sessionStore.GetLoginSession(c)
	c.JSON(http.StatusOK, session.GetSessionUser())
}

func (w *WebappApplication) apiv1GetCurrentUserOrgs(c *gin.Context) {
	currentUser := w.sessionStore.GetLoginSession(c).GetSessionUser()
	orgs, err := w.backend.itf.GetUserOrgs(currentUser.Id)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, orgs)
}

func (w *WebappApplication) apiv1UpdateCurrentUser(c *gin.Context) {
	user := users.User{}
	err := c.BindJSON(&user)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// Force set current user ID on the bound user just in case
	// the input is malicious and tries to update information about
	// another user.
	sess := w.sessionStore.GetLoginSession(c)
	currentUser := sess.GetSessionUser()

	user.Id = currentUser.Id
	user.Email = currentUser.Email
	user.FusionAuthUserId = currentUser.FusionAuthUserId

	err = w.backend.itf.WrapDatabaseTx(sess.GetAuditTrailId(c), func(*sqlx.Tx) error {
		// Probably good idea to keep information about the user in sync
		// with FusionAuth.
		resp, _, err := w.fusionauth.UpdateUser(currentUser.FusionAuthUserId, fusionauth.UserRequest{
			SendSetPasswordEmail: false,
			SkipVerification:     true,
			User: fusionauth.User{
				Email:    user.Email,
				FullName: user.FullName,
			},
		})

		if err != nil {
			return err
		}

		if resp.StatusCode != http.StatusOK {
			return errors.New("Failed to update user with FA.")
		}

		return nil
	}, func(tx *sqlx.Tx) error {
		return w.backend.itf.Users.UpdateUser(tx, &user)
	})

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusNoContent)
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
