package vendors

func (m *VendorManager) GetNumVendorsForEngagement(engagementId int64) (int, error) {
	ret := 0
	err := m.db.Get(&ret, `
		SELECT COUNT(id)
		FROM vendors
		WHERE engagement_id = $1
	`, engagementId)
	return ret, err
}

func (m *VendorManager) GetNumVendorProductsForEngagement(engagementId int64) (int, error) {
	ret := 0
	err := m.db.Get(&ret, `
		SELECT COUNT(vp.id)
		FROM vendor_products AS vp
		INNER JOIN vendors AS v
			ON v.id = vp.vendor_id
		WHERE v.engagement_id = $1
	`, engagementId)
	return ret, err
}
