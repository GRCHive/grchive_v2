package systems

import (
	"github.com/jmoiron/sqlx"
)

type System struct {
	Id           int64  `db:"id"`
	EngagementId int64  `db:"engagement_id"`
	UniqueId     string `db:"unique_id"`
	Name         string `db:"name"`
	Description  string `db:"description"`
	Purpose      string `db:"purpose"`
}

type SystemManager struct {
	db *sqlx.DB
}

func CreateSystemManager(db *sqlx.DB) *SystemManager {
	return &SystemManager{
		db: db,
	}
}
