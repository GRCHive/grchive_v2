package main

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/grchive/grchive-v2/shared/backend"
	"gitlab.com/grchive/grchive-v2/shared/backend/roles"
	"net/http"
)

func (w *WebappApplication) apiv1EnsureAuth(c *gin.Context) {
	// For now just ensure we have a valid session.
	// TODO: #9jjxxu
	// Support API keys for non browser users.
	sess := w.sessionStore.GetLoginSession(c)
	if !sess.IsLoggedIn() {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	err := w.middleware.StoreCurrentUserIntoContext(c, sess.GetSessionUser())
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
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

		orgsR := apiR.Group("/orgs", w.acl.ACLUserEmailVerified)
		{
			orgsR.POST("/", w.apiv1CreateNewOrg)

			singleOrgR := orgsR.Group("/:orgId",
				w.middleware.LoadResourceIntoContext(backend.RIOrganization, "orgId"),
				w.middleware.LoadCurrentUserRolesIntoContext,
				w.acl.ACLUserHasValidRole,
			)
			{
				singleOrgR.GET("/",
					w.acl.ACLUserHasPermissions(roles.POrgProfileView),
					w.apiv1GetOrg)
			}
		}
	}
}
