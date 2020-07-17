package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"gitlab.com/grchive/grchive-v2/shared/backend/roles"
	"gitlab.com/grchive/grchive-v2/shared/backend/users"
	"gitlab.com/grchive/grchive-v2/shared/fusionauth"
	"gitlab.com/grchive/grchive-v2/shared/gin_middleware/gin_backend_utility"
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
		c.AbortWithError(http.StatusInternalServerError, &gin_backend_utility.WebappError{
			Err:     err,
			Context: "Get user orgs.",
		})
		return
	}
	c.JSON(http.StatusOK, orgs)
}

func (w *WebappApplication) apiv1UpdateCurrentUser(c *gin.Context) {
	user := users.User{}
	err := c.BindJSON(&user)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, &gin_backend_utility.WebappError{
			Err:     err,
			Context: "Read user from JSON body.",
		})
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

	err = w.backend.itf.WrapDatabaseTx(w.middleware.GetAuditTrailId(c), func(*sqlx.Tx) error {
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
		c.AbortWithError(http.StatusInternalServerError, &gin_backend_utility.WebappError{
			Err:     err,
			Context: "Update user.",
		})
		return
	}

	c.Status(http.StatusNoContent)
}

func (w *WebappApplication) apiv1ResendEmailVerification(c *gin.Context) {
	user := w.sessionStore.GetLoginSession(c).GetSessionUser()
	_, emailErr, err := w.fusionauth.ResendEmailVerification(user.Email)
	if emailErr != nil || err != nil {
		c.AbortWithError(http.StatusInternalServerError, &gin_backend_utility.WebappError{
			Err:     err,
			Context: "Resend email verification.",
		})
	} else {
		c.Status(http.StatusNoContent)
	}
}

func (w *WebappApplication) apiv1CheckCurrentUserPermissions(c *gin.Context) {
	hasPermission := func() {
		c.JSON(http.StatusOK, true)
	}

	noPermission := func() {
		c.JSON(http.StatusOK, false)
	}

	permissions, ok := c.GetQueryArray("permissions")
	if !ok {
		noPermission()
		return
	}

	w.acl.ACLCheckPermissionHandler(c, hasPermission, func(err error) {
		noPermission()
	}, roles.PermissionArrayFromStrings(permissions...)...)
}
