package vendors

import (
	"github.com/jmoiron/sqlx"
)

func (m *VendorManager) CreateVendor(tx *sqlx.Tx, vendor *Vendor) error {
	rows, err := tx.NamedQuery(`
		INSERT INTO vendors (engagement_id, name, description, url)
		VALUES (:engagement_id, :name, :description, :url)
		RETURNING id
	`, vendor)
	if err != nil {
		return err
	}
	defer rows.Close()
	rows.Next()
	return rows.Scan(&vendor.Id)
}
