import { v4 as uuidv4 } from 'uuid'
import { MachineState, createEmptyMachineState } from '@client/ts/types/machineState'

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
    StaticExternalIp:   string | null
    State:              MachineState
}

export function createEmptyServer() : RawServer {
    return {
        Id: -1,
        Inventory: createEmptyInventory(),
        PhysicalLocation: '',
        StaticExternalIp: null,
        State: createEmptyMachineState(),
    }
}

export interface RawDesktop extends RawBaseInventory {
    PhysicalLocation:   string
    State:              MachineState
}

export function createEmptyDesktop() : RawDesktop {
    return {
        Id: -1,
        Inventory: createEmptyInventory(),
        PhysicalLocation: '',
        State: createEmptyMachineState(),
    }
}

export interface RawLaptop extends RawBaseInventory {
    State:              MachineState
}


export function createEmptyLaptop() : RawLaptop {
    return {
        Id: -1,
        Inventory: createEmptyInventory(),
        State: createEmptyMachineState(),
    }
}

export interface RawMobile extends RawBaseInventory {
    State:              MachineState
}

export function createEmptyMobile() : RawMobile {
    return {
        Id: -1,
        Inventory: createEmptyInventory(),
        State: createEmptyMachineState(),
    }
}

export interface RawEmbedded extends RawBaseInventory {
    State:              MachineState
}

export function createEmptyEmbedded() : RawEmbedded {
    return {
        Id: -1,
        Inventory: createEmptyInventory(),
        State: createEmptyMachineState(),
    }
}
