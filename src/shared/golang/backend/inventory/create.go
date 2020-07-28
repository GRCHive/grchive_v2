package inventory

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"reflect"
)

func (m *InventoryManager) CreateBaseInventory(tx *sqlx.Tx, inv *Inventory) error {
	rows, err := tx.NamedQuery(`
		INSERT INTO inventory (
			engagement_id,
			unique_id,
			name,
			description,
			brand,
			model
		)
		VALUES (
			:engagement_id,
			:unique_id,
			:name,
			:description,
			:brand,
			:model
		)
		RETURNING id
	`, inv)
	if err != nil {
		return err
	}
	defer rows.Close()
	rows.Next()
	return rows.Scan(&inv.Id)
}

func (m *InventoryManager) CreateInventory(tx *sqlx.Tx, it InventoryType, inv interface{}) error {
	refInv := reflect.ValueOf(inv).Elem()
	baseInv := refInv.FieldByName("Inventory")

	typedBase := baseInv.Addr().Interface().(*Inventory)
	err := m.CreateBaseInventory(tx, typedBase)
	if err != nil {
		return err
	}

	tblName, err := inventoryTypeToTableName(it)
	if err != nil {
		return err
	}

	var rows *sqlx.Rows
	err = nil
	switch it {
	case ITServer:
		rows, err = tx.NamedQuery(fmt.Sprintf(`
			INSERT INTO %s (
				inventory_id,
				physical_location,
				operating_system,
				hypervisor,
				static_external_ip
			)
			VALUES (
				:inventory.id,
				:physical_location,
				:operating_system,
				:hypervisor,
				:static_external_ip
			)
			RETURNING id
		`, tblName), inv)
	case ITDesktop:
		rows, err = tx.NamedQuery(fmt.Sprintf(`
			INSERT INTO %s (
				inventory_id,
				physical_location,
				operating_system
			)
			VALUES (
				:inventory.id,
				:physical_location,
				:operating_system
			)
			RETURNING id
		`, tblName), inv)
	case ITLaptop:
		rows, err = tx.NamedQuery(fmt.Sprintf(`
			INSERT INTO %s (
				inventory_id,
				operating_system
			)
			VALUES (
				:inventory.id,
				:operating_system
			)
			RETURNING id
		`, tblName), inv)
	case ITMobile:
		rows, err = tx.NamedQuery(fmt.Sprintf(`
			INSERT INTO %s (
				inventory_id,
				operating_system
			)
			VALUES (
				:inventory.id,
				:operating_system
			)
			RETURNING id
		`, tblName), inv)
	case ITEmbedded:
		rows, err = tx.NamedQuery(fmt.Sprintf(`
			INSERT INTO %s (
				inventory_id,
				operating_system
			)
			VALUES (
				:inventory.id,
				:operating_system
			)
			RETURNING id
		`, tblName), inv)
	default:
		err = errors.New("Unsupported inventory type for creation.")
	}

	if err != nil {
		return err
	}

	defer rows.Close()
	rows.Next()
	return rows.Scan(refInv.FieldByName("Id").Addr().Interface())
}
