package systems

import (
	"github.com/jmoiron/sqlx"
)

func (m *SystemManager) DeleteSystem(tx *sqlx.Tx, dbId int64) error {
	_, err := tx.Exec(`
		DELETE FROM systems
		WHERE id = $1
	`, dbId)
	return err
}
