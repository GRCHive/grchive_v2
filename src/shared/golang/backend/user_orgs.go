package backend

import (
	"gitlab.com/grchive/grchive-v2/shared/backend/orgs"
)

func (b *BackendInterface) GetUserOrgs(userId int64) ([]*orgs.Organization, error) {
	orgs := []*orgs.Organization{}
	err := b.db.Select(&orgs, `
		SELECT org.*
		FROM organizations AS org
		INNER JOIN user_orgs AS uo
			ON uo.org_id = org.id
		WHERE uo.user_id = $1
	`, userId)
	return orgs, err
}
