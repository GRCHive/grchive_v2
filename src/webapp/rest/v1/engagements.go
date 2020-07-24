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
			Code:    gin_backend_utility.GECBadRequest,
			Message: gin_backend_utility.GEMBadRequest,
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
			Code:    gin_backend_utility.GECBadRequest,
			Message: gin_backend_utility.GEMBadRequest,
		})
		return
	}

	user, err := w.middleware.GetCurrentUserFromContext(c)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, &gin_backend_utility.WebappError{
			Err:     err,
			Context: "apiv1ListAllOrgEngagementsForCurrentUser - Obtain current user in context",
			Code:    gin_backend_utility.GECBadRequest,
			Message: gin_backend_utility.GEMBadRequest,
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

type CreateEngagementInputs struct {
	NewEngagement  engagements.Engagement  `json:"new"`
	BaseEngagement *engagements.Engagement `json:"base"`
}

func (w *WebappApplication) apiv1CreateEngagement(c *gin.Context) {
	org, err := w.middleware.GetResourceFromContext(c, backend.RIOrganization)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, &gin_backend_utility.WebappError{
			Err:     err,
			Context: "apiv1CreateEngagement - Obtain org in context",
			Code:    gin_backend_utility.GECBadRequest,
			Message: gin_backend_utility.GEMBadRequest,
		})
		return
	}

	inputs := CreateEngagementInputs{}
	err = c.BindJSON(&inputs)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, &gin_backend_utility.WebappError{
			Err:     err,
			Context: "apiv1CreateEngagement - Read engagement data from JSON body.",
			Code:    gin_backend_utility.GECBadRequest,
			Message: gin_backend_utility.GEMBadRequest,
		})
		return
	}

	engagement := &inputs.NewEngagement
	engagement.OrgId = org.(*orgs.Organization).Id

	if engagement.Roles == nil || len(*engagement.Roles) == 0 {
		c.AbortWithError(http.StatusBadRequest, &gin_backend_utility.WebappError{
			Err:     nil,
			Context: "apiv1CreateEngagement - Must have at least one role.",
			Code:    gin_backend_utility.GECBadRequest,
			Message: gin_backend_utility.GEMBadRequest,
		})
		return
	}

	err = w.backend.itf.WrapDatabaseTx(w.middleware.GetAuditTrailIdWithEngagementOverride(engagement, c), func(tx *sqlx.Tx) error {
		return w.backend.itf.Engagements.CreateEngagement(tx, engagement)
	}, func(tx *sqlx.Tx) error {
		return w.backend.itf.Engagements.LinkEngagementToRoles(tx, engagement)
	}, func(tx *sqlx.Tx) error {
		if inputs.BaseEngagement != nil {
			return w.backend.itf.RollForwardEngagement(tx, inputs.BaseEngagement.Id, engagement.Id)
		}
		return nil
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

func (w *WebappApplication) apiv1UpdateEngagement(c *gin.Context) {
	currentEngagement, err := w.middleware.GetResourceFromContext(c, backend.RIEngagement)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, &gin_backend_utility.WebappError{
			Err:     err,
			Context: "apiv1UpdateEngagement - Obtain engagement in context",
			Code:    gin_backend_utility.GECBadRequest,
			Message: gin_backend_utility.GEMBadRequest,
		})
		return
	}

	editEngagement := engagements.Engagement{}
	err = c.BindJSON(&editEngagement)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, &gin_backend_utility.WebappError{
			Err:     err,
			Context: "apiv1UpdateEngagement - Obtain org in request.",
			Code:    gin_backend_utility.GECBadRequest,
			Message: gin_backend_utility.GEMBadRequest,
		})
		return
	}

	typedCurrentEng := currentEngagement.(*engagements.Engagement)
	typedCurrentEng.Name = editEngagement.Name
	typedCurrentEng.Description = editEngagement.Description
	typedCurrentEng.StartTime = editEngagement.StartTime
	typedCurrentEng.EndTime = editEngagement.EndTime

	err = w.backend.itf.WrapDatabaseTx(w.middleware.GetAuditTrailId(c), func(tx *sqlx.Tx) error {
		return w.backend.itf.Engagements.UpdateEngagement(tx, typedCurrentEng)
	})

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, &gin_backend_utility.WebappError{
			Err:     err,
			Context: "apiv1UpdateEngagement - Update engagement",
		})
		return
	}

	c.JSON(http.StatusOK, typedCurrentEng)
}

func (w *WebappApplication) apiv1GetEngagement(c *gin.Context) {
	eng, err := w.middleware.GetResourceFromContext(c, backend.RIEngagement)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, &gin_backend_utility.WebappError{
			Err:     err,
			Context: "apiv1GetEngagement - Obtain engagement in context",
			Code:    gin_backend_utility.GECBadRequest,
			Message: gin_backend_utility.GEMBadRequest,
		})
		return
	}

	c.JSON(http.StatusOK, eng.(*engagements.Engagement))
}
