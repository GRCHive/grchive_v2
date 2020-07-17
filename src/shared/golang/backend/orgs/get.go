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

func (m *OrgManager) GetSuborgsFromId(id int64) ([]*Organization, error) {
	orgs := []*Organization{}
	err := m.db.Select(&orgs, `
		WITH RECURSIVE suborgs AS (
			SELECT *
			FROM organizations
			WHERE id = $1
			UNION
				SELECT o.*
				FROM organizations AS o
				INNER JOIN suborgs AS sub
					ON sub.id = o.parent_org_id
		)
		SELECT *
		FROM suborgs
	`, id)
	return orgs, err
}

func (m *OrgManager) GetParentOrgsFromId(id int64) ([]*Organization, error) {
	orgs := []*Organization{}
	err := m.db.Select(&orgs, `
		WITH RECURSIVE suborgs AS (
			SELECT *
			FROM organizations
			WHERE id = $1
			UNION
				SELECT o.*
				FROM organizations AS o
				INNER JOIN suborgs AS sub
					ON sub.parent_org_id = o.id
		)
		SELECT *
		FROM suborgs
	`, id)
	return orgs, err
}
