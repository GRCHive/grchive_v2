<script lang="ts">

import Vue from 'vue'
import Component from 'vue-class-component'
import { Watch, Prop } from 'vue-property-decorator'
import { GrchiveApi } from '@client/ts/main'
import { Permission } from '@client/ts/types/roles'

@Component
export default class RestrictRolePermissionBase extends Vue {
    @Prop({required: true})
    permissions! : Permission[]

    @Prop({required: true})
    orgId!: number

    hasPermissions: boolean | null = null

    get tooltipStr() : string {
        return `You do not have the required permissions to access this feature. Please request the ${this.permissions.join(', ')} permission(s) from your administrator.`
    }

    get isLoading() : boolean {
        return this.hasPermissions === null
    }

    @Watch('permissions')
    refreshPermission() {
        GrchiveApi.user.checkCurrentUserPermissions(this.orgId, this.permissions).then((resp : boolean | null) => {
            this.hasPermissions = resp
        })
    }

    mounted() {
        this.refreshPermission()
    }
}

</script>
