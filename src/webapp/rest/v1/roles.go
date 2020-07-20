package main

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/grchive/grchive-v2/shared/backend"
	"gitlab.com/grchive/grchive-v2/shared/backend/orgs"
	"gitlab.com/grchive/grchive-v2/shared/gin_middleware/gin_backend_utility"
	"net/http"
)

func (w *WebappApplication) apiv1ListAllOrgRoles(c *gin.Context) {
	org, err := w.middleware.GetResourceFromContext(c, backend.RIOrganization)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, &gin_backend_utility.WebappError{
			Err:     err,
			Context: "apiv1ListAllOrgRoles - Obtain org in context",
		})
		return
	}

	roles, err := w.backend.itf.Roles.ListRolesForOrg(org.(*orgs.Organization).Id)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, &gin_backend_utility.WebappError{
			Err:     err,
			Context: "apiv1ListAllOrgRoles - Get roles",
		})
		return
	}

	c.JSON(http.StatusOK, roles)
}

func (w *WebappApplication) apiv1ListAllOrgRolesForCurrentUser(c *gin.Context) {
	org, err := w.middleware.GetResourceFromContext(c, backend.RIOrganization)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, &gin_backend_utility.WebappError{
			Err:     err,
			Context: "apiv1ListAllOrgRolesForCurrentUser - Obtain org in context",
		})
		return
	}

	user, err := w.middleware.GetCurrentUserFromContext(c)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, &gin_backend_utility.WebappError{
			Err:     err,
			Context: "apiv1ListAllOrgRolesForCurrentUser - Obtain current user in context",
		})
		return
	}

	roles, err := w.backend.itf.Roles.ListRolesForOrgAndUser(org.(*orgs.Organization).Id, user.Id)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, &gin_backend_utility.WebappError{
			Err:     err,
			Context: "apiv1ListAllOrgRolesForCurrentUser - Get roles",
		})
		return
	}

	c.JSON(http.StatusOK, roles)
}
