package engagements

import (
	"gitlab.com/grchive/grchive-v2/shared/backend/roles"
)

func (m *EngagementManager) GetEngagementFromId(engId int64) (*Engagement, error) {
	eng := Engagement{}
	err := m.db.Get(&eng, `
		WITH latest_status AS (
			SELECT DISTINCT ON (engagement_id) engagement_id, is_closed
			FROM engagement_status
			ORDER BY engagement_id, id DESC	
		)
		SELECT eng.*, COALESCE(st.is_closed, false) AS "is_closed"
		FROM engagements AS eng
		LEFT JOIN latest_status AS st
			ON st.engagement_id = eng.id
		WHERE eng.id = $1
	`, engId)

	if err != nil {
		return nil, err
	}

	roles, err := m.GetRolesForEngagement(engId)
	if err != nil {
		return nil, err
	}

	eng.Roles = &roles
	return &eng, err
}

func (m *EngagementManager) GetRolesForEngagement(engId int64) ([]*roles.Role, error) {
	roles := make([]*roles.Role, 0)
	err := m.db.Select(&roles, `
		SELECT r.*
		FROM roles AS r
		INNER JOIN engagement_roles AS er
			ON er.role_id = r.id
		WHERE er.engagement_id = $1
	`, engId)
	return roles, err
}
