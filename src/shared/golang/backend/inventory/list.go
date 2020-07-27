package inventory

import (
	"errors"
	"fmt"
	"reflect"
)

func (m *InventoryManager) ListInventoryForEngagement(it InventoryType, engagementId int64) (interface{}, error) {
	var ret interface{}
	var typedColumns string
	singleEle := CreateTypedInventory(it)
	switch it {
	case ITServer:
		ret = make([]*InventoryServer, 0)
		typedColumns = `
			tbl.physical_location AS "physical_location",
			tbl.operating_system AS "operating_system",
			tbl.hypervisor AS "hypervisor",
			text(tbl.static_external_ip) AS "static_external_ip",
		`
	default:
		return nil, errors.New("Not yet supported.")
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
			inv.model AS "inventory.model"
		FROM %s AS tbl
		INNER JOIN inventory AS inv
			ON tbl.inventory_id = inv.id
		WHERE inv.engagement_id = $1
	`, typedColumns, tblName), engagementId)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	refRet := reflect.ValueOf(ret)
	for rows.Next() {
		err = rows.StructScan(singleEle)
		if err != nil {
			return nil, err
		}

		refRet = reflect.Append(refRet, reflect.ValueOf(singleEle))
	}

	return refRet.Interface(), nil
}
