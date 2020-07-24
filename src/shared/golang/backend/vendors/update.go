package vendors

import (
	"github.com/jmoiron/sqlx"
)

func (m *VendorManager) UpdateVendor(tx *sqlx.Tx, vendor *Vendor) error {
	_, err := tx.NamedExec(`
		UPDATE vendors
		SET name = :name,
			description = :description,
			url = :url
		WHERE id = :id
	`, vendor)
	return err
}
