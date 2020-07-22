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

func (m *UserManager) GetUserFromId(id int64) (*User, error) {
	rows, err := m.db.Queryx(`
		SELECT *
		FROM users
		WHERE id = $1
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

func (m *UserManager) MustGetUserFromId(id int64) (*User, error) {
	user := User{}
	err := m.db.Get(&user, `
		SELECT *
		FROM users
		WHERE id = $1
	`, id)
	return &user, err
}

func (m *UserManager) IsUserInOrg(orgId int64, userId int64) (bool, error) {
	ret := false
	err := m.db.Get(&ret, `
		SELECT EXISTS(
			SELECT 1
			FROM users AS u
			INNER JOIN user_orgs AS uo
				ON uo.user_id = u.id
			WHERE uo.org_id = $1
				AND u.id = $2
		)
	`, orgId, userId)
	return ret, err
}
