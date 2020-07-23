package gin_backend_utility

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/grchive/grchive-v2/shared/backend"
	"gitlab.com/grchive/grchive-v2/shared/backend/audit"
	"gitlab.com/grchive/grchive-v2/shared/backend/engagements"
	"gitlab.com/grchive/grchive-v2/shared/backend/orgs"
	"gitlab.com/grchive/grchive-v2/shared/backend/users"
)

func (m *MiddlewareClient) GetAuditTrailId(c *gin.Context) audit.AuditTrailId {
	ret := audit.AuditTrailId{}

	user, err := m.GetCurrentUserFromContext(c)
	if err == nil {
		ret.UserId = &user.Id

		ipAddr := c.ClientIP()
		ret.IpAddress = &ipAddr
	}

	org, err := m.GetResourceFromContext(c, backend.RIOrganization)
	if err == nil {
		ret.OrgId = &org.(*orgs.Organization).Id
	}

	eng, err := m.GetResourceFromContext(c, backend.RIEngagement)
	if err == nil {
		ret.EngagementId = &eng.(*engagements.Engagement).Id
	}

	threadId, err := m.GetCommentThreadIdFromContext(c)
	if err == nil {
		ret.ThreadId = &threadId
	}

	return ret
}

func (m *MiddlewareClient) GetAuditTrailIdWithUserOverride(u *users.User, c *gin.Context) audit.AuditTrailId {
	ret := m.GetAuditTrailId(c)
	if u != nil {
		ret.UserId = &u.Id

		ipAddr := c.ClientIP()
		ret.IpAddress = &ipAddr
	}
	return ret
}

func (m *MiddlewareClient) GetAuditTrailIdWithOrgOverride(o *orgs.Organization, c *gin.Context) audit.AuditTrailId {
	ret := m.GetAuditTrailId(c)
	if o != nil {
		ret.OrgId = &o.Id
	}
	return ret
}

func (m *MiddlewareClient) GetAuditTrailIdWithEngagementOverride(e *engagements.Engagement, c *gin.Context) audit.AuditTrailId {
	ret := m.GetAuditTrailId(c)
	if e != nil {
		ret.EngagementId = &e.Id
	}
	return ret
}
