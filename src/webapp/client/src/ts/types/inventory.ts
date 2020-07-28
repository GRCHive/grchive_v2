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
    ITMobile,
    ITEmbedded
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

export function createEmptyDesktop() : RawDesktop {
    return {
        Id: -1,
        Inventory: createEmptyInventory(),
        PhysicalLocation: '',
        OperatingSystem: '',
    }
}

export interface RawLaptop extends RawBaseInventory {
    OperatingSystem:    string
}


export function createEmptyLaptop() : RawLaptop {
    return {
        Id: -1,
        Inventory: createEmptyInventory(),
        OperatingSystem: '',
    }
}

export interface RawMobile extends RawBaseInventory {
    OperatingSystem:    string
}

export function createEmptyMobile() : RawMobile {
    return {
        Id: -1,
        Inventory: createEmptyInventory(),
        OperatingSystem: '',
    }
}

export interface RawEmbedded extends RawBaseInventory {
    OperatingSystem:    string
}

export function createEmptyEmbedded() : RawEmbedded {
    return {
        Id: -1,
        Inventory: createEmptyInventory(),
        OperatingSystem: '',
    }
}
