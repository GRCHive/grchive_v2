package session

import (
	"github.com/gorilla/sessions"
)

type SessionStore struct {
	store *sessions.CookieStore
}

func CreateSessionStore(authKey string, encryptKey string) *SessionStore {
	store := SessionStore{
		store: sessions.NewCookieStore([]byte(authKey), []byte(encryptKey)),
	}
	return &store
}
