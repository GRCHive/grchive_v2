package inventory

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

func (m *InventoryManager) DeleteInventory(tx *sqlx.Tx, it InventoryType, id int64) error {
	tblName, err := inventoryTypeToTableName(it)
	if err != nil {
		return err
	}

	_, err = tx.Exec(fmt.Sprintf(`
		DELETE FROM %s
		WHERE id = $1
	`, tblName), id)
	return err
}
