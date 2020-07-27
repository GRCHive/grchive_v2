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

	singleR := r.Group("/:invId",
		w.middleware.LoadResourceIntoContext(resource, "invId"),
		w.middleware.CheckResourcePartOfEngagement(resource),
	)

	{
		singleR.GET("/",
			w.acl.ACLUserHasPermissions(crud.View),
			w.apiv1CreateGetInventory(resource),
		)

		singleR.PUT("/",
			w.acl.ACLUserHasPermissions(crud.Update),
			w.apiv1CreateUpdateInventory(resource),
		)

		singleR.DELETE("/",
			w.acl.ACLUserHasPermissions(crud.Delete),
			w.apiv1CreateDeleteInventory(resource),
		)

		w.addCommentEndpoints(
			resource,
			singleR,
			w.acl.ACLUserHasPermissions(crud.View),
		)
	}
}

func (w *WebappApplication) apiv1CreateDeleteInventory(resource backend.ResourceIdentifier) gin.HandlerFunc {
	it := backend.ResourceToInventoryType(resource)
	return func(c *gin.Context) {
		inv, err := w.middleware.GetResourceFromContext(c, resource)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, &gin_backend_utility.WebappError{
				Err:     err,
				Context: "apiv1CreateDeleteInventory - Obtain inventory in context",
				Code:    gin_backend_utility.GECBadRequest,
				Message: gin_backend_utility.GEMBadRequest,
			})
			return
		}

		id := inventory.GetInventoryId(inv)

		err = w.backend.itf.WrapDatabaseTx(w.middleware.GetAuditTrailId(c), func(tx *sqlx.Tx) error {
			return w.backend.itf.Inventory.DeleteInventory(tx, it, id)
		})

		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, &gin_backend_utility.WebappError{
				Err:     err,
				Context: "apiv1CreateDeleteInventory - Create inventory",
			})
			return
		}

		c.Status(http.StatusNoContent)
	}
}

func (w *WebappApplication) apiv1CreateUpdateInventory(resource backend.ResourceIdentifier) gin.HandlerFunc {
	it := backend.ResourceToInventoryType(resource)
	return func(c *gin.Context) {
		inv, err := w.middleware.GetResourceFromContext(c, resource)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, &gin_backend_utility.WebappError{
				Err:     err,
				Context: "apiv1CreateUpdateInventory - Obtain inventory in context",
				Code:    gin_backend_utility.GECBadRequest,
				Message: gin_backend_utility.GEMBadRequest,
			})
			return
		}
		invBase := inventory.GetBaseInventory(inv)

		updatedInv := inventory.CreateTypedInventory(it)
		err = c.BindJSON(updatedInv)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, &gin_backend_utility.WebappError{
				Err:     err,
				Context: "apiv1CreateUpdateInventory - Read inventory data from JSON body.",
				Code:    gin_backend_utility.GECBadRequest,
				Message: gin_backend_utility.GEMBadRequest,
			})
			return
		}
		updatedBase := inventory.GetBaseInventory(updatedInv)

		inventory.SetInventoryId(updatedInv, inventory.GetInventoryId(inv))
		updatedBase.Id = invBase.Id
		updatedBase.EngagementId = invBase.EngagementId
		updatedBase.UniqueId = invBase.UniqueId

		err = w.backend.itf.WrapDatabaseTx(w.middleware.GetAuditTrailId(c), func(tx *sqlx.Tx) error {
			return w.backend.itf.Inventory.UpdateInventory(tx, it, updatedInv)
		})

		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, &gin_backend_utility.WebappError{
				Err:     err,
				Context: "apiv1CreateUpdateInventory - Update inventory",
			})
			return
		}

		c.JSON(http.StatusOK, updatedInv)
	}
}

func (w *WebappApplication) apiv1CreateGetInventory(resource backend.ResourceIdentifier) gin.HandlerFunc {
	return func(c *gin.Context) {
		inv, err := w.middleware.GetResourceFromContext(c, resource)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, &gin_backend_utility.WebappError{
				Err:     err,
				Context: "apiv1CreateGetVendor - Obtain inventory in context",
				Code:    gin_backend_utility.GECBadRequest,
				Message: gin_backend_utility.GEMBadRequest,
			})
			return
		}
		c.JSON(http.StatusOK, inv)
	}
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
