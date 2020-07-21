package risks

import (
	"github.com/jmoiron/sqlx"
)

func (m *RiskManager) UpdateRisk(tx *sqlx.Tx, risk *Risk) error {
	_, err := tx.NamedExec(`
		UPDATE risks
		SET name = :name,
			human_id = :human_id,
			severity = :severity,
			severity_justification = :severity_justification
		WHERE id = :id
	`, risk)
	return err
}
