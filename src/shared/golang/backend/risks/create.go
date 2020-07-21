package risks

import (
	"github.com/jmoiron/sqlx"
)

func (m *RiskManager) CreateRisk(tx *sqlx.Tx, risk *Risk) error {
	rows, err := tx.NamedQuery(`
		INSERT INTO risks (name, description, engagement_id, human_id, severity, severity_justification)
		VALUES (:name, :description, :engagement_id, :human_id, :severity, :severity_justification)
		RETURNING id
	`, risk)

	if err != nil {
		return err
	}
	defer rows.Close()
	rows.Next()
	return rows.Scan(&risk.Id)
}
