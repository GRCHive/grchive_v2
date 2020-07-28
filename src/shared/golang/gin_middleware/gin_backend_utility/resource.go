package gin_backend_utility

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/grchive/grchive-v2/shared/backend"
	"gitlab.com/grchive/grchive-v2/shared/backend/controls"
	"gitlab.com/grchive/grchive-v2/shared/backend/databases"
	"gitlab.com/grchive/grchive-v2/shared/backend/engagements"
	"gitlab.com/grchive/grchive-v2/shared/backend/gl"
	"gitlab.com/grchive/grchive-v2/shared/backend/inventory"
	"gitlab.com/grchive/grchive-v2/shared/backend/orgs"
	"gitlab.com/grchive/grchive-v2/shared/backend/risks"
	"gitlab.com/grchive/grchive-v2/shared/backend/systems"
	"gitlab.com/grchive/grchive-v2/shared/backend/users"
	"gitlab.com/grchive/grchive-v2/shared/backend/vendors"
	"net/http"
)

func (m *MiddlewareClient) CheckResourcePartOfOrg(resource backend.ResourceIdentifier) gin.HandlerFunc {
	return func(c *gin.Context) {
		org, err := m.GetResourceFromContext(c, backend.RIOrganization)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, &WebappError{
				Err:     err,
				Context: "CheckResourcePartOfOrg - Get org",
				Code:    GECBadRequest,
				Message: GEMBadRequest,
			})
			return
		}

		torg := org.(*orgs.Organization)

		rsc, err := m.GetResourceFromContext(c, resource)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, &WebappError{
				Err:     err,
				Context: "CheckResourcePartOfOrg - Get resource",
				Code:    GECBadRequest,
				Message: GEMBadRequest,
			})
			return
		}

		mismatch := false
		switch resource {
		case backend.RIEngagement:
			trsc := rsc.(*engagements.Engagement)
			mismatch = trsc.OrgId != torg.Id
		case backend.RIUser:
			trsc := rsc.(*users.User)
			exists, err := m.Itf.Users.IsUserInOrg(torg.Id, trsc.Id)
			if err != nil {
				c.AbortWithError(http.StatusBadRequest, &WebappError{
					Err:     err,
					Context: "CheckResourcePartOfOrg - Failed to check user in org",
					Code:    GECBadRequest,
					Message: GEMBadRequestTree,
				})
				return
			}
			mismatch = !exists
		default:
			mismatch = true
		}

		if mismatch {
			c.AbortWithError(http.StatusBadRequest, &WebappError{
				Err:     nil,
				Context: "CheckResourcePartOfOrg - Resource Org mismatch",
				Code:    GECBadRequest,
				Message: GEMBadRequestTree,
			})

		} else {
			c.Next()
		}
	}
}

func (m *MiddlewareClient) CheckResourcePartOfEngagement(resource backend.ResourceIdentifier) gin.HandlerFunc {
	return func(c *gin.Context) {
		engagement, err := m.GetResourceFromContext(c, backend.RIEngagement)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, &WebappError{
				Err:     err,
				Context: "CheckResourcePartOfEngagement - Get engagement",
				Code:    GECBadRequest,
				Message: GEMBadRequest,
			})
			return
		}

		tengagement := engagement.(*engagements.Engagement)

		rsc, err := m.GetResourceFromContext(c, resource)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, &WebappError{
				Err:     err,
				Context: "CheckResourcePartOfEngagement - Get resource",
				Code:    GECBadRequest,
				Message: GEMBadRequest,
			})
			return
		}

		mismatch := false
		switch resource {
		case backend.RIRisk:
			trsc := rsc.(*risks.Risk)
			mismatch = trsc.EngagementId != tengagement.Id
		case backend.RIControl:
			trsc := rsc.(*controls.Control)
			mismatch = trsc.EngagementId != tengagement.Id
		case backend.RIGLAccount:
			trsc := rsc.(*gl.GLAccount)
			mismatch = trsc.EngagementId != tengagement.Id
		case backend.RIVendor:
			trsc := rsc.(*vendors.Vendor)
			mismatch = trsc.EngagementId != tengagement.Id
		case backend.RIInventoryServer:
			fallthrough
		case backend.RIInventoryDesktop:
			fallthrough
		case backend.RIInventoryLaptop:
			fallthrough
		case backend.RIInventoryMobile:
			fallthrough
		case backend.RIInventoryEmbedded:
			engId := inventory.GetInventoryEngagementId(rsc)
			mismatch = engId != tengagement.Id
		case backend.RIDatabase:
			trsc := rsc.(*databases.Database)
			mismatch = trsc.EngagementId != tengagement.Id
		case backend.RISystem:
			trsc := rsc.(*systems.System)
			mismatch = trsc.EngagementId != tengagement.Id
		default:
			mismatch = true
		}

		if mismatch {
			c.AbortWithError(http.StatusBadRequest, &WebappError{
				Err:     nil,
				Context: "CheckResourcePartOfEngagement - Resource Engagement mismatch",
				Code:    GECBadRequest,
				Message: GEMBadRequestTree,
			})

		} else {
			c.Next()
		}
	}
}

func (m *MiddlewareClient) CheckResourcePartOfVendor(resource backend.ResourceIdentifier) gin.HandlerFunc {
	return func(c *gin.Context) {
		vendor, err := m.GetResourceFromContext(c, backend.RIVendor)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, &WebappError{
				Err:     err,
				Context: "CheckResourcePartOfVendor - Get vendor",
				Code:    GECBadRequest,
				Message: GEMBadRequest,
			})
			return
		}

		tvendor := vendor.(*vendors.Vendor)
		rsc, err := m.GetResourceFromContext(c, resource)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, &WebappError{
				Err:     err,
				Context: "CheckResourcePartOfVendor - Get resource",
				Code:    GECBadRequest,
				Message: GEMBadRequest,
			})
			return
		}

		mismatch := false
		switch resource {
		case backend.RIVendorProduct:
			trsc := rsc.(*vendors.VendorProduct)
			mismatch = trsc.VendorId != tvendor.Id
		default:
			mismatch = true
		}

		if mismatch {
			c.AbortWithError(http.StatusBadRequest, &WebappError{
				Err:     nil,
				Context: "CheckResourcePartOfVendor - Resource Vendor mismatch",
				Code:    GECBadRequest,
				Message: GEMBadRequestTree,
			})
		} else {
			c.Next()
		}
	}
}
