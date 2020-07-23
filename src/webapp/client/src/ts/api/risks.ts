import { ApiHttpHandler } from '@client/ts/api/handler'
import {
    RawRisk,
} from '@client/ts/types/risks'

export class RiskApiClient {
    handler : ApiHttpHandler
    constructor(handler : ApiHttpHandler) {
        this.handler = handler
    }

    createRisk(orgId : number, risk : RawRisk) : Promise<RawRisk> {
        return this.handler.post<RawRisk>(`/orgs/${orgId}/engagements/${risk.EngagementId}/risks`, {json: risk})
    }

    updateRisk(orgId : number, risk : RawRisk) : Promise<RawRisk> {
        return this.handler.put<RawRisk>(`/orgs/${orgId}/engagements/${risk.EngagementId}/risks/${risk.Id}`, {json: risk})
    }

    deleteRisk(orgId : number, engagementId : number, riskId : number) : Promise<void> {
        return this.handler.delete<void>(`/orgs/${orgId}/engagements/${engagementId}/risks/${riskId}`, {})
    }

    getRisk(orgId : number, engagementId: number, riskId : number) : Promise<RawRisk> {
        return this.handler.get<RawRisk>(`/orgs/${orgId}/engagements/${engagementId}/risks/${riskId}`, {})
    }

    listRisks(orgId : number, engagementId : number): Promise<RawRisk[]> {
        return this.handler.get<RawRisk[]>(`/orgs/${orgId}/engagements/${engagementId}/risks`, {})
    }
}
