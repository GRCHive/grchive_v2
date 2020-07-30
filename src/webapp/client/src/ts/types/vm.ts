import { MachineState, createEmptyMachineState } from '@client/ts/types/machineState'
export interface VirtualMachine {
    Id                  : number
    EngagementId        : number
    UniqueId            : string
    Name                : string
    Description         : string
    VCpus               : number
    MemoryBytes         : number
    State               : MachineState
}

export function createEmptyVirtualMachine() : VirtualMachine {
    return {
        Id: -1,
        EngagementId: -1,
        UniqueId: '',
        Name: '',
        Description: '',
        VCpus: 0,
        MemoryBytes: 0,
        State: createEmptyMachineState(),
    }
}
