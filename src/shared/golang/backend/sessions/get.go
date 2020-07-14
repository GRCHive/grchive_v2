package sessions

// Returns the session with the given ID. Returns a nil session if not found.
func (m *SessionManager) GetSessionFromId(id string) (*Session, error) {
	rows, err := m.db.Queryx(`
		SELECT *
		FROM user_sessions
		WHERE id = $1
	`, id)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, nil
	}

	sess := Session{}
	err = rows.StructScan(&sess)
	return &sess, err
}
