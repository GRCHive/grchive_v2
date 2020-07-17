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

func (w *WebappApplication) apiv1CreateSuborg(c *gin.Context) {
	org := orgs.Organization{}
	err := c.BindJSON(&org)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, &gin_backend_utility.WebappError{
			Err:     err,
			Context: "apiv1CreateSuborg - Obtain org in request.",
		})
		return
	}

	currentOrg, err := w.middleware.GetResourceFromContext(c, backend.RIOrganization)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, &gin_backend_utility.WebappError{
			Err:     err,
			Context: "apiv1CreateSuborg - Obtain org in context",
		})
		return
	}

	sess := w.sessionStore.GetLoginSession(c)
	currentUser := sess.GetSessionUser()
	org.ParentOrgId = &currentOrg.(*orgs.Organization).Id
	org.OwnerUserId = currentUser.Id

	err = w.backend.itf.WrapDatabaseTx(sess.GetAuditTrailId(c), func(tx *sqlx.Tx) error {
		return w.backend.itf.Orgs.CreateOrg(tx, &org)
	})

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, &gin_backend_utility.WebappError{
			Err:     err,
			Context: "apiv1CreateSuborg - Create suborg.",
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

func (w *WebappApplication) apiv1UpdateOrg(c *gin.Context) {
	org, err := w.middleware.GetResourceFromContext(c, backend.RIOrganization)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, &gin_backend_utility.WebappError{
			Err:     err,
			Context: "apiv1UpdateOrg - Obtain org in context",
		})
		return
	}

	editOrg := orgs.Organization{}
	err = c.BindJSON(&editOrg)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, &gin_backend_utility.WebappError{
			Err:     err,
			Context: "apiv1UpdateOrg - Obtain org in request.",
		})
		return
	}

	// Copy what can be edited into the session org since we know for sure that
	// the user has access to that.
	// TODO: CU #9phgnd
	// Determine how to securely perform changes to ParentOrgId and OwnerUserId with proper ACL checks.
	torg := org.(*orgs.Organization)
	torg.Name = editOrg.Name
	torg.Description = editOrg.Description

	sess := w.sessionStore.GetLoginSession(c)
	err = w.backend.itf.WrapDatabaseTx(sess.GetAuditTrailId(c), func(tx *sqlx.Tx) error {
		return w.backend.itf.Orgs.UpdateOrg(tx, torg)
	})

	c.JSON(http.StatusOK, torg)
}

func (w *WebappApplication) apiv1GetSuborgs(c *gin.Context) {
	org, err := w.middleware.GetResourceFromContext(c, backend.RIOrganization)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, &gin_backend_utility.WebappError{
			Err:     err,
			Context: "apiv1GetSuborg - Obtain org in context",
		})
		return
	}

	suborgs, err := w.backend.itf.Orgs.GetSuborgsFromId(org.(*orgs.Organization).Id)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, &gin_backend_utility.WebappError{
			Err:     err,
			Context: "apiv1GetSuborg - Get suborgs",
		})
		return
	}

	c.JSON(http.StatusOK, suborgs)
}

func (w *WebappApplication) apiv1GetParentOrgs(c *gin.Context) {
	org, err := w.middleware.GetResourceFromContext(c, backend.RIOrganization)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, &gin_backend_utility.WebappError{
			Err:     err,
			Context: "apiv1GetParentOrgs - Obtain org in context",
		})
		return
	}

	parents, err := w.backend.itf.Orgs.GetParentOrgsFromId(org.(*orgs.Organization).Id)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, &gin_backend_utility.WebappError{
			Err:     err,
			Context: "apiv1GetParentOrgs - Get parent orgs",
		})
		return
	}

	c.JSON(http.StatusOK, parents)
}
