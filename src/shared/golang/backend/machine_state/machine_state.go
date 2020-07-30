package machine_state

import (
	"github.com/jmoiron/sqlx"
)

type MachineStateManager struct {
	db *sqlx.DB
}

func CreateMachineStateManager(db *sqlx.DB) *MachineStateManager {
	return &MachineStateManager{
		db: db,
	}
}

type MachineState struct {
	Id              int64   `db:"id"`
	EngagementId    int64   `db:"engagement_id"`
	UniqueId        string  `db:"unique_id"`
	Hypervisor      *string `db:"hypervisor"`
	OperatingSystem *string `db:"operating_system"`
}
