package databases

func (m *DatabaseManager) GetNumDatabasesForEngagement(engagementId int64) (int, error) {
	ret := 0
	err := m.db.Get(&ret, `
		SELECT COUNT(id)
		FROM databases
		WHERE engagement_id = $1
	`, engagementId)
	return ret, err
}
