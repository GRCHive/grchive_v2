package backend

import (
	"gitlab.com/grchive/grchive-v2/shared/backend/sessions"
	"gitlab.com/grchive/grchive-v2/shared/backend/users"
)

type BackendInterface struct {
	Sessions *sessions.SessionManager
	Users    *users.UserManager
}
