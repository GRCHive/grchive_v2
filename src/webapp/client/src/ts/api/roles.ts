import { ApiHttpHandler } from '@client/ts/api/handler'
import { Role } from '@client/ts/types/roles'

export class RoleApiClient {
    handler : ApiHttpHandler
    constructor(handler : ApiHttpHandler) {
        this.handler = handler
    }

    async listRoles(orgId : number) : Promise<Role[]> {
        return this.handler.get<Role[]>(`/orgs/${orgId}/roles`, {})
    }
}
