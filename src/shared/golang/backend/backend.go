package backend

import (
	"github.com/jmoiron/sqlx"
	"gitlab.com/grchive/grchive-v2/shared/backend/audit"
	"gitlab.com/grchive/grchive-v2/shared/backend/comments"
	"gitlab.com/grchive/grchive-v2/shared/backend/controls"
	"gitlab.com/grchive/grchive-v2/shared/backend/engagements"
	"gitlab.com/grchive/grchive-v2/shared/backend/gl"
	"gitlab.com/grchive/grchive-v2/shared/backend/orgs"
	"gitlab.com/grchive/grchive-v2/shared/backend/risks"
	"gitlab.com/grchive/grchive-v2/shared/backend/roles"
	"gitlab.com/grchive/grchive-v2/shared/backend/sessions"
	"gitlab.com/grchive/grchive-v2/shared/backend/users"
	"gitlab.com/grchive/grchive-v2/shared/backend/utility"
)

type BackendInterface struct {
	db          *sqlx.DB
	Sessions    *sessions.SessionManager
	Users       *users.UserManager
	Orgs        *orgs.OrgManager
	Roles       *roles.RoleManager
	Engagements *engagements.EngagementManager
	Risks       *risks.RiskManager
	Controls    *controls.ControlManager
	Comments    *comments.CommentManager
	GL          *gl.GLManager
}

func CreateBackendInterface(db *sqlx.DB) *BackendInterface {
	return &BackendInterface{
		db:          db,
		Sessions:    sessions.CreateSessionManager(db),
		Users:       users.CreateUserManager(db),
		Orgs:        orgs.CreateOrgManager(db),
		Roles:       roles.CreateRoleManager(db),
		Engagements: engagements.CreateEngagementManager(db),
		Risks:       risks.CreateRiskManager(db),
		Controls:    controls.CreateControlManager(db),
		Comments:    comments.CreateCommentManager(db),
		GL:          gl.CreateGLManager(db),
	}
}

func (m *BackendInterface) WrapDatabaseTx(auditId audit.AuditTrailId, fns ...utility.TxHandler) error {
	return audit.WrapDatabaseAuditTx(m.db, auditId, fns...)
}
