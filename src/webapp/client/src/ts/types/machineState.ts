export enum MachineStateType {
    None,
    Hypervisor,
    OperatingSystem
}

export interface MachineState {
    Id              : number
    EngagementId    : number
    UniqueId        : string
    Hypervisor      : string | null
    OperatingSystem : string | null
}

export function createEmptyMachineState() : MachineState {
    return {
        Id: -1,
        EngagementId: -1,
        UniqueId: '',
        Hypervisor: null,
        OperatingSystem: null
    }
}
