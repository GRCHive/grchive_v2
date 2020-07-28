package backend

import (
	"fmt"
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
			gl_accounts.engagement_id = $2 AND
			gl_accounts.parent_account_id = map_acc.from_id
	`, fromEngagementId, toEngagementId)
	return err
}

func (b *BackendInterface) rollForwardVendors(tx *sqlx.Tx, fromEngagementId int64, toEngagementId int64) error {
	// First copy over vendors then copy over products and do the vendor_id replacement during the product copy.
	_, err := tx.Exec(`
		INSERT INTO vendors (
			engagement_id,
			name,
			description,
			url
		)
		SELECT
			$2,
			v.name,
			v.description,
			v.url
		FROM vendors AS v
		WHERE v.engagement_id = $1
	`, fromEngagementId, toEngagementId)

	if err != nil {
		return err
	}

	_, err = tx.Exec(`
		WITH vendors_from AS (
			SELECT *
			FROM vendors
			WHERE engagement_id = $1
		), vendors_to AS (
			SELECT *
			FROM vendors
			WHERE engagement_id = $2
		), map_vendors AS (
			SELECT f.id AS "from_id", t.id AS "to_id"
			FROM vendors_from AS f
			INNER JOIN vendors_to AS t
				ON t.name = f.name
		)
		INSERT INTO vendor_products (
			vendor_id,
			name,
			description,
			url
		)
		SELECT
			mv.to_id,
			vp.name,
			vp.description,
			vp.url
		FROM vendor_products AS vp
		INNER JOIN vendors AS v
			ON v.id = vp.vendor_id
		INNER JOIN map_vendors AS mv
			ON mv.from_id = v.id
		WHERE v.engagement_id = $1
	`, fromEngagementId, toEngagementId)
	return err
}

func (b *BackendInterface) rollForwardInventory(tx *sqlx.Tx, fromEngagementId int64, toEngagementId int64) error {
	// Roll-forward base inventory data (inventory) first and then roll forward the specific types of inventory next and doing
	// inventory_id replacement too.
	_, err := tx.Exec(`
		INSERT INTO inventory (
			engagement_id,
			unique_id,
			name,
			description,
			brand,
			model
		)
		SELECT
			$2,
			inv.unique_id,
			inv.name,
			inv.description,
			inv.brand,
			inv.model
		FROM inventory AS inv
		WHERE inv.engagement_id = $1
	`, fromEngagementId, toEngagementId)

	if err != nil {
		return err
	}

	commonCte := `
		WITH inventory_from AS (
			SELECT *
			FROM inventory
			WHERE engagement_id = $1
		), inventory_to AS (
			SELECT *
			FROM inventory
			WHERE engagement_id = $2
		), map_inventory AS (
			SELECT f.id AS "from_id", t.id AS "to_id"
			FROM inventory_from AS f
			INNER JOIN inventory_to AS t
				ON t.name = f.name
		)
	`
	_, err = tx.Exec(fmt.Sprintf(`
		%s
		INSERT INTO inventory_servers (
			inventory_id,
			physical_location,
			operating_system,
			hypervisor,
			static_external_ip
		)
		SELECT
			map.to_id,
			inv.physical_location,
			inv.operating_system,
			inv.hypervisor,
			inv.static_external_ip
		FROM inventory_servers AS inv
		INNER JOIN map_inventory AS map
			ON map.from_id = inv.inventory_id
	`, commonCte), fromEngagementId, toEngagementId)

	if err != nil {
		return err
	}

	_, err = tx.Exec(fmt.Sprintf(`
		%s
		INSERT INTO inventory_desktops (
			inventory_id,
			physical_location,
			operating_system
		)
		SELECT
			map.to_id,
			inv.physical_location,
			inv.operating_system
		FROM inventory_desktops AS inv
		INNER JOIN map_inventory AS map
			ON map.from_id = inv.inventory_id
	`, commonCte), fromEngagementId, toEngagementId)

	if err != nil {
		return err
	}

	_, err = tx.Exec(fmt.Sprintf(`
		%s
		INSERT INTO inventory_laptops (
			inventory_id,
			operating_system
		)
		SELECT
			map.to_id,
			inv.operating_system
		FROM inventory_laptops AS inv
		INNER JOIN map_inventory AS map
			ON map.from_id = inv.inventory_id
	`, commonCte), fromEngagementId, toEngagementId)

	if err != nil {
		return err
	}

	_, err = tx.Exec(fmt.Sprintf(`
		%s
		INSERT INTO inventory_mobile (
			inventory_id,
			operating_system
		)
		SELECT
			map.to_id,
			inv.operating_system
		FROM inventory_mobile AS inv
		INNER JOIN map_inventory AS map
			ON map.from_id = inv.inventory_id
	`, commonCte), fromEngagementId, toEngagementId)

	if err != nil {
		return err
	}

	_, err = tx.Exec(fmt.Sprintf(`
		%s
		INSERT INTO inventory_embedded (
			inventory_id,
			operating_system
		)
		SELECT
			map.to_id,
			inv.operating_system
		FROM inventory_embedded AS inv
		INNER JOIN map_inventory AS map
			ON map.from_id = inv.inventory_id
	`, commonCte), fromEngagementId, toEngagementId)

	if err != nil {
		return err
	}

	return nil
}

func (b *BackendInterface) rollForwardDatabases(tx *sqlx.Tx, fromEngagementId int64, toEngagementId int64) error {
	_, err := tx.Exec(`
		INSERT INTO databases (
			engagement_id,
			unique_id,
			name,
			description,
			type_id,
			other_type,
			version
		)
		SELECT
			$2,
			db.unique_id,
			db.name,
			db.description,
			db.type_id,
			db.other_type,
			db.version
		FROM databases AS db
		WHERE db.engagement_id = $1
	`, fromEngagementId, toEngagementId)
	return err
}

func (b *BackendInterface) rollForwardSystems(tx *sqlx.Tx, fromEngagementId int64, toEngagementId int64) error {
	_, err := tx.Exec(`
		INSERT INTO systems (
			engagement_id,
			unique_id,
			name,
			description,
			purpose
		)
		SELECT
			$2,
			sys.unique_id,
			sys.name,
			sys.description,
			sys.purpose
		FROM systems AS sys
		WHERE sys.engagement_id = $1
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

	if err := b.rollForwardVendors(tx, fromEngagementId, toEngagementId); err != nil {
		return err
	}

	if err := b.rollForwardInventory(tx, fromEngagementId, toEngagementId); err != nil {
		return err
	}

	if err := b.rollForwardDatabases(tx, fromEngagementId, toEngagementId); err != nil {
		return err
	}

	if err := b.rollForwardSystems(tx, fromEngagementId, toEngagementId); err != nil {
		return err
	}

	return nil
}
