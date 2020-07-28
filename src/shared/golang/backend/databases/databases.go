package databases

import (
	"github.com/jmoiron/sqlx"
)

type Database struct {
	Id           int64   `db:"id"`
	EngagementId int64   `db:"engagement_id"`
	UniqueId     string  `db:"unique_id"`
	Name         string  `db:"name"`
	Description  string  `db:"description"`
	TypeId       int32   `db:"type_id"`
	OtherType    *string `db:"other_type"`
	Version      string  `db:"version"`
}

type DatabaseManager struct {
	db *sqlx.DB
}

func CreateDatabaseManager(db *sqlx.DB) *DatabaseManager {
	return &DatabaseManager{
		db: db,
	}
}
