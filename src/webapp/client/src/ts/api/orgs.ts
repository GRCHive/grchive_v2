import { RawOrganization } from '@client/ts/types/orgs'
import { ApiHttpHandler } from '@client/ts/api/handler'

export class OrgApiClient {
    handler : ApiHttpHandler
    constructor(handler : ApiHttpHandler) {
        this.handler = handler
    }

    createOrg(org : RawOrganization) : Promise<RawOrganization | null> {
        return this.handler.post('/orgs', {json : org})
    }

    updateOrg(org : RawOrganization) : Promise<RawOrganization | null> {
        return this.handler.put(`/orgs/${org.Id}`, {json : org})
    }

    getOrg(orgId : number) : Promise<RawOrganization | null> {
        return this.handler.get(`/orgs/${orgId}`, {})
    }

    getSuborgs(orgId: number) : Promise<RawOrganization[] | null> {
        return this.handler.get(`/orgs/${orgId}/suborgs`, {})
    }
}
