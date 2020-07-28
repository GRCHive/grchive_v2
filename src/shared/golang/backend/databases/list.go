package databases

func (m *DatabaseManager) ListDatabasesForEngagement(engagementId int64) ([]*Database, error) {
	dbs := []*Database{}
	err := m.db.Select(&dbs, `
		SELECT *
		FROM databases
		WHERE engagement_id = $1
		ORDER BY id DESC
	`, engagementId)
	return dbs, err
}
