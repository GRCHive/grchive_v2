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

    createSuborg(parentOrgId: number, org : RawOrganization) : Promise<RawOrganization | null> {
        return this.handler.post(`/orgs/${parentOrgId}`, {json : org})
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

    // Returns a list that starts with the org that was passed in
    async getParentOrgs(orgId : number) : Promise<RawOrganization[] | null> {
        return this.handler.get(`/orgs/${orgId}/parents`, {})
    }
}
