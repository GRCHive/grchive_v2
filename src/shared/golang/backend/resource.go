package backend

import (
	"errors"
	"fmt"
	"gitlab.com/grchive/grchive-v2/shared/backend/controls"
	"gitlab.com/grchive/grchive-v2/shared/backend/inventory"
	"gitlab.com/grchive/grchive-v2/shared/backend/risks"
	"gitlab.com/grchive/grchive-v2/shared/backend/roles"
	"reflect"
	"strconv"
)

type ResourceIdentifier int64

func (r ResourceIdentifier) Str() string {
	return fmt.Sprintf("resource-%d", r)
}

const (
	RIUser ResourceIdentifier = iota
	RIOrganization
	RIEngagement
	RIRisk
	RIControl
	RIGeneralLedger
	RIGLAccount
	RIVendor
	RIVendorProduct
	RIInventoryServer
	RIInventoryDesktop
	RIInventoryLaptop
	RIInventoryMobile
	RIInventoryEmbedded
	RIDatabase
	RISystem
)

type GenericResourceIdentifier struct {
	OrgId        int64  `json:"orgId"`
	EngagementId int64  `json:"engagementId"`
	RiskId       *int64 `json:"riskId"`
	ControlId    *int64 `json:"controlId"`
	GlAccountId  *int64 `json:"glAccountId"`
	VendorId     *int64 `json:"vendorId"`
	ServerId     *int64 `json:"serverId"`
	DesktopId    *int64 `json:"desktopId"`
	LaptopId     *int64 `json:"laptopId"`
	MobileId     *int64 `json:"mobileId"`
	EmbeddedId   *int64 `json:"embeddedId"`
	DatabaseId   *int64 `json:"databaseId"`
	SystemId     *int64 `json:"systemId"`
}

func ResourceToInventoryType(resource ResourceIdentifier) inventory.InventoryType {
	switch resource {
	case RIInventoryServer:
		return inventory.ITServer
	case RIInventoryDesktop:
		return inventory.ITDesktop
	case RIInventoryLaptop:
		return inventory.ITLaptop
	case RIInventoryMobile:
		return inventory.ITMobile
	case RIInventoryEmbedded:
		return inventory.ITEmbedded
	default:
		return inventory.ITNull
	}
}

func CreateResource(resource ResourceIdentifier) (interface{}, error) {
	switch resource {
	case RIControl:
		return &controls.Control{}, nil
	case RIRisk:
		return &risks.Risk{}, nil
	}
	return "", errors.New("Unsupported resource for generic createion.")
}

func resourceToTableName(resource ResourceIdentifier) (string, error) {
	switch resource {
	case RIControl:
		return "controls", nil
	case RIRisk:
		return "risks", nil
	}
	return "", errors.New("Unsupported resource for table name extraction.")
}

func resourceToFKDatabaseId(resource ResourceIdentifier) (string, error) {
	switch resource {
	case RIRisk:
		return "risk_id", nil
	case RIControl:
		return "control_id", nil
	}
	return "", errors.New("Unsupported resource for foreign key ID name extraction.")
}

func ResourceToEndpointName(resource ResourceIdentifier) string {
	switch resource {
	case RIControl:
		return "controls"
	case RIRisk:
		return "risks"
	}
	return ""
}

func ResourceRelationshipToCrudPermissions(from ResourceIdentifier, to ResourceIdentifier) roles.CrudPermissions {
	if from > to {
		return ResourceRelationshipToCrudPermissions(to, from)
	}

	crud := roles.CrudPermissions{
		List:   roles.PNull,
		Create: roles.PNull,
		View:   roles.PNull,
		Update: roles.PNull,
		Delete: roles.PNull,
	}

	switch from {
	case RIRisk:
		switch to {
		case RIControl:
			crud.List = roles.PRelRisksControlsList
			crud.Create = roles.PRelRisksControlsCreate
			crud.Delete = roles.PRelRisksControlsDelete
		}
	}

	return crud
}

func ResourceToCrudPermissions(resource ResourceIdentifier) roles.CrudPermissions {
	crud := roles.CrudPermissions{
		List:   roles.PNull,
		Create: roles.PNull,
		View:   roles.PNull,
		Update: roles.PNull,
		Delete: roles.PNull,
	}

	switch resource {
	case RIControl:
		crud.List = roles.PControlsList
		crud.Create = roles.PControlsCreate
		crud.View = roles.PControlsView
		crud.Update = roles.PControlsUpdate
		crud.Delete = roles.PControlsDelete
	case RIRisk:
		crud.List = roles.PRisksList
		crud.Create = roles.PRisksCreate
		crud.View = roles.PRisksView
		crud.Update = roles.PRisksUpdate
		crud.Delete = roles.PRisksDelete
	case RIInventoryServer:
		crud.List = roles.PServersList
		crud.Create = roles.PServersCreate
		crud.View = roles.PServersView
		crud.Update = roles.PServersUpdate
		crud.Delete = roles.PServersDelete
	case RIInventoryDesktop:
		crud.List = roles.PDesktopsList
		crud.Create = roles.PDesktopsCreate
		crud.View = roles.PDesktopsView
		crud.Update = roles.PDesktopsUpdate
		crud.Delete = roles.PDesktopsDelete
	case RIInventoryLaptop:
		crud.List = roles.PLaptopsList
		crud.Create = roles.PLaptopsCreate
		crud.View = roles.PLaptopsView
		crud.Update = roles.PLaptopsUpdate
		crud.Delete = roles.PLaptopsDelete
	case RIInventoryMobile:
		crud.List = roles.PMobileList
		crud.Create = roles.PMobileCreate
		crud.View = roles.PMobileView
		crud.Update = roles.PMobileUpdate
		crud.Delete = roles.PMobileDelete
	case RIInventoryEmbedded:
		crud.List = roles.PEmbeddedList
		crud.Create = roles.PEmbeddedCreate
		crud.View = roles.PEmbeddedView
		crud.Update = roles.PEmbeddedUpdate
		crud.Delete = roles.PEmbeddedDelete
	}

	return crud
}

func (b *BackendInterface) GetResourceFromGenericId(g GenericResourceIdentifier) (interface{}, error) {
	if g.ControlId != nil {
		return b.GetResource(RIControl, strconv.FormatInt(*g.ControlId, 10))
	} else if g.RiskId != nil {
		return b.GetResource(RIRisk, strconv.FormatInt(*g.RiskId, 10))
	}
	return nil, errors.New("Unsupported GetResourceFromGenericId option.")
}

func (b *BackendInterface) GetResource(id ResourceIdentifier, key string) (interface{}, error) {
	switch id {
	case RIOrganization:
		orgId, err := strconv.ParseInt(key, 10, 64)
		if err != nil {
			return nil, err
		}
		return b.Orgs.GetOrgFromId(orgId)
	case RIEngagement:
		engId, err := strconv.ParseInt(key, 10, 64)
		if err != nil {
			return nil, err
		}
		return b.Engagements.GetEngagementFromId(engId)
	case RIRisk:
		riskId, err := strconv.ParseInt(key, 10, 64)
		if err != nil {
			return nil, err
		}
		return b.Risks.GetRiskFromId(riskId)
	case RIUser:
		userId, err := strconv.ParseInt(key, 10, 64)
		if err != nil {
			return nil, err
		}
		return b.Users.MustGetUserFromId(userId)
	case RIControl:
		ctrlId, err := strconv.ParseInt(key, 10, 64)
		if err != nil {
			return nil, err
		}
		return b.Controls.GetControlFromId(ctrlId)
	case RIGLAccount:
		accId, err := strconv.ParseInt(key, 10, 64)
		if err != nil {
			return nil, err
		}
		return b.GL.GetGLAccountFromId(accId)
	case RIVendor:
		vendorId, err := strconv.ParseInt(key, 10, 64)
		if err != nil {
			return nil, err
		}
		return b.Vendors.GetVendorFromId(vendorId)
	case RIVendorProduct:
		productId, err := strconv.ParseInt(key, 10, 64)
		if err != nil {
			return nil, err
		}
		return b.Vendors.GetVendorProductFromId(productId)
	case RIInventoryServer:
		fallthrough
	case RIInventoryDesktop:
		fallthrough
	case RIInventoryLaptop:
		fallthrough
	case RIInventoryMobile:
		fallthrough
	case RIInventoryEmbedded:
		invId, err := strconv.ParseInt(key, 10, 64)
		if err != nil {
			return nil, err
		}
		return b.Inventory.GetInventory(ResourceToInventoryType(id), invId)
	case RIDatabase:
		dbId, err := strconv.ParseInt(key, 10, 64)
		if err != nil {
			return nil, err
		}
		return b.Databases.GetDatabaseFromId(dbId)
	case RISystem:
		sysId, err := strconv.ParseInt(key, 10, 64)
		if err != nil {
			return nil, err
		}
		return b.Systems.GetSystemFromId(sysId)
	default:
		return nil, errors.New("Unsupported resource identifier.")
	}
}

func GetResourceId(inv interface{}) int64 {
	ref := reflect.ValueOf(inv).Elem()
	return ref.FieldByName("Id").Int()
}

func GetResourceEngagementId(inv interface{}) int64 {
	ref := reflect.ValueOf(inv).Elem()
	return ref.FieldByName("EngagementId").Int()
}
