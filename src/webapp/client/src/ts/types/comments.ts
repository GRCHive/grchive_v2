export interface CommentThreadId {
    orgId : number
    engagementId : number
    riskId? : number
    controlId? : number
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
