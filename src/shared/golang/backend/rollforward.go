package backend

import (
	"github.com/jmoiron/sqlx"
)

func (b *BackendInterface) RollForwardEngagement(tx *sqlx.Tx, fromEngagementId int64, toEngagementId int64) error {
	_, err := tx.Exec(`
		INSERT INTO risks (
			engagement_id,
			name,
			human_id,
			severity,
			severity_justification,
			description
		)
		SELECT
			$2,
			r.name,
			r.human_id,
			r.severity,
			r.severity_justification,
			r.description
		FROM risks AS r
		WHERE r.engagement_id = $1
	`, fromEngagementId, toEngagementId)

	if err != nil {
		return err
	}

	_, err = tx.Exec(`
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
		SELECT
			$2,
			c.name,
			c.human_id,
			c.description,
			c.likelihood,
			c.likelihood_justification,
			c.control_type,
			c.owner_id,
			c.is_manual,
			c.freq_type,
			c.freq_other,
			c.freq_interval
		FROM controls AS c
		WHERE c.engagement_id = $1
	`, fromEngagementId, toEngagementId)
	return err
}
