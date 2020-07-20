package engagements

import (
	"github.com/jmoiron/sqlx"
	"gitlab.com/grchive/grchive-v2/shared/backend/roles"
	"time"
)

type Engagement struct {
	Id          int64          `db:"id"`
	Name        string         `db:"name"`
	Description string         `db:"description"`
	OrgId       int64          `db:"org_id"`
	CreatedTime time.Time      `db:"created_time"`
	StartTime   *time.Time     `db:"start_time"`
	EndTime     *time.Time     `db:"end_time"`
	IsClosed    bool           `db:"is_closed"`
	Roles       *[]*roles.Role `db:"roles"`
}

type EngagementManager struct {
	db *sqlx.DB
}

func CreateEngagementManager(db *sqlx.DB) *EngagementManager {
	return &EngagementManager{
		db: db,
	}
}
