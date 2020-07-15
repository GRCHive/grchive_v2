package orgs

import (
	"github.com/jmoiron/sqlx"
)

type Organization struct {
	Id          int64  `db:"id"`
	Name        string `db:"name"`
	Description string `db:"description"`
	ParentOrgId *int64 `db:"parent_org_id"`
	OwnerUserId int64  `db:"owner_user_id"`
}

type OrgManager struct {
	db *sqlx.DB
}

func CreateOrgManager(db *sqlx.DB) *OrgManager {
	return &OrgManager{
		db: db,
	}
}
