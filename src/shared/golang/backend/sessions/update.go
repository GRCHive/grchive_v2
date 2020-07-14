package sessions

import (
	"github.com/jmoiron/sqlx"
)

func (m *SessionManager) UpdateSession(tx *sqlx.Tx, session *Session) error {
	_, err := tx.NamedExec(`
		UPDATE user_sessions
		SET access_token = :access_token,
			refresh_token = :refresh_token,
			session_expiration_time = :session_expiration_time,
			jwt_expiration_time = :jwt_expiration_time
		WHERE id = :id
	`, session)
	return err
}
