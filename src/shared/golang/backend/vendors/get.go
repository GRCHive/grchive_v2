package vendors

func (m *VendorManager) GetVendorFromId(id int64) (*Vendor, error) {
	vendor := Vendor{}
	err := m.db.Get(&vendor, `
		SELECT *
		FROM vendors
		WHERE id = $1
	`, id)
	return &vendor, err
}

func (m *VendorManager) GetVendorProductFromId(id int64) (*VendorProduct, error) {
	product := VendorProduct{}
	err := m.db.Get(&product, `
		SELECT *
		FROM vendor_products
		WHERE id = $1
	`, id)
	return &product, err
}
