package inventory

import (
	"errors"
	"github.com/jmoiron/sqlx"
	"gitlab.com/grchive/grchive-v2/shared/backend/machine_state"
	"reflect"
)

type InventoryType int64

const (
	ITServer InventoryType = iota
	ITDesktop
	ITLaptop
	ITMobile
	ITEmbedded
	ITNull
)

type InventoryManager struct {
	db *sqlx.DB
}

func CreateInventoryManager(db *sqlx.DB) *InventoryManager {
	return &InventoryManager{
		db: db,
	}
}

type Inventory struct {
	Id           int64  `db:"id"`
	EngagementId int64  `db:"engagement_id"`
	UniqueId     string `db:"unique_id"`
	Name         string `db:"name"`
	Description  string `db:"description"`
	Brand        string `db:"brand"`
	Model        string `db:"model"`
}

type InventoryServer struct {
	Id               int64                      `db:"id"`
	Inventory        Inventory                  `db:"inventory"`
	PhysicalLocation string                     `db:"physical_location"`
	State            machine_state.MachineState `db:"state"`
	StaticExternalIp *string                    `db:"static_external_ip"`
}

type InventoryDesktop struct {
	Id               int64                      `db:"id"`
	Inventory        Inventory                  `db:"inventory"`
	State            machine_state.MachineState `db:"state"`
	PhysicalLocation string                     `db:"physical_location"`
}

type InventoryLaptop struct {
	Id        int64                      `db:"id"`
	Inventory Inventory                  `db:"inventory"`
	State     machine_state.MachineState `db:"state"`
}

type InventoryMobile struct {
	Id        int64                      `db:"id"`
	Inventory Inventory                  `db:"inventory"`
	State     machine_state.MachineState `db:"state"`
}

type InventoryEmbedded struct {
	Id        int64                      `db:"id"`
	Inventory Inventory                  `db:"inventory"`
	State     machine_state.MachineState `db:"state"`
}

func inventoryTypeToTableName(it InventoryType) (string, error) {
	switch it {
	case ITServer:
		return "inventory_servers", nil
	case ITDesktop:
		return "inventory_desktops", nil
	case ITLaptop:
		return "inventory_laptops", nil
	case ITMobile:
		return "inventory_mobile", nil
	case ITEmbedded:
		return "inventory_embedded", nil
	default:
		return "", errors.New("Inventory type table not yet supported.")
	}
}

func inventoryToUniqueColumns(it InventoryType) (string, error) {
	switch it {
	case ITServer:
		return `
			tbl.physical_location AS "physical_location",
			text(tbl.static_external_ip) AS "static_external_ip",
		`, nil
	case ITDesktop:
		return `
			tbl.physical_location AS "physical_location",
		`, nil
	case ITLaptop:
		fallthrough
	case ITMobile:
		fallthrough
	case ITEmbedded:
		return ``, nil
	default:
		return "", errors.New("Inventory type columns not yet supported.")
	}
}

func CreateTypedInventory(it InventoryType) interface{} {
	switch it {
	case ITServer:
		return &InventoryServer{}
	case ITDesktop:
		return &InventoryDesktop{}
	case ITLaptop:
		return &InventoryLaptop{}
	case ITMobile:
		return &InventoryMobile{}
	case ITEmbedded:
		return &InventoryEmbedded{}
	default:
		return nil
	}
}

func SetInventoryEngagementId(inv interface{}, id int64) {
	ref := reflect.ValueOf(inv).Elem()
	baseInventory := ref.FieldByName("Inventory")
	refId := baseInventory.FieldByName("EngagementId")
	refId.Set(reflect.ValueOf(id))
}

func GetInventoryEngagementId(inv interface{}) int64 {
	ref := reflect.ValueOf(inv).Elem()
	baseInventory := ref.FieldByName("Inventory")
	refId := baseInventory.FieldByName("EngagementId")
	return refId.Int()
}

func GetBaseInventory(inv interface{}) *Inventory {
	ref := reflect.ValueOf(inv).Elem()
	baseInventory := ref.FieldByName("Inventory")
	return baseInventory.Addr().Interface().(*Inventory)
}

func GetInventoryId(inv interface{}) int64 {
	ref := reflect.ValueOf(inv).Elem()
	return ref.FieldByName("Id").Int()
}

func GetMachineState(inv interface{}) *machine_state.MachineState {
	ref := reflect.ValueOf(inv).Elem()
	return ref.FieldByName("State").Addr().Interface().(*machine_state.MachineState)
}

func SetInventoryId(inv interface{}, id int64) {
	ref := reflect.ValueOf(inv).Elem()
	field := ref.FieldByName("Id")
	field.Set(reflect.ValueOf(id))
}
