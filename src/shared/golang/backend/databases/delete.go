package databases

import (
	"github.com/jmoiron/sqlx"
)

func (m *DatabaseManager) DeleteDatabase(tx *sqlx.Tx, dbId int64) error {
	_, err := tx.Exec(`
		DELETE FROM databases
		WHERE id = $1
	`, dbId)
	return err
}
