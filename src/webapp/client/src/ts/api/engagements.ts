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

    async listOrgEngagements(orgId : number) : Promise<RawEngagement[]> {
        return this.handler.get<RawEngagement[]>(`/orgs/${orgId}/engagements`, {}).then((resp : RawEngagement[]) => {
            resp.forEach(cleanRawEngagementFromJSON)
            return resp
        })
    }

    async createEngagement(eng : RawEngagement, baseEngagement : RawEngagement | null) : Promise<RawEngagement> {
        return this.handler.post<RawEngagement>(`/orgs/${eng.OrgId}/engagements`, {json: {
                new: engagementToJson(eng),
                base: !!baseEngagement ? engagementToJson(baseEngagement) : null,
            }}).then((resp : RawEngagement) => {
            cleanRawEngagementFromJSON(resp)
            return resp
        })
    }


    async updateEngagement(eng : RawEngagement) : Promise<RawEngagement> {
        return this.handler.put<RawEngagement>(`/orgs/${eng.OrgId}/engagements/${eng.Id}`, {json: engagementToJson(eng)}).then((resp : RawEngagement) => {
            cleanRawEngagementFromJSON(resp)
            return resp
        })
    }


    async getEngagement(orgId : number, engId : number) : Promise<RawEngagement> {
        return this.handler.get<RawEngagement>(`/orgs/${orgId}/engagements/${engId}`, {}).then((resp : RawEngagement) => {
            cleanRawEngagementFromJSON(resp)
            return resp
        })
    }
}
