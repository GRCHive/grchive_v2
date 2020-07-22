package controls

import (
	"github.com/jmoiron/sqlx"
)

func (m *ControlManager) DeleteControl(tx *sqlx.Tx, controlId int64) error {
	_, err := tx.Exec(`
		DELETE FROM controls
		WHERE id = $1
	`, controlId)
	return err
}
