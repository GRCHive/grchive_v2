package systems

func (m *SystemManager) GetSystemFromId(dbId int64) (*System, error) {
	db := System{}
	err := m.db.Get(&db, `
		SELECT *
		FROM systems
		WHERE id = $1
	`, dbId)
	return &db, err
}
