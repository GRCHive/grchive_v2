package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"gitlab.com/grchive/grchive-v2/shared/backend"
	"gitlab.com/grchive/grchive-v2/shared/backend/engagements"
	"gitlab.com/grchive/grchive-v2/shared/backend/orgs"
	"gitlab.com/grchive/grchive-v2/shared/gin_middleware/gin_backend_utility"
	"net/http"
)

func (w *WebappApplication) apiv1ListAllOrgEngagements(c *gin.Context) {
	org, err := w.middleware.GetResourceFromContext(c, backend.RIOrganization)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, &gin_backend_utility.WebappError{
			Err:     err,
			Context: "apiv1ListAllOrgEngagements - Obtain org in context",
		})
		return
	}

	engs, err := w.backend.itf.Engagements.ListEngagementsForOrgId(org.(*orgs.Organization).Id)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, &gin_backend_utility.WebappError{
			Err:     err,
			Context: "apiv1ListAllOrgEngagements - List engagements",
		})
		return
	}

	c.JSON(http.StatusOK, engs)
}

func (w *WebappApplication) apiv1ListAllOrgEngagementsForCurrentUser(c *gin.Context) {
	org, err := w.middleware.GetResourceFromContext(c, backend.RIOrganization)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, &gin_backend_utility.WebappError{
			Err:     err,
			Context: "apiv1ListAllOrgEngagementsForCurrentUser - Obtain org in context",
		})
		return
	}

	user, err := w.middleware.GetCurrentUserFromContext(c)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, &gin_backend_utility.WebappError{
			Err:     err,
			Context: "apiv1ListAllOrgEngagementsForCurrentUser - Obtain current user in context",
		})
		return
	}

	engs, err := w.backend.itf.Engagements.ListEngagementsForOrgIdAndUserId(org.(*orgs.Organization).Id, user.Id)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, &gin_backend_utility.WebappError{
			Err:     err,
			Context: "apiv1ListAllOrgEngagementsForCurrentUser - List engagements",
		})
		return
	}

	c.JSON(http.StatusOK, engs)
}

func (w *WebappApplication) apiv1CreateEngagement(c *gin.Context) {
	org, err := w.middleware.GetResourceFromContext(c, backend.RIOrganization)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, &gin_backend_utility.WebappError{
			Err:     err,
			Context: "apiv1CreateEngagement - Obtain org in context",
		})
		return
	}

	engagement := engagements.Engagement{}
	err = c.BindJSON(&engagement)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, &gin_backend_utility.WebappError{
			Err:     err,
			Context: "apiv1CreateEngagement - Read engagement from JSON body.",
		})
		return
	}
	engagement.OrgId = org.(*orgs.Organization).Id

	err = w.backend.itf.WrapDatabaseTx(w.middleware.GetAuditTrailIdWithEngagementOverride(&engagement, c), func(tx *sqlx.Tx) error {
		return w.backend.itf.Engagements.CreateEngagement(tx, &engagement)
	})

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, &gin_backend_utility.WebappError{
			Err:     err,
			Context: "apiv1CreateEngagement - Create engagement",
		})
		return
	}

	c.JSON(http.StatusOK, engagement)
}
