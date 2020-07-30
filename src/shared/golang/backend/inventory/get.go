package inventory

import (
	"fmt"
)

func (m *InventoryManager) GetInventory(it InventoryType, id int64) (interface{}, error) {
	tblName, err := inventoryTypeToTableName(it)
	if err != nil {
		return nil, err
	}

	typedColumns, err := inventoryToUniqueColumns(it)
	if err != nil {
		return nil, err
	}

	ret := CreateTypedInventory(it)
	err = m.db.Get(ret, fmt.Sprintf(`
		SELECT
			tbl.id AS "id",
			%s
			inv.id AS "inventory.id",
			inv.engagement_id AS "inventory.engagement_id",
			inv.unique_id AS "inventory.unique_id",
			inv.name AS "inventory.name",
			inv.description AS "inventory.description",
			inv.brand AS "inventory.brand",
			inv.model AS "inventory.model",
			state.id AS "state.id",
			state.engagement_id AS "state.engagement_id",
			state.unique_id AS "state.unique_id",
			state.hypervisor AS "state.hypervisor",
			state.operating_system AS "state.operating_system"
		FROM %s AS tbl
		INNER JOIN inventory AS inv
			ON tbl.inventory_id = inv.id
		INNER JOIN machine_state AS state
			ON state.id = tbl.state_id
		WHERE tbl.id = $1
	`, typedColumns, tblName), id)
	return ret, err
}
