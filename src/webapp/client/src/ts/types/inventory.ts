import { v4 as uuidv4 } from 'uuid'
export interface RawInventory {
    Id:                 number
    EngagementId:       number
    UniqueId:           string
    Name:               string
    Description:        string
    Brand:              string
    Model:              string
}

function createEmptyInventory() : RawInventory {
    return {
        Id: -1,
        EngagementId: -1,
        UniqueId: uuidv4(),
        Name: '',
        Description: '',
        Brand: '',
        Model: '',
    }
}

export enum InventoryType {
    ITServer,
    ITDesktop,
    ITLaptop,
    ITMobileDevice,
    ITEmbeddedDevice
}

export interface RawBaseInventory {
    Id:                 number
    Inventory:          RawInventory
}

export interface RawServer extends RawBaseInventory {
    PhysicalLocation:   string
    OperatingSystem:    string | null
    Hypervisor:         string | null
    StaticExternalIp:   string | null
}

export function createEmptyServer() : RawServer {
    return {
        Id: -1,
        Inventory: createEmptyInventory(),
        PhysicalLocation: '',
        OperatingSystem: null,
        Hypervisor: null,
        StaticExternalIp: null,
    }
}

export interface RawDesktop extends RawBaseInventory {
    PhysicalLocation:   string
    OperatingSystem:    string
}

export interface RawLaptop extends RawBaseInventory {
    OperatingSystem:    string
}

export interface RawMobileDevice extends RawBaseInventory {
    OperatingSystem:    string
}

export interface RawEmbeddedDevice extends RawBaseInventory {
    OperatingSystem:    string
}
