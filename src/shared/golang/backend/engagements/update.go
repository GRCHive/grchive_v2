package engagements

import (
	"github.com/jmoiron/sqlx"
)

func (m *EngagementManager) UpdateEngagement(tx *sqlx.Tx, engagement *Engagement) error {
	_, err := tx.NamedExec(`
		UPDATE engagements
		SET name = :name,
			description = :description,
			start_time = :start_time,
			end_time = :end_time
		WHERE id = :id
	`, engagement)
	return err
}
