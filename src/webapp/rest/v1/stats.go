package main

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/grchive/grchive-v2/shared/backend"
	"gitlab.com/grchive/grchive-v2/shared/backend/engagements"
	"gitlab.com/grchive/grchive-v2/shared/gin_middleware/gin_backend_utility"
	"net/http"
)

func (w *WebappApplication) apiv1GetEngagementScopingStats(c *gin.Context) {
	currentEngagement, err := w.middleware.GetResourceFromContext(c, backend.RIEngagement)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, &gin_backend_utility.WebappError{
			Err:     err,
			Context: "apiv1GetEngagementScopingStats - Obtain engagement in context",
			Code:    gin_backend_utility.GECBadRequest,
			Message: gin_backend_utility.GEMBadRequest,
		})
		return
	}

	stats, err := w.backend.itf.GetEngagementScopingStats(currentEngagement.(*engagements.Engagement).Id)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, &gin_backend_utility.WebappError{
			Err:     err,
			Context: "apiv1GetEngagementScopingStats - Get stats",
		})
		return
	}

	c.JSON(http.StatusOK, stats)
}
