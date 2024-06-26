package inventory

import (
	"errors"
	"fmt"
	"reflect"
)

func (m *InventoryManager) ListInventoryForEngagement(it InventoryType, engagementId int64) (interface{}, error) {
	var ret interface{}
	switch it {
	case ITServer:
		ret = make([]*InventoryServer, 0)
	case ITDesktop:
		ret = make([]*InventoryDesktop, 0)
	case ITLaptop:
		ret = make([]*InventoryLaptop, 0)
	case ITMobile:
		ret = make([]*InventoryMobile, 0)
	case ITEmbedded:
		ret = make([]*InventoryEmbedded, 0)
	default:
		return nil, errors.New("Not yet supported.")
	}

	typedColumns, err := inventoryToUniqueColumns(it)
	if err != nil {
		return nil, err
	}

	tblName, err := inventoryTypeToTableName(it)
	if err != nil {
		return nil, err
	}

	rows, err := m.db.Queryx(fmt.Sprintf(`
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
		WHERE inv.engagement_id = $1
		ORDER BY tbl.id DESC
	`, typedColumns, tblName), engagementId)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	refRet := reflect.ValueOf(ret)
	for rows.Next() {
		singleEle := CreateTypedInventory(it)
		err = rows.StructScan(singleEle)
		if err != nil {
			return nil, err
		}

		refRet = reflect.Append(refRet, reflect.ValueOf(singleEle))
	}

	return refRet.Interface(), nil
}
