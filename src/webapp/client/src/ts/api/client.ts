import { UserApiClient } from '@client/ts/api/user'
import { OrgApiClient } from '@client/ts/api/orgs'
import { EngagementApiClient } from '@client/ts/api/engagements'
import { RoleApiClient } from '@client/ts/api/roles'
import { RiskApiClient } from '@client/ts/api/risks'
import { ControlApiClient } from '@client/ts/api/controls'
import { CommentApiClient } from '@client/ts/api/comments'
import { GeneralLedgerApiClient } from '@client/ts/api/gl'
import { VendorApiClient } from '@client/ts/api/vendor'
import { ApiHttpHandler } from '@client/ts/api/handler'

export class ApiClient {
    handler : ApiHttpHandler
    user : UserApiClient
    orgs : OrgApiClient
    engagements: EngagementApiClient
    roles : RoleApiClient
    risks : RiskApiClient
    controls : ControlApiClient
    comments : CommentApiClient
    gl: GeneralLedgerApiClient
    vendors: VendorApiClient

    constructor() {
        this.handler = new ApiHttpHandler()
        this.user = new UserApiClient(this.handler)
        this.orgs = new OrgApiClient(this.handler)
        this.engagements = new EngagementApiClient(this.handler)
        this.roles = new RoleApiClient(this.handler)
        this.risks = new RiskApiClient(this.handler)
        this.controls = new ControlApiClient(this.handler)
        this.comments = new CommentApiClient(this.handler)
        this.gl = new GeneralLedgerApiClient(this.handler)
        this.vendors = new VendorApiClient(this.handler)
    }
}
