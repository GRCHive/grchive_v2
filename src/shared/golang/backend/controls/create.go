package controls

import (
	"github.com/jmoiron/sqlx"
)

func (m *ControlManager) CreateControl(tx *sqlx.Tx, control *Control) error {
	rows, err := tx.NamedQuery(`
		INSERT INTO controls (
			engagement_id,
			name, 
			human_id,
			description,
			likelihood,
			likelihood_justification,
			control_type,
			owner_id,
			is_manual,
			freq_type,
			freq_other,
			freq_interval
		)
		VALUES (
			:engagement_id,
			:name, 
			:human_id,
			:description,
			:likelihood,
			:likelihood_justification,
			:control_type,
			:owner_id,
			:is_manual,
			:freq_type,
			:freq_other,
			:freq_interval
		)
		RETURNING id
	`, control)

	if err != nil {
		return err
	}
	defer rows.Close()
	rows.Next()
	return rows.Scan(&control.Id)
}
