import formatRFC3339 from 'date-fns/formatRFC3339'
import { Role } from '@client/ts/types/roles'

export interface RawEngagement {
    Id            : number
    Name          : string
    Description   : string
    OrgId         : number
    CreatedTime   : Date
    StartTime     : Date | null
    EndTime       : Date | null
    IsClosed      : boolean
    Roles         : Role[] | null
}

export interface RawEngagementScopingStats {
    NumRisks :                          number
    NumControls:                        number
    NumGLAccounts:                      number
    NumFinanciallyRelevantGLAccounts:   number
    NumVendors:                         number
    NumVendorProducts:                  number
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
        Roles: null,
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

export function engagementToJson(r : RawEngagement) : any {
    let ret : any = {
        ...r
    }

    ret.CreatedTime = formatRFC3339(r.CreatedTime)

    if (!!r.StartTime) {
        ret.StartTime = formatRFC3339(r.StartTime)
    }

    if (!!r.EndTime) {
        ret.EndTime = formatRFC3339(r.EndTime)
    }

    return ret
}
