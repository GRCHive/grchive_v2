package gin_backend_utility

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gitlab.com/grchive/grchive-v2/shared/backend"
	"gitlab.com/grchive/grchive-v2/shared/backend/engagements"
	"gitlab.com/grchive/grchive-v2/shared/backend/gl"
	"net/http"
)

const generalLedgerContextKey = "grchive-general-ledger"

func (m *MiddlewareClient) LoadGeneralLedgerIntoContext(c *gin.Context) {
	engagement, err := m.GetResourceFromContext(c, backend.RIEngagement)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, &WebappError{
			Err:     err,
			Context: "LoadGeneralLedgerIntoContext - Obtain engagement in context",
			Code:    GECBadRequest,
			Message: GEMBadRequest,
		})
		return
	}

	teng := engagement.(*engagements.Engagement)
	ledger, err := m.Itf.GL.GetGeneralLedgerForEngagement(teng.Id)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, &WebappError{
			Err:     err,
			Context: "LoadGeneralLedgerIntoContext - Load general ledger",
			Code:    GECBadRequest,
			Message: GEMBadRequest,
		})
		return
	}

	c.Set(generalLedgerContextKey, ledger)
	c.Next()
}

func (m *MiddlewareClient) GetGeneralLedgerFromContext(c *gin.Context) (*gl.GeneralLedger, error) {
	val, ok := c.Get(generalLedgerContextKey)
	if !ok {
		return nil, errors.New("Could not find general ledger in context.")
	}

	return val.(*gl.GeneralLedger), nil

}
