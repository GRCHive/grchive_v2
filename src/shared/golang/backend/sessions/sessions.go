package sessions

import (
	"github.com/jmoiron/sqlx"
	"time"
)

type SessionManager struct {
	db *sqlx.DB
}

func CreateSessionManager(db *sqlx.DB) *SessionManager {
	return &SessionManager{
		db: db,
	}
}

type Session struct {
	Id                string    `db:"id"`
	SessionExpiration time.Time `db:"session_expiration_time"`
	JwtExpiration     time.Time `db:"jwt_expiration_time"`
	AccessToken       string    `db:"access_token"`
	RefreshToken      string    `db:"refresh_token"`
	UserId            int64     `db:"user_id"`
}
