import { ApiHttpHandler } from '@client/ts/api/handler'
import { RawComment, cleanRawCommentFromJson } from '@client/ts/types/comments'
import { ResourceId } from '@client/ts/types/resource'
import { resourceIdToApiUrl } from '@client/ts/api/url'

function threadIdToUrl(id : ResourceId) : string {
    return `${resourceIdToApiUrl(id)}/comments`
}

export class CommentApiClient {
    handler : ApiHttpHandler
    constructor(handler : ApiHttpHandler) {
        this.handler = handler
    }

    async listComments(id : ResourceId) : Promise<RawComment[]> {
        return this.handler.get<RawComment[]>(threadIdToUrl(id), {}).then((resp : RawComment[]) => {
            resp.forEach(cleanRawCommentFromJson)
            return resp
        })
    }

    async createComment(id : ResourceId, content : string) : Promise<RawComment> {
         return this.handler.post<RawComment>(threadIdToUrl(id), { json: content }).then((resp : RawComment) => {
            cleanRawCommentFromJson(resp)
            return resp
        })
    }

    async updateComment(id : ResourceId, commentId : number, content : string) : Promise<RawComment> {
         return this.handler.put<RawComment>(`${threadIdToUrl(id)}/${commentId}`, { json: content }).then((resp : RawComment) => {
            cleanRawCommentFromJson(resp)
            return resp
        })
    }

    deleteComment(id : ResourceId, commentId : number) : Promise<void> {
         return this.handler.delete<void>(`${threadIdToUrl(id)}/${commentId}`, {})
    }
}
