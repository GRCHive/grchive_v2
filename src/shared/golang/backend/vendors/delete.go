package vendors

import (
	"github.com/jmoiron/sqlx"
)

func (m *VendorManager) DeleteVendor(tx *sqlx.Tx, id int64) error {
	_, err := tx.Exec(`
		DELETE FROM vendors
		WHERE id = $1
	`, id)
	return err
}

func (m *VendorManager) DeleteVendorProduct(tx *sqlx.Tx, id int64) error {
	_, err := tx.Exec(`
		DELETE FROM vendor_products
		WHERE id = $1
	`, id)
	return err
}
