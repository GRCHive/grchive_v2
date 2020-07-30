package vm

import (
	"github.com/jmoiron/sqlx"
	"gitlab.com/grchive/grchive-v2/shared/backend/machine_state"
)

type VirtualMachineManager struct {
	db *sqlx.DB
}

func CreateVirtualMachineManager(db *sqlx.DB) *VirtualMachineManager {
	return &VirtualMachineManager{
		db: db,
	}
}

type VirtualMachine struct {
	Id           int64                      `db:"id"`
	EngagementId int64                      `db:"engagement_id"`
	UniqueId     string                     `db:"unique_id"`
	Name         string                     `db:"name"`
	Description  string                     `db:"description"`
	VCpus        int32                      `db:"vcpus"`
	MemoryBytes  int64                      `db:"memory_bytes"`
	State        machine_state.MachineState `db:"state"`
	// This is the machine state this VM is deployed on.
	BaseStateId *int64 `db:"base_state_id"`
}
