package gin_backend_utility

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gitlab.com/grchive/grchive-v2/shared/backend"
	"gitlab.com/grchive/grchive-v2/shared/backend/controls"
	"gitlab.com/grchive/grchive-v2/shared/backend/databases"
	"gitlab.com/grchive/grchive-v2/shared/backend/engagements"
	"gitlab.com/grchive/grchive-v2/shared/backend/gl"
	"gitlab.com/grchive/grchive-v2/shared/backend/inventory"
	"gitlab.com/grchive/grchive-v2/shared/backend/orgs"
	"gitlab.com/grchive/grchive-v2/shared/backend/risks"
	"gitlab.com/grchive/grchive-v2/shared/backend/vendors"
	"net/http"
)

const threadIdContextKey = "grchive-comment-thread-id"

func (m *MiddlewareClient) GetCommentThreadIdFromContext(c *gin.Context) (int64, error) {
	val, ok := c.Get(threadIdContextKey)
	if !ok {
		return -1, errors.New("Could not find comment thread ID.")
	}

	return val.(int64), nil
}

func (m *MiddlewareClient) LoadCommentThreadIdIntoContext(resource backend.ResourceIdentifier) gin.HandlerFunc {
	return func(c *gin.Context) {
		org, err := m.GetResourceFromContext(c, backend.RIOrganization)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, &WebappError{
				Err:     err,
				Context: "LoadCommentThreadIdIntoContext - Get Org in context",
				Code:    GECBadRequest,
				Message: GEMBadRequest,
			})
			return
		}

		engagement, err := m.GetResourceFromContext(c, backend.RIEngagement)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, &WebappError{
				Err:     err,
				Context: "LoadCommentThreadIdIntoContext - Get engagement in context",
				Code:    GECBadRequest,
				Message: GEMBadRequest,
			})
			return
		}

		var rsc interface{}

		switch resource {
		case backend.RIGeneralLedger:
			rsc, err = m.GetGeneralLedgerFromContext(c)
		default:
			rsc, err = m.GetResourceFromContext(c, resource)
		}
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, &WebappError{
				Err:     err,
				Context: "LoadCommentThreadIdIntoContext - Get resource in context",
				Code:    GECBadRequest,
				Message: GEMBadRequest,
			})
			return
		}

		torg := org.(*orgs.Organization)
		teng := engagement.(*engagements.Engagement)

		var threadId int64

		switch resource {
		case backend.RIRisk:
			trsc := rsc.(*risks.Risk)
			threadId, err = m.Itf.Comments.GetThreadIdForRisk(trsc.Id, teng.Id, torg.Id)
		case backend.RIControl:
			trsc := rsc.(*controls.Control)
			threadId, err = m.Itf.Comments.GetThreadIdForControl(trsc.Id, teng.Id, torg.Id)
		case backend.RIGeneralLedger:
			trsc := rsc.(*gl.GeneralLedger)
			threadId, err = m.Itf.Comments.GetThreadIdForGeneralLedger(trsc.Id, teng.Id, torg.Id)
		case backend.RIGLAccount:
			trsc := rsc.(*gl.GLAccount)
			threadId, err = m.Itf.Comments.GetThreadIdForGLAccount(trsc.Id, teng.Id, torg.Id)
		case backend.RIVendor:
			trsc := rsc.(*vendors.Vendor)
			threadId, err = m.Itf.Comments.GetThreadIdForVendor(trsc.Id, teng.Id, torg.Id)
		case backend.RIInventoryServer:
			trsc := rsc.(*inventory.InventoryServer)
			threadId, err = m.Itf.Comments.GetThreadIdForServer(trsc.Id, teng.Id, torg.Id)
		case backend.RIInventoryDesktop:
			trsc := rsc.(*inventory.InventoryDesktop)
			threadId, err = m.Itf.Comments.GetThreadIdForDesktop(trsc.Id, teng.Id, torg.Id)
		case backend.RIInventoryLaptop:
			trsc := rsc.(*inventory.InventoryLaptop)
			threadId, err = m.Itf.Comments.GetThreadIdForLaptop(trsc.Id, teng.Id, torg.Id)
		case backend.RIInventoryMobile:
			trsc := rsc.(*inventory.InventoryMobile)
			threadId, err = m.Itf.Comments.GetThreadIdForMobile(trsc.Id, teng.Id, torg.Id)
		case backend.RIInventoryEmbedded:
			trsc := rsc.(*inventory.InventoryEmbedded)
			threadId, err = m.Itf.Comments.GetThreadIdForEmbedded(trsc.Id, teng.Id, torg.Id)
		case backend.RIDatabase:
			trsc := rsc.(*databases.Database)
			threadId, err = m.Itf.Comments.GetThreadIdForDatabase(trsc.Id, teng.Id, torg.Id)
		default:
			err = errors.New("Unsupported resource for comments.")
		}

		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, &WebappError{
				Err:     err,
				Context: "LoadCommentThreadIdIntoContext - Get thread id",
			})
			return
		}

		c.Set(threadIdContextKey, threadId)
		c.Next()
	}
}
