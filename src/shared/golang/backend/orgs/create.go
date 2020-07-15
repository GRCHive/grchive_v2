package orgs

import (
	"github.com/jmoiron/sqlx"
)

func (m *OrgManager) CreateOrg(tx *sqlx.Tx, org *Organization) error {
	rows, err := tx.NamedQuery(`
		INSERT INTO organizations (name, description, parent_org_id, owner_user_id)
		VALUES (:name, :description, :parent_org_id, :owner_user_id)
		RETURNING id
	`, org)

	if err != nil {
		return err
	}
	defer rows.Close()
	rows.Next()
	return rows.Scan(&org.Id)
}
