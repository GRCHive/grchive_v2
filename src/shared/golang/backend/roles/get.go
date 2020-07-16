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

func (m *RoleManager) UserHasPermissions(userId int64, permissions ...Permission) error {
	query, args, err := sqlx.In(`
		SELECT COUNT(*)
		FROM permissions AS p
		INNER JOIN role_permissions AS rp
			ON rp.permission_id = p.id
		INNER JOIN user_roles AS ur
			ON ur.role_id = rp.role_id
		WHERE ur.user_id = ? AND p.human_name IN  (?)
	`, userId, permissions)

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
