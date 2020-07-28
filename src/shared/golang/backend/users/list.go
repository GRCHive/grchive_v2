package users

func (m *UserManager) GetOrgUsers(orgId int64) ([]*User, error) {
	users := []*User{}
	err := m.db.Select(&users, `
		SELECT u.*
		FROM users AS u
		INNER JOIN user_orgs AS uo
			ON uo.user_id = u.id
		WHERE uo.org_id = $1
		ORDER BY u.id DESC
	`, orgId)
	return users, err
}
