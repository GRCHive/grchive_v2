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

func (m *EngagementManager) LinkEngagementToRoles(tx *sqlx.Tx, engagement *Engagement) error {
	if engagement.Roles == nil {
		return nil
	}

	for _, r := range *engagement.Roles {
		_, err := tx.Exec(`
			INSERT INTO engagement_roles (engagement_id, role_id, org_id)
			VALUES ($1, $2, $3)
		`, engagement.Id, r.Id, engagement.OrgId)

		if err != nil {
			return err
		}
	}

	return nil
}
