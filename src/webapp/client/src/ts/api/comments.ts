import { ApiHttpHandler } from '@client/ts/api/handler'
import { CommentThreadId, RawComment, cleanRawCommentFromJson } from '@client/ts/types/comments'

function threadIdToUrl(id : CommentThreadId) : string {
    const baseUrl : string = `/orgs/${id.orgId}/engagements/${id.engagementId}`

    let apiUrl : string =''
    if (!!id.riskId) {
        apiUrl = `${baseUrl}/risks/${id.riskId}`
    } else if (!!id.controlId) {
        apiUrl = `${baseUrl}/controls/${id.controlId}`
    }
    return `${apiUrl}/comments`
}

export class CommentApiClient {
    handler : ApiHttpHandler
    constructor(handler : ApiHttpHandler) {
        this.handler = handler
    }

    async listComments(id : CommentThreadId) : Promise<RawComment[] | null> {
        return this.handler.get<RawComment[]>(threadIdToUrl(id), {}).then((resp : RawComment[] | null) => {
            if (!resp) {
                return null
            }

            resp.forEach(cleanRawCommentFromJson)
            return resp
        })
    }

    async createComment(id : CommentThreadId, content : string) : Promise<RawComment | null> {
         return this.handler.post<RawComment>(threadIdToUrl(id), { json: content }).then((resp : RawComment | null) => {
            if (!resp) {
                return null
            }

            cleanRawCommentFromJson(resp)
            return resp
        })
       
    }
}
