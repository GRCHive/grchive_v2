import { ApiHttpHandler } from '@client/ts/api/handler'
import { CommentThreadId, RawComment, cleanRawCommentFromJson } from '@client/ts/types/comments'
import {
    InventoryType
} from '@client/ts/types/inventory'
import { createBaseApiUrl } from '@client/ts/api/inventory'

function threadIdToUrl(id : CommentThreadId) : string {
    const baseUrl : string = `/orgs/${id.orgId}/engagements/${id.engagementId}`

    let apiUrl : string =''
    if (!!id.riskId) {
        apiUrl = `${baseUrl}/risks/${id.riskId}`
    } else if (!!id.controlId) {
        apiUrl = `${baseUrl}/controls/${id.controlId}`
    } else if (!!id.glAccountId) {
        if (id.glAccountId === -1) {
            apiUrl = `${baseUrl}/gl`
        } else {
            apiUrl = `${baseUrl}/gl/accs/${id.glAccountId}`
        }
    } else if (!!id.vendorId) {
        apiUrl = `${baseUrl}/vendors/${id.vendorId}`
    } else if (!!id.serverId) {
        apiUrl = `${createBaseApiUrl(InventoryType.ITServer, id.orgId, id.engagementId)}/${id.serverId}`
    } else if (!!id.desktopId) {
        apiUrl = `${createBaseApiUrl(InventoryType.ITDesktop, id.orgId, id.engagementId)}/${id.desktopId}`
    } else if (!!id.laptopId) {
        apiUrl = `${createBaseApiUrl(InventoryType.ITLaptop, id.orgId, id.engagementId)}/${id.laptopId}`
    } else if (!!id.mobileId) {
        apiUrl = `${createBaseApiUrl(InventoryType.ITMobile, id.orgId, id.engagementId)}/${id.mobileId}`
    } else if (!!id.embeddedId) {
        apiUrl = `${createBaseApiUrl(InventoryType.ITEmbedded, id.orgId, id.engagementId)}/${id.embeddedId}`
    } else if (!!id.databaseId) {
        apiUrl = `${baseUrl}/databases/${id.databaseId}`
    } else if (!!id.systemId) {
        apiUrl = `${baseUrl}/systems/${id.systemId}`
    }
    return `${apiUrl}/comments`
}

export class CommentApiClient {
    handler : ApiHttpHandler
    constructor(handler : ApiHttpHandler) {
        this.handler = handler
    }

    async listComments(id : CommentThreadId) : Promise<RawComment[]> {
        return this.handler.get<RawComment[]>(threadIdToUrl(id), {}).then((resp : RawComment[]) => {
            resp.forEach(cleanRawCommentFromJson)
            return resp
        })
    }

    async createComment(id : CommentThreadId, content : string) : Promise<RawComment> {
         return this.handler.post<RawComment>(threadIdToUrl(id), { json: content }).then((resp : RawComment) => {
            cleanRawCommentFromJson(resp)
            return resp
        })
    }

    async updateComment(id : CommentThreadId, commentId : number, content : string) : Promise<RawComment> {
         return this.handler.put<RawComment>(`${threadIdToUrl(id)}/${commentId}`, { json: content }).then((resp : RawComment) => {
            cleanRawCommentFromJson(resp)
            return resp
        })
    }

    deleteComment(id : CommentThreadId, commentId : number) : Promise<void> {
         return this.handler.delete<void>(`${threadIdToUrl(id)}/${commentId}`, {})
    }
}
