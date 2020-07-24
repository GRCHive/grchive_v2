package risks

func (m *RiskManager) GetNumRisksForEngagement(engagementId int64) (int, error) {
	ret := 0
	err := m.db.Get(&ret, `
		SELECT COUNT(id)
		FROM risks
		WHERE engagement_id = $1
	`, engagementId)
	return ret, err
}
