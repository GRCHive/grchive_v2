package sessions

import (
	"github.com/jmoiron/sqlx"
)

func (m *SessionManager) UpdateSession(tx *sqlx.Tx, session *Session) error {
	_, err := tx.NamedExec(`
	`, session)
	return err
}
