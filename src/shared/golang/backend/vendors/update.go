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

func (m *VendorManager) UpdateVendorProduct(tx *sqlx.Tx, product *VendorProduct) error {
	_, err := tx.NamedExec(`
		UPDATE vendor_products
		SET name = :name,
			description = :description,
			url = :url
		WHERE id = :id
	`, product)
	return err
}
