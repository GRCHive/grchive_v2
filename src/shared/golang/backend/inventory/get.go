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
			inv.model AS "inventory.model"
		FROM %s AS tbl
		INNER JOIN inventory AS inv
			ON tbl.inventory_id = inv.id
		WHERE tbl.id = $1
	`, typedColumns, tblName), id)
	return ret, err
}
