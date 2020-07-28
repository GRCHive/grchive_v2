package inventory

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
)

func (m *InventoryManager) UpdateInventory(tx *sqlx.Tx, it InventoryType, inv interface{}) error {
	tblName, err := inventoryTypeToTableName(it)
	if err != nil {
		return err
	}

	baseInv := GetBaseInventory(inv)
	_, err = tx.NamedExec(`
		UPDATE inventory
		SET name = :name,
			description = :description,
			brand = :brand,
			model = :model
		WHERE id = :id
	`, baseInv)

	if err != nil {
		return err
	}

	switch it {
	case ITServer:
		_, err = tx.NamedExec(fmt.Sprintf(`
			UPDATE %s
			SET physical_location = :physical_location,
				operating_system = :operating_system,
				hypervisor = :hypervisor,
				static_external_ip = :static_external_ip
			WHERE id = :id
		`, tblName), inv)
	case ITDesktop:
		_, err = tx.NamedExec(fmt.Sprintf(`
			UPDATE %s
			SET physical_location = :physical_location,
				operating_system = :operating_system
			WHERE id = :id
		`, tblName), inv)
	case ITLaptop:
		_, err = tx.NamedExec(fmt.Sprintf(`
			UPDATE %s
			SET operating_system = :operating_system
			WHERE id = :id
		`, tblName), inv)
	case ITMobile:
		_, err = tx.NamedExec(fmt.Sprintf(`
			UPDATE %s
			SET operating_system = :operating_system
			WHERE id = :id
		`, tblName), inv)
	case ITEmbedded:
		_, err = tx.NamedExec(fmt.Sprintf(`
			UPDATE %s
			SET operating_system = :operating_system
			WHERE id = :id
		`, tblName), inv)
	default:
		return errors.New("Unsupported inventory for updating.")

	}
	return err
}
