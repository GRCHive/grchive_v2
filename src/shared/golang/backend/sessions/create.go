package sessions

import (
	"github.com/jmoiron/sqlx"
)

func (m *SessionManager) CreateSession(tx *sqlx.Tx, session *Session) error {
	_, err := tx.NamedExec(`
		INSERT INTO user_sessions (id, session_expiration_time, jwt_expiration_time, access_token, refresh_token, user_id)
		VALUES (:id, :session_expiration_time, :jwt_expiration_time, :access_token, :refresh_token, :user_id)
	`, session)
	return err
}
