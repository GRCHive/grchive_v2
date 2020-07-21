package risks

func (m *RiskManager) ListRisksForEngagementId(engagementId int64) ([]*Risk, error) {
	ret := make([]*Risk, 0)
	err := m.db.Select(&ret, `
		SELECT *
		FROM risks
		WHERE engagement_id = $1
		ORDER BY id DESC
	`, engagementId)
	return ret, err
}
