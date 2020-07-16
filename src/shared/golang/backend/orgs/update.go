package orgs

import (
	"github.com/jmoiron/sqlx"
)

func (m *OrgManager) UpdateOrg(tx *sqlx.Tx, org *Organization) error {
	_, err := tx.NamedExec(`
		UPDATE organizations
		SET name = :name,
			description = :description
		WHERE id = :id
	`, org)
	return err
}
