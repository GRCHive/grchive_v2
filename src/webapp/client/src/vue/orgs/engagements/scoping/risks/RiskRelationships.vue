<template>
    <div>
        <v-tabs>
            <restrict-role-permission-tab
                :permissions="controlPermissions"
                :to="controlsTo"
                :org-id="currentOrg.Id"
                :engagement-id="currentEngagement.Id"
            >
                Controls
            </restrict-role-permission-tab>
        </v-tabs>
        <router-view></router-view>
    </div>
</template>

<script lang="ts">

import Vue from 'vue'
import Component from 'vue-class-component'
import { Watch } from 'vue-property-decorator'
import { Permission, PermissionAnd, PermissionOr, SinglePermission } from '@client/ts/types/roles'
import { RawOrganization } from '@client/ts/types/orgs'
import { RawEngagement } from '@client/ts/types/engagements'
import { ErrorHandler } from '@client/ts/main'
import { createManualUnauthorizedError } from '@client/ts/error'
import RestrictRolePermissionTab from '@client/vue/loading/RestrictRolePermissionTab.vue'

@Component({
    components: {
        RestrictRolePermissionTab
    }
})
export default class RiskRelationships extends Vue {
    get currentOrg() : RawOrganization | null {
        return this.$store.state.org.rawOrg
    }

    get currentEngagement() : RawEngagement | null {
        return this.$store.state.engagements.rawEngagement
    }

    get controlPermissions() : Permission[] {
        return [Permission.PRelRisksControlsList, Permission.PControlsList]
    }

    get controlsTo() : any {
        return {
            name: 'riskControlRelationships',
            params: this.$route.params,
        }
    }

    get fullPermissions(): any {
        if (!this.currentOrg || !this.currentEngagement) {
            return null
        }

        return {
            controls: this.$store.getters['permission/currentUserHasPermissions'](this.currentOrg!.Id, this.currentEngagement!.Id, this.controlPermissions),
        }
    }

    @Watch('fullPermissions')
    @Watch('$route')
    redirectToSubpage() {
        if (this.$route.name !== 'riskRelationships') {
            return
        }

        if (!this.currentOrg || !this.currentEngagement) {
            return
        }

        if (this.fullPermissions.controls === null) {
            return
        }

        if (this.fullPermissions.controls === true) {
            this.$router.replace({
                name: 'riskControlRelationships',
                params: this.$route.params
            })
        } else {
            ErrorHandler.failurePageOnManualError(
                createManualUnauthorizedError('You do not have sufficient privileges to view any risk relationships.',
                    new PermissionOr([
                        new PermissionAnd(this.controlPermissions.map((ele : Permission) => new SinglePermission(ele))),
                    ]),
                )
            )
        }
    }

    mounted() {
        this.redirectToSubpage()
    }
}

</script>
