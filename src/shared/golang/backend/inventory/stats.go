package inventory

import (
	"fmt"
)

func (m *InventoryManager) GetNumInventoryTypeForEngagement(typ InventoryType, engagementId int64) (int, error) {
	tblName, err := inventoryTypeToTableName(typ)
	if err != nil {
		return 0, err
	}

	ret := 0
	err = m.db.Get(&ret, fmt.Sprintf(`
		SELECT COUNT(tbl.id)
		FROM %s AS tbl
		INNER JOIN inventory AS inv
			ON inv.id = tbl.inventory_id
		WHERE inv.engagement_id = $1
	`, tblName), engagementId)
	return ret, err
}

func (m *InventoryManager) GetNumInventoryForEngagement(engagementId int64) (map[InventoryType]int, error) {
	var err error
	ret := map[InventoryType]int{}

	ret[ITServer], err = m.GetNumInventoryTypeForEngagement(ITServer, engagementId)
	if err != nil {
		return nil, err
	}

	ret[ITDesktop], err = m.GetNumInventoryTypeForEngagement(ITDesktop, engagementId)
	if err != nil {
		return nil, err
	}

	ret[ITLaptop], err = m.GetNumInventoryTypeForEngagement(ITLaptop, engagementId)
	if err != nil {
		return nil, err
	}

	ret[ITMobile], err = m.GetNumInventoryTypeForEngagement(ITMobile, engagementId)
	if err != nil {
		return nil, err
	}

	ret[ITEmbedded], err = m.GetNumInventoryTypeForEngagement(ITEmbedded, engagementId)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
