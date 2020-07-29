import { UserApiClient } from '@client/ts/api/user'
import { OrgApiClient } from '@client/ts/api/orgs'
import { EngagementApiClient } from '@client/ts/api/engagements'
import { RoleApiClient } from '@client/ts/api/roles'
import { RiskApiClient } from '@client/ts/api/risks'
import { ControlApiClient } from '@client/ts/api/controls'
import { CommentApiClient } from '@client/ts/api/comments'
import { GeneralLedgerApiClient } from '@client/ts/api/gl'
import { VendorApiClient } from '@client/ts/api/vendor'
import { InventoryApiClient } from '@client/ts/api/inventory'
import { DatabaseApiClient } from '@client/ts/api/database'
import { SystemApiClient } from '@client/ts/api/system'
import { RelationshipApiClient } from '@client/ts/api/relationships'
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
    inventory : InventoryApiClient
    databases : DatabaseApiClient
    systems : SystemApiClient
    relationships : RelationshipApiClient

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
        this.inventory = new InventoryApiClient(this.handler)
        this.databases = new DatabaseApiClient(this.handler)
        this.systems = new SystemApiClient(this.handler)
        this.relationships = new RelationshipApiClient(this.handler)
    }
}
