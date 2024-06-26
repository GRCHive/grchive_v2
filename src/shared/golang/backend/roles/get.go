package roles

import (
	"errors"
	"github.com/jmoiron/sqlx"
)

var roleDoesNotHavePermission = errors.New("Role does not have request permissions.")

func (m *RoleManager) GetUserRolesForOrg(userId int64, orgId int64) ([]*Role, error) {
	role := []*Role{}
	err := m.db.Select(&role, `
		SELECT r.*
		FROM roles AS r
		INNER JOIN user_roles AS ur
			ON ur.role_id = r.id
		WHERE ur.user_id = $1 AND ur.org_id = $2
	`, userId, orgId)
	return role, err
}

func (m *RoleManager) UserHasPermissions(userId int64, orgId int64, permissions ...Permission) error {
	query, args, err := sqlx.In(`
		SELECT COUNT(DISTINCT p.id)
		FROM permissions AS p
		INNER JOIN role_permissions AS rp
			ON rp.permission_id = p.id
		INNER JOIN user_roles AS ur
			ON ur.role_id = rp.role_id
		WHERE ur.user_id = ?
			AND ur.org_id = ?
			AND p.human_name IN  (?)
	`, userId, orgId, permissions)

	if err != nil {
		return err
	}

	query = m.db.Rebind(query)
	count := 0
	err = m.db.Get(&count, query, args...)
	if err != nil {
		return err
	}

	if count != len(permissions) {
		return roleDoesNotHavePermission
	}

	return nil
}

func (m *RoleManager) UserHasPermissionsInRoles(userId int64, roles []*Role, permissions ...Permission) error {
	roleIds := make([]int64, len(roles))
	for i, r := range roles {
		roleIds[i] = r.Id
	}

	query, args, err := sqlx.In(`
		SELECT COUNT(DISTINCT p.id)
		FROM permissions AS p
		INNER JOIN role_permissions AS rp
			ON rp.permission_id = p.id
		INNER JOIN user_roles AS ur
			ON ur.role_id = rp.role_id
		WHERE ur.user_id = ?
			AND ur.role_id IN (?) 
			AND p.human_name IN  (?)
	`, userId, roleIds, permissions)

	if err != nil {
		return err
	}

	query = m.db.Rebind(query)
	count := 0
	err = m.db.Get(&count, query, args...)
	if err != nil {
		return err
	}

	if count != len(permissions) {
		return roleDoesNotHavePermission
	}

	return nil
}
