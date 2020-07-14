package sessions

import (
	"github.com/jmoiron/sqlx"
)

func (m *SessionManager) DeleteSession(tx *sqlx.Tx, sessionId string) error {
	_, err := tx.Exec(`
	`, sessionId)
	return err
}
