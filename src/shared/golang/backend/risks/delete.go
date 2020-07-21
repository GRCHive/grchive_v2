package risks

import (
	"github.com/jmoiron/sqlx"
)

func (m *RiskManager) DeleteRisk(tx *sqlx.Tx, riskId int64) error {
	_, err := tx.Exec(`
		DELETE FROM risks
		WHERE id = $1
	`, riskId)
	return err
}
