export interface RawEngagement {
    Id            : number
    Name          : string
    Description   : string
    OrgId         : number
    CreatedTime   : Date
    StartTime     : Date | null
    EndTime       : Date | null
    IsClosed      : boolean
}

export function createEmptyEngagement() : RawEngagement {
    return {
        Id: -1,
        Name: '',
        Description: '',
        OrgId: -1,
        CreatedTime: new Date(),
        StartTime: null,
        EndTime: null,
        IsClosed: false,
    }
}

export function cleanRawEngagementFromJSON(r : RawEngagement) {
    r.CreatedTime = new Date(r.CreatedTime)
    if (!!r.StartTime) {
        r.StartTime = new Date(r.StartTime)
    }

    if (!!r.EndTime) {
        r.EndTime = new Date(r.EndTime)
    }
}
