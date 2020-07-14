import { RawUser } from '@client/ts/types/users'
import { ApiHttpHandler } from '@client/ts/api/handler'

export class UserApiClient {
    handler : ApiHttpHandler
    constructor(handler : ApiHttpHandler) {
        this.handler = handler
    }

    getCurrentUser() : Promise<RawUser> {
        return this.handler.get('/users/current', {})
    }

    resendEmailVerification() : Promise<void> {
        return this.handler.post('/users/current/verify', {})
    }
}
