export interface RawVendor {
    Id              : number
    EngagementId    : number
    Name            : string
    Description     : string
    Url             : string
}

export interface RawVendorProduct {
    Id              : number
    VendorId        : number
    Name            : string
    Description     : string
    Url             : string
}

export function createEmptyVendor() : RawVendor {
    return {
        Id: -1,
        EngagementId: -1,
        Name: '',
        Description: '',
        Url: '',
    }
}

export function createEmptyVendorProduct() : RawVendorProduct {
    return {
        Id: -1,
        VendorId: -1,
        Name: '',
        Description: '',
        Url: '',
    }
}
