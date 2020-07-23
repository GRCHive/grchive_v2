package gl

import (
	"github.com/jmoiron/sqlx"
)

func (m *GLManager) CreateGLAccount(tx *sqlx.Tx, acc *GLAccount) error {
	rows, err := tx.NamedQuery(`
		INSERT INTO gl_accounts (
			engagement_id,
			name,
			description,
			account_id,
			account_type,
			parent_account_id,
			financially_relevant
		)
		VALUES (
			:engagement_id,
			:name,
			:description,
			:account_id,
			:account_type,
			:parent_account_id,
			:financially_relevant
		)
		RETURNING id
	`, acc)

	if err != nil {
		return err
	}
	defer rows.Close()
	rows.Next()
	return rows.Scan(&acc.Id)
}
