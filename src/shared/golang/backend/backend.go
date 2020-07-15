package backend

import (
	"github.com/jmoiron/sqlx"
	"gitlab.com/grchive/grchive-v2/shared/backend/orgs"
	"gitlab.com/grchive/grchive-v2/shared/backend/sessions"
	"gitlab.com/grchive/grchive-v2/shared/backend/users"
	"gitlab.com/grchive/grchive-v2/shared/backend/utility"
)

type BackendInterface struct {
	db       *sqlx.DB
	Sessions *sessions.SessionManager
	Users    *users.UserManager
	Orgs     *orgs.OrgManager
}

func CreateBackendInterface(db *sqlx.DB) *BackendInterface {
	return &BackendInterface{
		db:       db,
		Sessions: sessions.CreateSessionManager(db),
		Users:    users.CreateUserManager(db),
		Orgs:     orgs.CreateOrgManager(db),
	}
}

func (m *BackendInterface) WrapDatabaseTx(fns ...utility.TxHandler) error {
	return utility.WrapDatabaseTx(m.db, fns...)
}
