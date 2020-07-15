import { Store } from 'vuex'
import { RootState } from '@client/ts/stores/store'
import { UserApiClient } from '@client/ts/api/user'
import { OrgApiClient } from '@client/ts/api/orgs'
import { ApiHttpHandler } from '@client/ts/api/handler'

export class ApiClient {
    handler : ApiHttpHandler
    user : UserApiClient
    orgs : OrgApiClient

    store : Store<RootState>

    constructor(store : Store<RootState>) {
        this.store = store
        this.handler = new ApiHttpHandler(this.store)
        this.user = new UserApiClient(this.handler)
        this.orgs = new OrgApiClient(this.handler)
    }
}
