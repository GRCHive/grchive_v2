package systems

func (m *SystemManager) GetNumSystemsForEngagement(engagementId int64) (int, error) {
	ret := 0
	err := m.db.Get(&ret, `
		SELECT COUNT(id)
		FROM systems
		WHERE engagement_id = $1
	`, engagementId)
	return ret, err
}
