package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"gitlab.com/grchive/grchive-v2/shared/backend"
	"gitlab.com/grchive/grchive-v2/shared/backend/controls"
	"gitlab.com/grchive/grchive-v2/shared/backend/engagements"
	"gitlab.com/grchive/grchive-v2/shared/gin_middleware/gin_backend_utility"
	"net/http"
)

func (w *WebappApplication) apiv1ListControls(c *gin.Context) {
	engagement, err := w.middleware.GetResourceFromContext(c, backend.RIEngagement)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, &gin_backend_utility.WebappError{
			Err:     err,
			Context: "apiv1ListControls - Obtain engagement in context",
		})
		return
	}

	controls, err := w.backend.itf.Controls.ListControlsForEngagementId(engagement.(*engagements.Engagement).Id)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, &gin_backend_utility.WebappError{
			Err:     err,
			Context: "apiv1ListControls - List controls",
		})
		return
	}

	c.JSON(http.StatusOK, controls)
}

func (w *WebappApplication) apiv1CreateControl(c *gin.Context) {
	engagement, err := w.middleware.GetResourceFromContext(c, backend.RIEngagement)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, &gin_backend_utility.WebappError{
			Err:     err,
			Context: "apiv1CreateControl - Obtain engagement in context",
		})
		return
	}

	control := controls.Control{}
	err = c.BindJSON(&control)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, &gin_backend_utility.WebappError{
			Err:     err,
			Context: "apiv1CreateEngagement - Read control from JSON body.",
		})
		return
	}
	control.EngagementId = engagement.(*engagements.Engagement).Id

	err = w.backend.itf.WrapDatabaseTx(w.middleware.GetAuditTrailId(c), func(tx *sqlx.Tx) error {
		return w.backend.itf.Controls.CreateControl(tx, &control)
	})

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, &gin_backend_utility.WebappError{
			Err:     err,
			Context: "apiv1CreateControl - Create control",
		})
		return
	}

	c.JSON(http.StatusOK, control)
}
