package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"gitlab.com/grchive/grchive-v2/shared/backend"
	"gitlab.com/grchive/grchive-v2/shared/backend/engagements"
	"gitlab.com/grchive/grchive-v2/shared/backend/gl"
	"gitlab.com/grchive/grchive-v2/shared/gin_middleware/gin_backend_utility"
	"net/http"
)

func (w *WebappApplication) apiv1ListGLAccounts(c *gin.Context) {
	engagement, err := w.middleware.GetResourceFromContext(c, backend.RIEngagement)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, &gin_backend_utility.WebappError{
			Err:     err,
			Context: "apiv1ListGLAccounts - Obtain engagement in context",
		})
		return
	}

	accs, err := w.backend.itf.GL.ListGLAccountsForEngagement(engagement.(*engagements.Engagement).Id)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, &gin_backend_utility.WebappError{
			Err:     err,
			Context: "apiv1ListGLAccounts - List GL accounts",
		})
		return
	}

	c.JSON(http.StatusOK, accs)
}

func (w *WebappApplication) apiv1CreateGLAccount(c *gin.Context) {
	engagement, err := w.middleware.GetResourceFromContext(c, backend.RIEngagement)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, &gin_backend_utility.WebappError{
			Err:     err,
			Context: "apiv1CreateGLAccount - Obtain engagement in context",
		})
		return
	}

	acc := gl.GLAccount{}
	err = c.BindJSON(&acc)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, &gin_backend_utility.WebappError{
			Err:     err,
			Context: "apiv1CreateGLAccount - Read account from JSON body.",
		})
		return
	}
	acc.EngagementId = engagement.(*engagements.Engagement).Id

	err = w.backend.itf.WrapDatabaseTx(w.middleware.GetAuditTrailId(c), func(tx *sqlx.Tx) error {
		return w.backend.itf.GL.CreateGLAccount(tx, &acc)
	})

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, &gin_backend_utility.WebappError{
			Err:     err,
			Context: "apiv1CreateGLAccount - Create account",
		})
		return
	}

	c.JSON(http.StatusOK, acc)
}
