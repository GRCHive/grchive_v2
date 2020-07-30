package machine_state

import (
	"github.com/jmoiron/sqlx"
)

func (m *MachineStateManager) UpdateMachineState(tx *sqlx.Tx, state *MachineState) error {
	_, err := tx.NamedExec(`
		UPDATE machine_state
		SET hypervisor = :hypervisor,
			operating_system = :operating_system
		WHERE id = :id
	`, state)
	return err
}
