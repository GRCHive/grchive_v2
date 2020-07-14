package session

import (
	"github.com/gorilla/sessions"
	gsess "gitlab.com/grchive/grchive-v2/shared/backend/sessions"
)

type SessionStore struct {
	store   *sessions.CookieStore
	backend *gsess.SessionManager
}

func CreateSessionStore(authKey string, encryptKey string, mgr *gsess.SessionManager) *SessionStore {
	store := SessionStore{
		store:   sessions.NewCookieStore([]byte(authKey), []byte(encryptKey)),
		backend: mgr,
	}
	return &store
}
