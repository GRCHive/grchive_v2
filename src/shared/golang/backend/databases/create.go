package databases

import (
	"github.com/jmoiron/sqlx"
)

func (m *DatabaseManager) CreateDatabase(tx *sqlx.Tx, db *Database) error {
	rows, err := tx.NamedQuery(`
		INSERT INTO databases (
			engagement_id,
			unique_id,
			name,
			description,
			type_id,
			other_type,
			version
		)
		VALUES (
			:engagement_id,
			:unique_id,
			:name,
			:description,
			:type_id,
			:other_type,
			:version
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
