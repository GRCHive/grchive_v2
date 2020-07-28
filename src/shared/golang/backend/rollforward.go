package backend

import (
	"github.com/jmoiron/sqlx"
)

func (b *BackendInterface) rollForwardRisks(tx *sqlx.Tx, fromEngagementId int64, toEngagementId int64) error {
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
	return err
}

func (b *BackendInterface) rollForwardControls(tx *sqlx.Tx, fromEngagementId int64, toEngagementId int64) error {
	_, err := tx.Exec(`
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

func (b *BackendInterface) rollForwardGeneralLedger(tx *sqlx.Tx, fromEngagementId int64, toEngagementId int64) error {
	// Only need to handle copying over accounts because a new ledger gets created automatically via a trigger.
	// Need to do the copy in two passes.
	// 	1) Copy over accounts and replace engagement_id
	// 	2) For the accounts we copied, replace parent_account_id to be the new account using the (engagement_id, account_id) key.
	_, err := tx.Exec(`
		INSERT INTO gl_accounts (
			engagement_id,
			name,
			description,
			account_id,
			account_type,
			parent_account_id,
			financially_relevant,
			gl_id
		)
		SELECT
			$2,
			gl.name,
			gl.description,
			gl.account_id,
			gl.account_type,
			gl.parent_account_id,
			gl.financially_relevant,
			gl.gl_id
		FROM gl_accounts AS gl
		WHERE gl.engagement_id = $1
	`, fromEngagementId, toEngagementId)

	if err != nil {
		return err
	}

	_, err = tx.Exec(`
		WITH eng_from AS (
			SELECT *
			FROM gl_accounts
			WHERE engagement_id = $1
		), eng_to AS (
			SELECT *
			FROM gl_accounts
			WHERE engagement_id = $2
		), map_acc AS (
			SELECT f.id AS "from_id", t.id AS "to_id"
			FROM eng_from AS f
			INNER JOIN eng_to AS t
				ON t.account_id = f.account_id
		)
		UPDATE gl_accounts
		SET parent_account_id = map_acc.to_id
		FROM map_acc
		WHERE
			gl_accounts.engagement_id = $2 
			gl_accounts.parent_account_id = map_acc.from_id
	`, fromEngagementId, toEngagementId)
	return err
}

func (b *BackendInterface) RollForwardEngagement(tx *sqlx.Tx, fromEngagementId int64, toEngagementId int64) error {
	if err := b.rollForwardRisks(tx, fromEngagementId, toEngagementId); err != nil {
		return err
	}

	if err := b.rollForwardControls(tx, fromEngagementId, toEngagementId); err != nil {
		return err
	}

	if err := b.rollForwardGeneralLedger(tx, fromEngagementId, toEngagementId); err != nil {
		return err
	}

	return nil
}
