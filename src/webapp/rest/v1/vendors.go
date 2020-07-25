package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"gitlab.com/grchive/grchive-v2/shared/backend"
	"gitlab.com/grchive/grchive-v2/shared/backend/engagements"
	"gitlab.com/grchive/grchive-v2/shared/backend/vendors"
	"gitlab.com/grchive/grchive-v2/shared/gin_middleware/gin_backend_utility"
	"net/http"
)

func (w *WebappApplication) apiv1ListVendors(c *gin.Context) {
	engagement, err := w.middleware.GetResourceFromContext(c, backend.RIEngagement)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, &gin_backend_utility.WebappError{
			Err:     err,
			Context: "apiv1ListVendors - Obtain engagement in context",
			Code:    gin_backend_utility.GECBadRequest,
			Message: gin_backend_utility.GEMBadRequest,
		})
		return
	}
	teng := engagement.(*engagements.Engagement)

	retVendors, err := w.backend.itf.Vendors.ListVendorsForEngagement(teng.Id)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, &gin_backend_utility.WebappError{
			Err:     err,
			Context: "apiv1ListVendors - List vendors",
		})
		return
	}

	c.JSON(http.StatusOK, retVendors)
}

func (w *WebappApplication) apiv1CreateVendor(c *gin.Context) {
	engagement, err := w.middleware.GetResourceFromContext(c, backend.RIEngagement)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, &gin_backend_utility.WebappError{
			Err:     err,
			Context: "apiv1CreateVendor - Obtain engagement in context",
			Code:    gin_backend_utility.GECBadRequest,
			Message: gin_backend_utility.GEMBadRequest,
		})
		return
	}
	teng := engagement.(*engagements.Engagement)

	vendor := vendors.Vendor{}
	err = c.BindJSON(&vendor)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, &gin_backend_utility.WebappError{
			Err:     err,
			Context: "apiv1CreateVendor - Read vendor from JSON body.",
			Code:    gin_backend_utility.GECBadRequest,
			Message: gin_backend_utility.GEMBadRequest,
		})
		return
	}
	vendor.EngagementId = teng.Id

	err = w.backend.itf.WrapDatabaseTx(w.middleware.GetAuditTrailId(c), func(tx *sqlx.Tx) error {
		return w.backend.itf.Vendors.CreateVendor(tx, &vendor)
	})

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, &gin_backend_utility.WebappError{
			Err:     err,
			Context: "apiv1CreateVendor - Create vendor",
		})
		return
	}

	c.JSON(http.StatusOK, vendor)
}

func (w *WebappApplication) apiv1GetVendor(c *gin.Context) {
	vendor, err := w.middleware.GetResourceFromContext(c, backend.RIVendor)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, &gin_backend_utility.WebappError{
			Err:     err,
			Context: "apiv1GetVendor - Obtain vendor in context",
			Code:    gin_backend_utility.GECBadRequest,
			Message: gin_backend_utility.GEMBadRequest,
		})
		return
	}

	c.JSON(http.StatusOK, vendor.(*vendors.Vendor))
}

func (w *WebappApplication) apiv1UpdateVendor(c *gin.Context) {
	currentVendor, err := w.middleware.GetResourceFromContext(c, backend.RIVendor)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, &gin_backend_utility.WebappError{
			Err:     err,
			Context: "apiv1UpdateVendor - Obtain vendor in context",
			Code:    gin_backend_utility.GECBadRequest,
			Message: gin_backend_utility.GEMBadRequest,
		})
		return
	}

	vendor := vendors.Vendor{}
	err = c.BindJSON(&vendor)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, &gin_backend_utility.WebappError{
			Err:     err,
			Context: "apiv1UpdateVendor - Read vendor from JSON body.",
			Code:    gin_backend_utility.GECBadRequest,
			Message: gin_backend_utility.GEMBadRequest,
		})
		return
	}

	tCurrentVendor := currentVendor.(*vendors.Vendor)
	vendor.Id = tCurrentVendor.Id
	vendor.EngagementId = tCurrentVendor.EngagementId

	err = w.backend.itf.WrapDatabaseTx(w.middleware.GetAuditTrailId(c), func(tx *sqlx.Tx) error {
		return w.backend.itf.Vendors.UpdateVendor(tx, &vendor)
	})

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, &gin_backend_utility.WebappError{
			Err:     err,
			Context: "apiv1UpdateVendor - Update vendor",
		})
		return
	}

	c.JSON(http.StatusOK, vendor)
}

func (w *WebappApplication) apiv1DeleteVendor(c *gin.Context) {
	currentVendor, err := w.middleware.GetResourceFromContext(c, backend.RIVendor)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, &gin_backend_utility.WebappError{
			Err:     err,
			Context: "apiv1DeleteVendor - Obtain vendor in context",
			Code:    gin_backend_utility.GECBadRequest,
			Message: gin_backend_utility.GEMBadRequest,
		})
		return
	}

	err = w.backend.itf.WrapDatabaseTx(w.middleware.GetAuditTrailId(c), func(tx *sqlx.Tx) error {
		return w.backend.itf.Vendors.DeleteVendor(tx, currentVendor.(*vendors.Vendor).Id)
	})

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, &gin_backend_utility.WebappError{
			Err:     err,
			Context: "apiv1DeleteVendor - Delete vendor",
		})
		return
	}
	c.Status(http.StatusNoContent)
}

func (w *WebappApplication) apiv1ListVendorProducts(c *gin.Context) {
	currentVendor, err := w.middleware.GetResourceFromContext(c, backend.RIVendor)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, &gin_backend_utility.WebappError{
			Err:     err,
			Context: "apiv1ListVendorProducts - Obtain vendor in context",
			Code:    gin_backend_utility.GECBadRequest,
			Message: gin_backend_utility.GEMBadRequest,
		})
		return
	}

	tVendor := currentVendor.(*vendors.Vendor)
	products, err := w.backend.itf.Vendors.ListVendorProductsForVendor(tVendor.Id)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, &gin_backend_utility.WebappError{
			Err:     err,
			Context: "apiv1ListVendorProducts - Get products",
		})
		return
	}
	c.JSON(http.StatusOK, products)
}

func (w *WebappApplication) apiv1CreateVendorProduct(c *gin.Context) {
	currentVendor, err := w.middleware.GetResourceFromContext(c, backend.RIVendor)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, &gin_backend_utility.WebappError{
			Err:     err,
			Context: "apiv1CreateVendorProduct - Obtain vendor in context",
			Code:    gin_backend_utility.GECBadRequest,
			Message: gin_backend_utility.GEMBadRequest,
		})
		return
	}

	tVendor := currentVendor.(*vendors.Vendor)

	product := vendors.VendorProduct{}
	err = c.BindJSON(&product)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, &gin_backend_utility.WebappError{
			Err:     err,
			Context: "apiv1CreateVendorVendor - Read product from JSON body.",
			Code:    gin_backend_utility.GECBadRequest,
			Message: gin_backend_utility.GEMBadRequest,
		})
		return
	}
	product.VendorId = tVendor.Id

	err = w.backend.itf.WrapDatabaseTx(w.middleware.GetAuditTrailId(c), func(tx *sqlx.Tx) error {
		return w.backend.itf.Vendors.CreateVendorProduct(tx, &product)
	})

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, &gin_backend_utility.WebappError{
			Err:     err,
			Context: "apiv1CreateVendorProduct - Create product",
		})
		return
	}

	c.JSON(http.StatusOK, product)
}

func (w *WebappApplication) apiv1GetVendorProduct(c *gin.Context) {
	product, err := w.middleware.GetResourceFromContext(c, backend.RIVendorProduct)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, &gin_backend_utility.WebappError{
			Err:     err,
			Context: "apiv1GetVendorProduct - Obtain vendor product in context",
			Code:    gin_backend_utility.GECBadRequest,
			Message: gin_backend_utility.GEMBadRequest,
		})
		return
	}

	c.JSON(http.StatusOK, product.(*vendors.VendorProduct))
}

func (w *WebappApplication) apiv1UpdateVendorProduct(c *gin.Context) {
	product, err := w.middleware.GetResourceFromContext(c, backend.RIVendorProduct)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, &gin_backend_utility.WebappError{
			Err:     err,
			Context: "apiv1UpdateVendorProduct - Obtain product in context",
			Code:    gin_backend_utility.GECBadRequest,
			Message: gin_backend_utility.GEMBadRequest,
		})
		return
	}

	newProduct := vendors.VendorProduct{}
	err = c.BindJSON(&newProduct)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, &gin_backend_utility.WebappError{
			Err:     err,
			Context: "apiv1UpdateVendorProduct - Read vendor product from JSON body.",
			Code:    gin_backend_utility.GECBadRequest,
			Message: gin_backend_utility.GEMBadRequest,
		})
		return
	}

	tproduct := product.(*vendors.VendorProduct)
	newProduct.Id = tproduct.Id
	newProduct.VendorId = tproduct.VendorId

	err = w.backend.itf.WrapDatabaseTx(w.middleware.GetAuditTrailId(c), func(tx *sqlx.Tx) error {
		return w.backend.itf.Vendors.UpdateVendorProduct(tx, &newProduct)
	})

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, &gin_backend_utility.WebappError{
			Err:     err,
			Context: "apiv1UpdateVendorProduct - Update vendor",
		})
		return
	}

	c.JSON(http.StatusOK, newProduct)
}

func (w *WebappApplication) apiv1DeleteVendorProduct(c *gin.Context) {
	currentVendorProduct, err := w.middleware.GetResourceFromContext(c, backend.RIVendorProduct)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, &gin_backend_utility.WebappError{
			Err:     err,
			Context: "apiv1DeleteVendorProduct - Obtain vendor in context",
			Code:    gin_backend_utility.GECBadRequest,
			Message: gin_backend_utility.GEMBadRequest,
		})
		return
	}

	err = w.backend.itf.WrapDatabaseTx(w.middleware.GetAuditTrailId(c), func(tx *sqlx.Tx) error {
		return w.backend.itf.Vendors.DeleteVendorProduct(tx, currentVendorProduct.(*vendors.VendorProduct).Id)
	})

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, &gin_backend_utility.WebappError{
			Err:     err,
			Context: "apiv1DeleteVendorProduct - Delete vendor",
		})
		return
	}
	c.Status(http.StatusNoContent)
}
