package roles

func (m *RoleManager) ListRolesForOrg(orgId int64) ([]*Role, error) {
	roles := make([]*Role, 0)
	err := m.db.Select(&roles, `
		SELECT r.*
		FROM roles AS r
		WHERE org_id = $1
	`, orgId)
	return roles, err
}

func (m *RoleManager) ListRolesForOrgAndUser(orgId int64, userId int64) ([]*Role, error) {
	roles := make([]*Role, 0)
	err := m.db.Select(&roles, `
		SELECT r.*
		FROM roles AS r
		INNER JOIN user_roles AS ur
			ON ur.role_id = r.id
		WHERE r.org_id = $1 AND ur.user_id = $2
	`, orgId, userId)
	return roles, err
}
