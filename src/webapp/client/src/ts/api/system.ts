import { ApiHttpHandler } from '@client/ts/api/handler'
import { RawSystem } from '@client/ts/types/systems'

function createBaseApiUrl(orgId : number, engagementId : number) : string {
    return `/orgs/${orgId}/engagements/${engagementId}/systems`
}

function createSingleBaseApiUrl(orgId : number, engagementId : number, systemId : number) : string {
    return `${createBaseApiUrl(orgId, engagementId)}/${systemId}`
}

export class SystemApiClient {
    handler : ApiHttpHandler
    constructor(handler : ApiHttpHandler) {
        this.handler = handler
    }

    listSystems(orgId : number, engagementId : number): Promise<RawSystem[]> {
        return this.handler.get<RawSystem[]>(createBaseApiUrl(orgId, engagementId), {})
    }

    createSystem(orgId: number, system : RawSystem) : Promise<RawSystem> {
        return this.handler.post<RawSystem>(createBaseApiUrl(orgId, system.EngagementId), {json: system})
    }

    getSystem(orgId : number, engagementId : number, systemId : number) : Promise<RawSystem> {
        return this.handler.get<RawSystem>(createSingleBaseApiUrl(orgId, engagementId, systemId), {})
    }

    deleteSystem(orgId : number, engagementId : number, systemId : number) : Promise<void> {
        return this.handler.delete<void>(createSingleBaseApiUrl(orgId, engagementId, systemId), {})
    }

    updateSystem(orgId : number, system : RawSystem) : Promise<RawSystem> {
        return this.handler.put<RawSystem>(createSingleBaseApiUrl(orgId, system.EngagementId, system.Id), {json : system})
    }
}
