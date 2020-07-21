package gin_backend_utility

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/grchive/grchive-v2/shared/backend"
	"gitlab.com/grchive/grchive-v2/shared/backend/engagements"
	"gitlab.com/grchive/grchive-v2/shared/backend/orgs"
	"gitlab.com/grchive/grchive-v2/shared/backend/risks"
	"net/http"
)

func (m *MiddlewareClient) CheckResourcePartOfOrg(resource backend.ResourceIdentifier) gin.HandlerFunc {
	return func(c *gin.Context) {
		org, err := m.GetResourceFromContext(c, backend.RIOrganization)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, &WebappError{
				Err:     err,
				Context: "CheckResourcePartOfOrg - Get org",
			})
			return
		}

		torg := org.(*orgs.Organization)

		rsc, err := m.GetResourceFromContext(c, resource)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, &WebappError{
				Err:     err,
				Context: "CheckResourcePartOfOrg - Get resource",
			})
			return
		}

		mismatch := false
		switch resource {
		case backend.RIEngagement:
			trsc := rsc.(*engagements.Engagement)
			mismatch = trsc.OrgId != torg.Id
		}

		if mismatch {
			c.AbortWithError(http.StatusBadRequest, &WebappError{
				Err:     nil,
				Context: "CheckResourcePartOfOrg - Resource Org mismatch",
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
			})
			return
		}

		tengagement := engagement.(*engagements.Engagement)

		rsc, err := m.GetResourceFromContext(c, resource)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, &WebappError{
				Err:     err,
				Context: "CheckResourcePartOfEngagement - Get resource",
			})
			return
		}

		mismatch := false
		switch resource {
		case backend.RIRisk:
			trsc := rsc.(*risks.Risk)
			mismatch = trsc.EngagementId != tengagement.Id
		}

		if mismatch {
			c.AbortWithError(http.StatusBadRequest, &WebappError{
				Err:     nil,
				Context: "CheckResourcePartOfEngagement - Resource Engagement mismatch",
			})

		} else {
			c.Next()
		}
	}
}
