export interface CommentThreadId {
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

export interface RawComment {
    Id           : number
    ThreadId     : number
    UserId       : number
    PostTime     : Date
    Content      : string
}

export function cleanRawCommentFromJson(r : RawComment) {
    r.PostTime = new Date(r.PostTime)
}
