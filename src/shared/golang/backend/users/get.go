package users

// Returns the user associated with the FusionAuthId. If this function returns
// nil for the User then the user does not exist. If err is set, the return value
// of User is undefined.
func (m *UserManager) GetUserFromFusionAuthId(id string) (*User, error) {
	rows, err := m.db.Queryx(`
		SELECT *
		FROM users
		WHERE fa_user_id = $1
	`, id)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, nil
	}

	user := User{}
	err = rows.StructScan(&user)
	return &user, err
}
