package backend

import (
	"errors"
	"fmt"
	"gitlab.com/grchive/grchive-v2/shared/backend/inventory"
	"gitlab.com/grchive/grchive-v2/shared/backend/roles"
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
)

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

func ResourceToCrudPermissions(resource ResourceIdentifier) roles.CrudPermissions {
	crud := roles.CrudPermissions{
		List:   roles.PNull,
		Create: roles.PNull,
		View:   roles.PNull,
		Update: roles.PNull,
		Delete: roles.PNull,
	}

	switch resource {
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
		invId, err := strconv.ParseInt(key, 10, 64)
		if err != nil {
			return nil, err
		}
		return b.Inventory.GetInventory(ResourceToInventoryType(id), invId)
	default:
		return nil, errors.New("Unsupported resource identifier.")
	}
}
