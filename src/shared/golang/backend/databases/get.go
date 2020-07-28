package databases

func (m *DatabaseManager) GetDatabaseFromId(dbId int64) (*Database, error) {
	db := Database{}
	err := m.db.Get(&db, `
		SELECT *
		FROM databases
		WHERE id = $1
	`, dbId)
	return &db, err
}
