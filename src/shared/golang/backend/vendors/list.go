package vendors

func (m *VendorManager) ListVendorsForEngagement(engagementId int64) ([]*Vendor, error) {
	ret := []*Vendor{}
	err := m.db.Select(&ret, `
		SELECT *
		FROM vendors
		WHERE engagement_id = $1
		ORDER BY id DESC
	`, engagementId)
	return ret, err
}
