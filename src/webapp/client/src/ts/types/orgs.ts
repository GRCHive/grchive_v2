export interface RawOrganization {
    Id          : number
    Name        : string
    Description : string
    ParentOrgId : number | null
    OwnerUserId : number
}

export function createEmptyOrg() : RawOrganization {
    return {
        Id: -1,
        Name: '',
        Description: '',
        ParentOrgId: null,
        OwnerUserId: -1
    }
}
