export interface RawRisk {
    Id                    : number
    Name                  : string
    HumanId               : string
    EngagementId          : number
    Severity              : number
    SeverityJustification : string
    Description           : string
}

export function createEmptyRisk() : RawRisk {
    return {
        Id: -1,
        Name: '',
        HumanId: '',
        EngagementId: -1,
        Severity: 0,
        SeverityJustification: '',
        Description: '',
    }
}
