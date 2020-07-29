export interface ResourceId {
    orgId : number
    engagementId : number
    riskId? : number
    controlId? : number
    glAccountId? : number
    vendorId?: number
    serverId? : number
    desktopId? : number
    laptopId? : number
    mobileId? : number
    embeddedId? : number
    databaseId? : number
    systemId? : number
}

export enum ResourceType {
    RTControl,
    RTRisk
}
