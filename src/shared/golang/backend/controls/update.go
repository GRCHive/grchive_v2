package controls

import (
	"github.com/jmoiron/sqlx"
)

func (m *ControlManager) UpdateControl(tx *sqlx.Tx, control *Control) error {
	_, err := tx.NamedExec(`
		UPDATE controls
		SET name = :name,
			human_id = :human_id,
			likelihood = :likelihood,
			likelihood_justification = :likelihood_justification,
			control_type = :control_type,
			owner_id = :owner_id,
			is_manual = :is_manual,
			freq_type = :freq_type,
			freq_other = :freq_other,
			freq_interval = :freq_interval,
			description = :description
		WHERE id = :id
	`, control)
	return err
}
