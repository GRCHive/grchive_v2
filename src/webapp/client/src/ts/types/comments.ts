import { ResourceId } from '@client/ts/types/resource'

export type CommentThreadId = ResourceId

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
