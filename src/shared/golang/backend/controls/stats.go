package controls

func (m *ControlManager) GetNumControlsForEngagement(engagementId int64) (int, error) {
	ret := 0
	err := m.db.Get(&ret, `
		SELECT COUNT(id)
		FROM controls
		WHERE engagement_id = $1
	`, engagementId)
	return ret, err
}
