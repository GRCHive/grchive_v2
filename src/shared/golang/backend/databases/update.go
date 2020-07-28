package databases

import (
	"github.com/jmoiron/sqlx"
)

func (m *DatabaseManager) UpdateDatabase(tx *sqlx.Tx, db *Database) error {
	_, err := tx.NamedExec(`
		UPDATE databases
		SET name = :name,
			description = :description,
			type_id = :type_id,
			other_type = :other_type,
			version = :version
		WHERE id = :id
	`, db)
	return err
}
