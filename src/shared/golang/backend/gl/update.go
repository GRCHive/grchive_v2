package gl

import (
	"github.com/jmoiron/sqlx"
)

func (m *GLManager) UpdateGLAccount(tx *sqlx.Tx, acc *GLAccount) error {
	_, err := tx.NamedExec(`
		UPDATE gl_accounts
		SET name = :name,
			description = :description,
			account_id = :account_id,
			account_type = :account_type,
			parent_account_id = :parent_account_id,
			financially_relevant = :financially_relevant
		WHERE id = :id
	`, acc)
	return err
}
