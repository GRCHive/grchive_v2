import { Store } from 'vuex'
import { RootState } from '@client/ts/stores/store'
import { UserApiClient } from '@client/ts/api/user'
import { OrgApiClient } from '@client/ts/api/orgs'
import { EngagementApiClient } from '@client/ts/api/engagements'
import { RoleApiClient } from '@client/ts/api/roles'
import { RiskApiClient } from '@client/ts/api/risks'
import { ControlApiClient } from '@client/ts/api/controls'
import { ApiHttpHandler } from '@client/ts/api/handler'

export class ApiClient {
    handler : ApiHttpHandler
    user : UserApiClient
    orgs : OrgApiClient
    engagements: EngagementApiClient
    roles : RoleApiClient
    risks : RiskApiClient
    controls : ControlApiClient

    store : Store<RootState>

    constructor(store : Store<RootState>) {
        this.store = store
        this.handler = new ApiHttpHandler(this.store)
        this.user = new UserApiClient(this.handler)
        this.orgs = new OrgApiClient(this.handler)
        this.engagements = new EngagementApiClient(this.handler)
        this.roles = new RoleApiClient(this.handler)
        this.risks = new RiskApiClient(this.handler)
        this.controls = new ControlApiClient(this.handler)
    }
}
