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
