package sessions

import (
	"github.com/jmoiron/sqlx"
)

func (m *SessionManager) DeleteSession(tx *sqlx.Tx, sessionId string) error {
	_, err := tx.Exec(`
		DELETE FROM user_sessions
		WHERE id = $1
	`, sessionId)
	return err
}
