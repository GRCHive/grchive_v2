package systems

func (m *SystemManager) ListSystemsForEngagement(engagementId int64) ([]*System, error) {
	dbs := []*System{}
	err := m.db.Select(&dbs, `
		SELECT *
		FROM systems
		WHERE engagement_id = $1
		ORDER BY id DESC
	`, engagementId)
	return dbs, err
}
