package systems

import (
	"github.com/jmoiron/sqlx"
)

func (m *SystemManager) UpdateSystem(tx *sqlx.Tx, db *System) error {
	_, err := tx.NamedExec(`
		UPDATE systems
		SET name = :name,
			description = :description,
			purpose = :purpose
		WHERE id = :id
	`, db)
	return err
}
