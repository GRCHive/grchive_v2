import { ApiHttpHandler } from '@client/ts/api/handler'
import { CommentThreadId, RawComment, cleanRawCommentFromJson } from '@client/ts/types/comments'

function threadIdToUrl(id : CommentThreadId) : string {
    const baseUrl : string = `/orgs/${id.orgId}/engagements/${id.engagementId}`

    let apiUrl : string =''
    if (!!id.riskId) {
        apiUrl = `${baseUrl}/risks/${id.riskId}`
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
}
