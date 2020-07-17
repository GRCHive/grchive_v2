package backend

import (
	"github.com/jmoiron/sqlx"
	"gitlab.com/grchive/grchive-v2/shared/backend/audit"
	"gitlab.com/grchive/grchive-v2/shared/backend/orgs"
	"gitlab.com/grchive/grchive-v2/shared/backend/roles"
	"gitlab.com/grchive/grchive-v2/shared/backend/sessions"
	"gitlab.com/grchive/grchive-v2/shared/backend/users"
	"gitlab.com/grchive/grchive-v2/shared/backend/utility"
)

type BackendInterface struct {
	db       *sqlx.DB
	Sessions *sessions.SessionManager
	Users    *users.UserManager
	Orgs     *orgs.OrgManager
	Roles    *roles.RoleManager
}

func CreateBackendInterface(db *sqlx.DB) *BackendInterface {
	return &BackendInterface{
		db:       db,
		Sessions: sessions.CreateSessionManager(db),
		Users:    users.CreateUserManager(db),
		Orgs:     orgs.CreateOrgManager(db),
		Roles:    roles.CreateRoleManager(db),
	}
}

func (m *BackendInterface) WrapDatabaseTx(auditId audit.AuditTrailId, fns ...utility.TxHandler) error {
	return audit.WrapDatabaseAuditTx(m.db, auditId, fns...)
}
