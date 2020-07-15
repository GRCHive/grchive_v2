package session

import (
	"github.com/gorilla/sessions"
	"gitlab.com/grchive/grchive-v2/shared/backend"
)

type SessionStore struct {
	SecureCookies bool

	store *sessions.CookieStore
	itf   *backend.BackendInterface
}

func CreateSessionStore(authKey string, encryptKey string, itf *backend.BackendInterface) *SessionStore {
	store := SessionStore{
		store: sessions.NewCookieStore([]byte(authKey), []byte(encryptKey)),
		itf:   itf,
	}
	return &store
}
