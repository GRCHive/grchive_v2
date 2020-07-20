import { ApiHttpHandler } from '@client/ts/api/handler'
import {
    RawEngagement,
    cleanRawEngagementFromJSON,
    engagementToJson
} from '@client/ts/types/engagements'

export class EngagementApiClient {
    handler : ApiHttpHandler
    constructor(handler : ApiHttpHandler) {
        this.handler = handler
    }

    async listOrgEngagements(orgId : number) : Promise<RawEngagement[] | null> {
        return this.handler.get<RawEngagement[]>(`/orgs/${orgId}/engagements`, {}).then((resp : RawEngagement[] | null) => {
            if (!resp) {
                return resp
            }

            resp.forEach(cleanRawEngagementFromJSON)
            return resp
        })
    }

    async createEngagement(eng : RawEngagement) : Promise<RawEngagement | null> {
        return this.handler.post<RawEngagement>(`/orgs/${eng.OrgId}/engagements`, {json: engagementToJson(eng)}).then((resp : RawEngagement | null) => {
            if (!resp) {
                return resp
            }

            cleanRawEngagementFromJSON(resp)
            return resp
        })
    }


    async updateEngagement(eng : RawEngagement) : Promise<RawEngagement | null> {
        return this.handler.put<RawEngagement>(`/orgs/${eng.OrgId}/engagements/${eng.Id}`, {json: engagementToJson(eng)}).then((resp : RawEngagement | null) => {
            if (!resp) {
                return resp
            }

            cleanRawEngagementFromJSON(resp)
            return resp
        })
    }


    async getEngagement(orgId : number, engId : number) : Promise<RawEngagement | null> {
        return this.handler.get<RawEngagement>(`/orgs/${orgId}/engagements/${engId}`, {}).then((resp : RawEngagement | null) => {
            if (!resp) {
                return resp
            }

            cleanRawEngagementFromJSON(resp)
            return resp
        })
    }
}
