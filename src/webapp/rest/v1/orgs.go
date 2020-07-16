package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"gitlab.com/grchive/grchive-v2/shared/backend"
	"gitlab.com/grchive/grchive-v2/shared/backend/orgs"
	"gitlab.com/grchive/grchive-v2/shared/gin_middleware/gin_backend_utility"
	"net/http"
)

func (w *WebappApplication) apiv1CreateNewOrg(c *gin.Context) {
	org := orgs.Organization{}
	err := c.BindJSON(&org)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, &gin_backend_utility.WebappError{
			Err:     err,
			Context: "Obtain org in request.",
		})
		return
	}

	sess := w.sessionStore.GetLoginSession(c)
	currentUser := sess.GetSessionUser()
	// This endpoint shouldn't allow creating a sub-org as this won't have
	// the proper ACL checks in place.
	org.ParentOrgId = nil
	org.OwnerUserId = currentUser.Id

	err = w.backend.itf.WrapDatabaseTx(sess.GetAuditTrailId(c), func(tx *sqlx.Tx) error {
		return w.backend.itf.Orgs.CreateOrg(tx, &org)
	})

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, &gin_backend_utility.WebappError{
			Err:     err,
			Context: "Create org.",
		})
		return
	}

	c.JSON(http.StatusOK, org)
}

func (w *WebappApplication) apiv1GetOrg(c *gin.Context) {
	org, err := w.middleware.GetResourceFromContext(c, backend.RIOrganization)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, &gin_backend_utility.WebappError{
			Err:     err,
			Context: "Obtain org in context",
		})
		return
	}
	c.JSON(http.StatusOK, org.(*orgs.Organization))
}
