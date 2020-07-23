package gl

import (
	"github.com/jmoiron/sqlx"
)

func (m *GLManager) DeleteGLAccount(tx *sqlx.Tx, accId int64) error {
	_, err := tx.Exec(`
		DELETE FROM gl_accounts
		WHERE id = $1
	`, accId)
	return err
}
