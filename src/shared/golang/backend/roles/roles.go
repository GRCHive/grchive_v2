package roles

import (
	"github.com/jmoiron/sqlx"
)

type Permission string

// This needs to be kept in sync with the database.
const (
	POrgProfileView   Permission = "org.profile.view"
	POrgProfileCreate            = "org.profile.create"
	POrgProfileUpdate            = "org.profile.update"
	POrgRolesView                = "org.roles.view"
	POrgRolesUpdate              = "org.roles.update"
	POrgRolesDelete              = "org.roles.delete"
	POrgRolesCreate              = "org.roles.create"
)

type RoleManager struct {
	db *sqlx.DB
}

func CreateRoleManager(db *sqlx.DB) *RoleManager {
	return &RoleManager{
		db: db,
	}
}

type Role struct {
	Id          int64  `db:"id"`
	OrgId       int64  `db:"org_id"`
	Name        string `db:"name"`
	Description string `db:"description"`
	IsAdminRole bool   `db:"is_admin_role"`
}
