package systems

import (
	"github.com/jmoiron/sqlx"
)

func (m *SystemManager) CreateSystem(tx *sqlx.Tx, db *System) error {
	rows, err := tx.NamedQuery(`
		INSERT INTO systems (
			engagement_id,
			unique_id,
			name,
			description,
			purpose
		)
		VALUES (
			:engagement_id,
			:unique_id,
			:name,
			:description,
			:purpose
		)
		RETURNING id
	`, db)
	if err != nil {
		return err
	}
	defer rows.Close()
	rows.Next()
	return rows.Scan(&db.Id)
}
