<script lang="ts">

import Vue from 'vue'
import Component from 'vue-class-component'
import { Watch, Prop } from 'vue-property-decorator'
import { Permission } from '@client/ts/types/roles'

@Component
export default class RestrictRolePermissionBase extends Vue {
    @Prop({required: true})
    permissions! : Permission[]

    @Prop({required: true})
    orgId!: number

    @Prop({ default: -1 })
    engagementId!: number

    get hasPermissions() : boolean | null {
        return this.$store.getters['permission/currentUserHasPermissions'](this.orgId, this.engagementId, this.permissions)
    }

    get tooltipStr() : string {
        return `You do not have the required permissions to access this feature. Please request the ${this.permissions.map(p => `'${p}'`).join(', ')} permission(s) from your administrator.`
    }

    get isLoading() : boolean {
        return this.hasPermissions === null
    }

    @Watch('permissions')
    @Watch('orgId')
    @Watch('engagementId')
    refreshPermission() {
        this.$store.dispatch('permission/initializeHasPermissions', {
            orgId: this.orgId,
            engagementId: this.engagementId,
            permissions: this.permissions,
        })
    }

    mounted() {
        this.refreshPermission()
    }
}

</script>
