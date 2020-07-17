package engagements

import (
	"github.com/jmoiron/sqlx"
)

func (m *EngagementManager) CreateEngagement(tx *sqlx.Tx, engagement *Engagement) error {
	rows, err := tx.NamedQuery(`
		INSERT INTO engagements (name, description, org_id, created_time, start_time, end_time)
		VALUES (:name, :description, :org_id, NOW(), :start_time, :end_time)
		RETURNING id
	`, engagement)

	if err != nil {
		return err
	}
	defer rows.Close()
	rows.Next()
	return rows.Scan(&engagement.Id)
}
