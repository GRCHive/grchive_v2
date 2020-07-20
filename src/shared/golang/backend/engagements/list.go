package engagements

func (m *EngagementManager) ListEngagementsForOrgId(orgId int64) ([]*Engagement, error) {
	engagements := []*Engagement{}
	err := m.db.Select(&engagements, `
		WITH latest_status AS (
			SELECT DISTINCT ON (engagement_id) engagement_id, is_closed
			FROM engagement_status
			ORDER BY engagement_id, id DESC	
		)
		SELECT eng.*, COALESCE(st.is_closed, false) AS "is_closed"
		FROM engagements AS eng
		LEFT JOIN latest_status AS st
			ON st.engagement_id = eng.id
		WHERE eng.org_id = $1
		ORDER BY eng.id DESC
	`, orgId)
	return engagements, err
}

func (m *EngagementManager) ListEngagementsForOrgIdAndUserId(orgId int64, userId int64) ([]*Engagement, error) {
	engagements := []*Engagement{}
	err := m.db.Select(&engagements, `
		WITH latest_status AS (
			SELECT DISTINCT ON (engagement_id) engagement_id, is_closed
			FROM engagement_status
			ORDER BY engagement_id, id DESC	
		)
		SELECT eng.*, COALESCE(st.is_closed, false) AS "is_closed"
		FROM engagements AS eng
		INNER JOIN engagement_roles AS er
			ON er.engagement_id = eng.id
		INNER JOIN user_roles AS ur
			ON ur.role_id = er.role_id
		LEFT JOIN latest_status AS st
			ON st.engagement_id = eng.id
		WHERE eng.org_id = $1 AND ur.user_id = $2
		ORDER BY eng.id DESC
	`, orgId, userId)
	return engagements, err
}
