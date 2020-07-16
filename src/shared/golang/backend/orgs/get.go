package orgs

func (m *OrgManager) GetOrgFromId(id int64) (*Organization, error) {
	org := Organization{}
	err := m.db.Get(&org, `
		SELECT *
		FROM organizations
		WHERE id = $1
	`, id)
	return &org, err
}
