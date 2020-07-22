export enum ControlFrequencyType {
	CFTAdHoc     = 1,
	CFTInterval  = 2,
	CFTOther     = 3,
}

export enum ControlType {
	CTAccess           = 1,
	CTAuthorization    = 2,
	CTConfiguration    = 3,
	CTGitc             = 4,
	CTInterface        = 5,
	CTManagementReview = 6,
	CTReconciliation   = 7,
	CTReport           = 8,
}

export interface RawControl {
    Id                     : number
    EngagementId           : number
	Name                   : string
	HumanId                : string
	Likelihood             : number
	LikelihoodJustification: string
	ControlType            : ControlType
	OwnerId                : number | null
	IsManual               : boolean
	FreqType               : ControlFrequencyType
	FreqOther              : string
	FreqInterval           : string
	Description            : string
}

export function createEmptyControl() : RawControl {
    return {
        Id: -1,
        EngagementId: -1,
        Name: '',
        HumanId: '',
        Likelihood: 0,
        LikelihoodJustification: '',
        ControlType: ControlType.CTAccess,
        OwnerId: null,
        IsManual: true,
        FreqType: ControlFrequencyType.CFTAdHoc,
        FreqOther: '',
        FreqInterval: '',
        Description: '',
    }
}

export function controlTypeToString(ct : ControlType) : string {
    switch (ct) {
        case ControlType.CTAccess:
            return 'Access'
        case ControlType.CTAuthorization:
            return 'Authorization'
        case ControlType.CTConfiguration:
            return 'Configuration'
        case ControlType.CTGitc:
            return 'GITC'
        case ControlType.CTInterface:
            return 'Interface'
        case ControlType.CTManagementReview:
            return 'Management Review'
        case ControlType.CTReconciliation:
            return 'Reconciliation'
        case ControlType.CTReport:
            return 'Report/IPE'
    }
}
