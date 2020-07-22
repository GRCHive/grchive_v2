package controls

func (m *ControlManager) GetControlFromId(id int64) (*Control, error) {
	control := Control{}
	err := m.db.Get(&control, `
		SELECT *
		FROM controls
		WHERE id = $1
	`, id)
	return &control, err
}
