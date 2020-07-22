package controls

func (m *ControlManager) ListControlsForEngagementId(engagementId int64) ([]*Control, error) {
	controls := make([]*Control, 0)
	err := m.db.Select(&controls, `
		SELECT *
		FROM controls
		WHERE engagement_id = $1
		ORDER BY id DESC
	`, engagementId)
	return controls, err
}
