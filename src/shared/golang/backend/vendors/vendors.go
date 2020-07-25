package vendors

import (
	"github.com/jmoiron/sqlx"
)

type Vendor struct {
	Id           int64  `db:"id"`
	EngagementId int64  `db:"engagement_id"`
	Name         string `db:"name"`
	Description  string `db:"description"`
	Url          string `db:"url"`
}

type VendorProduct struct {
	Id          int64  `db:"id"`
	VendorId    int64  `db:"vendor_id"`
	Name        string `db:"name"`
	Description string `db:"description"`
	Url         string `db:"url"`
}

type VendorManager struct {
	db *sqlx.DB
}

func CreateVendorManager(db *sqlx.DB) *VendorManager {
	return &VendorManager{
		db: db,
	}
}
