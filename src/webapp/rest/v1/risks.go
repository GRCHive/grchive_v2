package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"gitlab.com/grchive/grchive-v2/shared/backend"
	"gitlab.com/grchive/grchive-v2/shared/backend/engagements"
	"gitlab.com/grchive/grchive-v2/shared/backend/risks"
	"gitlab.com/grchive/grchive-v2/shared/gin_middleware/gin_backend_utility"
	"net/http"
)

func (w *WebappApplication) apiv1ListRisks(c *gin.Context) {
	engagement, err := w.middleware.GetResourceFromContext(c, backend.RIEngagement)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, &gin_backend_utility.WebappError{
			Err:     err,
			Context: "apiv1ListRisks - Obtain engagement in context",
		})
		return
	}

	risks, err := w.backend.itf.Risks.ListRisksForEngagementId(engagement.(*engagements.Engagement).Id)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, &gin_backend_utility.WebappError{
			Err:     err,
			Context: "apiv1ListRisks - List risks",
		})
		return
	}

	c.JSON(http.StatusOK, risks)
}

func (w *WebappApplication) apiv1CreateRisk(c *gin.Context) {
	engagement, err := w.middleware.GetResourceFromContext(c, backend.RIEngagement)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, &gin_backend_utility.WebappError{
			Err:     err,
			Context: "apiv1CreateRisk - Obtain engagement in context",
		})
		return
	}

	risk := risks.Risk{}
	err = c.BindJSON(&risk)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, &gin_backend_utility.WebappError{
			Err:     err,
			Context: "apiv1CreateRisk - Read risk from JSON body.",
		})
		return
	}
	risk.EngagementId = engagement.(*engagements.Engagement).Id

	err = w.backend.itf.WrapDatabaseTx(w.middleware.GetAuditTrailId(c), func(tx *sqlx.Tx) error {
		return w.backend.itf.Risks.CreateRisk(tx, &risk)
	})

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, &gin_backend_utility.WebappError{
			Err:     err,
			Context: "apiv1CreateRisk - Create risk",
		})
		return
	}

	c.JSON(http.StatusOK, risk)
}

func (w *WebappApplication) apiv1GetRisk(c *gin.Context) {
	risk, err := w.middleware.GetResourceFromContext(c, backend.RIRisk)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, &gin_backend_utility.WebappError{
			Err:     err,
			Context: "apiv1GetRisk - Obtain risk in context",
		})
		return
	}

	c.JSON(http.StatusOK, risk.(*risks.Risk))
}

func (w *WebappApplication) apiv1UpdateRisk(c *gin.Context) {
	currentRisk, err := w.middleware.GetResourceFromContext(c, backend.RIRisk)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, &gin_backend_utility.WebappError{
			Err:     err,
			Context: "apiv1UpdateRisk - Obtain risk in context",
		})
		return
	}

	risk := risks.Risk{}
	err = c.BindJSON(&risk)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, &gin_backend_utility.WebappError{
			Err:     err,
			Context: "apiv1UpdateRisk - Read risk from JSON body.",
		})
		return
	}

	tCurrentRisk := currentRisk.(*risks.Risk)
	risk.Id = tCurrentRisk.Id
	risk.EngagementId = tCurrentRisk.EngagementId

	err = w.backend.itf.WrapDatabaseTx(w.middleware.GetAuditTrailId(c), func(tx *sqlx.Tx) error {
		return w.backend.itf.Risks.UpdateRisk(tx, &risk)
	})

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, &gin_backend_utility.WebappError{
			Err:     err,
			Context: "apiv1UpdateRisk - Update risk",
		})
		return
	}

	c.JSON(http.StatusOK, risk)
}

func (w *WebappApplication) apiv1DeleteRisk(c *gin.Context) {
	currentRisk, err := w.middleware.GetResourceFromContext(c, backend.RIRisk)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, &gin_backend_utility.WebappError{
			Err:     err,
			Context: "apiv1DeleteRisk - Obtain risk in context",
		})
		return
	}

	err = w.backend.itf.WrapDatabaseTx(w.middleware.GetAuditTrailId(c), func(tx *sqlx.Tx) error {
		return w.backend.itf.Risks.DeleteRisk(tx, currentRisk.(*risks.Risk).Id)
	})

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, &gin_backend_utility.WebappError{
			Err:     err,
			Context: "apiv1DeleteRisk - Delete risk",
		})
		return
	}
	c.Status(http.StatusNoContent)
}
