import { UserApiClient } from '@client/ts/api/user'
import { ApiHttpHandler } from '@client/ts/api/handler'

class ApiClient {
    handler : ApiHttpHandler
    user : UserApiClient

    constructor() {
        this.handler = new ApiHttpHandler()
        this.user = new UserApiClient(this.handler)
    }
}

export const GrchiveApi = new ApiClient()
