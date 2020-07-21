import { RawUser } from '@client/ts/types/users'
import { RawOrganization } from '@client/ts/types/orgs'
import { ApiHttpHandler } from '@client/ts/api/handler'
import { Permission } from '@client/ts/types/roles'
import * as qs from 'query-string'

export class UserApiClient {
    handler : ApiHttpHandler
    constructor(handler : ApiHttpHandler) {
        this.handler = handler
    }

    getCurrentUser() : Promise<RawUser | null> {
        return this.handler.get('/users/current', {})
    }

    getCurrentUserOrgs() : Promise<RawOrganization[] | null> {
        return this.handler.get('/users/current/orgs', {})
    }

    updateCurrentUser(user : RawUser) : Promise<void | null> {
        return this.handler.put('/users/current', {json : user})
    }

    resendEmailVerification() : Promise<void | null> {
        return this.handler.post('/users/current/verify', {})
    }

    checkCurrentUserPermissions(orgId : number, engagementId : number, permissions : Permission[]) : Promise<boolean | null> {
        if (engagementId == -1) {
            return this.handler.get(`/users/current/orgs/${orgId}/permissions`, {
                searchParams: qs.stringify({
                    permissions
                })
            })
        } else {
            return this.handler.get(`/users/current/orgs/${orgId}/engagement/${engagementId}/permissions`, {
                searchParams: qs.stringify({
                    permissions
                })
            })
        }
    }
}
