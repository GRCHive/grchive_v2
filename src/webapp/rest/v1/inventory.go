package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"gitlab.com/grchive/grchive-v2/shared/backend"
	"gitlab.com/grchive/grchive-v2/shared/backend/engagements"
	"gitlab.com/grchive/grchive-v2/shared/backend/inventory"
	"gitlab.com/grchive/grchive-v2/shared/gin_middleware/gin_backend_utility"
	"net/http"
)

func (w *WebappApplication) apiv1CreateInventoryEndpoints(parentR *gin.RouterGroup, prefix string, resource backend.ResourceIdentifier) {
	crud := backend.ResourceToCrudPermissions(resource)

	r := parentR.Group("/" + prefix)
	r.GET(
		"/",
		w.acl.ACLUserHasPermissions(crud.List),
		w.apiv1CreateListInventory(resource),
	)

	r.POST(
		"/",
		w.acl.ACLUserHasPermissions(crud.Create),
		w.apiv1CreateCreateInventory(resource),
	)
}

func (w *WebappApplication) apiv1CreateListInventory(resource backend.ResourceIdentifier) gin.HandlerFunc {
	it := backend.ResourceToInventoryType(resource)
	return func(c *gin.Context) {
		engagement, err := w.middleware.GetResourceFromContext(c, backend.RIEngagement)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, &gin_backend_utility.WebappError{
				Err:     err,
				Context: "apiv1CreateListInventory - Obtain engagement in context",
				Code:    gin_backend_utility.GECBadRequest,
				Message: gin_backend_utility.GEMBadRequest,
			})
			return
		}
		teng := engagement.(*engagements.Engagement)

		retInventory, err := w.backend.itf.Inventory.ListInventoryForEngagement(it, teng.Id)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, &gin_backend_utility.WebappError{
				Err:     err,
				Context: "apiv1CreateListInventory - List inventory",
			})
			return
		}

		c.JSON(http.StatusOK, retInventory)
	}
}

func (w *WebappApplication) apiv1CreateCreateInventory(resource backend.ResourceIdentifier) gin.HandlerFunc {
	it := backend.ResourceToInventoryType(resource)
	return func(c *gin.Context) {
		engagement, err := w.middleware.GetResourceFromContext(c, backend.RIEngagement)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, &gin_backend_utility.WebappError{
				Err:     err,
				Context: "apiv1CreateCreateInventory - Obtain engagement in context",
				Code:    gin_backend_utility.GECBadRequest,
				Message: gin_backend_utility.GEMBadRequest,
			})
			return
		}
		teng := engagement.(*engagements.Engagement)

		inv := inventory.CreateTypedInventory(it)
		err = c.BindJSON(inv)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, &gin_backend_utility.WebappError{
				Err:     err,
				Context: "apiv1CreateCreateInventory - Read inventory data from JSON body.",
				Code:    gin_backend_utility.GECBadRequest,
				Message: gin_backend_utility.GEMBadRequest,
			})
			return
		}
		inventory.SetInventoryEngagementId(inv, teng.Id)

		err = w.backend.itf.WrapDatabaseTx(w.middleware.GetAuditTrailId(c), func(tx *sqlx.Tx) error {
			return w.backend.itf.Inventory.CreateInventory(tx, it, inv)
		})

		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, &gin_backend_utility.WebappError{
				Err:     err,
				Context: "apiv1CreateCreateInventory - Create inventory",
			})
			return
		}

		c.JSON(http.StatusOK, inv)
	}
}
