import { RawUser } from '@client/ts/types/users'
import { RawOrganization } from '@client/ts/types/orgs'
import { ApiHttpHandler } from '@client/ts/api/handler'

export class UserApiClient {
    handler : ApiHttpHandler
    constructor(handler : ApiHttpHandler) {
        this.handler = handler
    }

    getCurrentUser() : Promise<RawUser | null> {
        return this.handler.get('/users/current', {})
    }

    getUserOrgs() : Promise<RawOrganization[] | null> {
        return this.handler.get('/users/current/orgs', {})
    }

    updateUser(user : RawUser) : Promise<void | null> {
        return this.handler.put('/users/current', {json : user})
    }

    resendEmailVerification() : Promise<void | null> {
        return this.handler.post('/users/current/verify', {})
    }
}
