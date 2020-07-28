import { v4 as uuidv4 } from 'uuid'

export interface RawSystem {
    Id                  : number
    EngagementId        : number
    UniqueId            : string
    Name                : string
    Description         : string
    Purpose             : string
}

export function createEmptySystem() : RawSystem {
    return {
        Id: -1,
        EngagementId: -1,
        UniqueId: uuidv4(),
        Name: '',
        Description: '',
        Purpose: '',
    }
}
