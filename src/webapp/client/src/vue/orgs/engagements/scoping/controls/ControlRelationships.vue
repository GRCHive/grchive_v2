<template>
    <div>
        <v-tabs>
            <restrict-role-permission-tab
                :permissions="riskPermissions"
                :to="riskTo"
                :org-id="currentOrg.Id"
                :engagement-id="currentEngagement.Id"
            >
                Risks
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
export default class ControlRelationships extends Vue {
    get currentOrg() : RawOrganization | null {
        return this.$store.state.org.rawOrg
    }

    get currentEngagement() : RawEngagement | null {
        return this.$store.state.engagements.rawEngagement
    }

    get riskPermissions() : Permission[] {
        return [Permission.PRelRisksControlsList, Permission.PRisksList]
    }

    get riskTo() : any {
        return {
            name: 'controlRiskRelationships',
            params: this.$route.params,
        }
    }

    get fullPermissions(): any {
        if (!this.currentOrg || !this.currentEngagement) {
            return null
        }

        return {
            risks: this.$store.getters['permission/currentUserHasPermissions'](this.currentOrg!.Id, this.currentEngagement!.Id, this.riskPermissions),
        }
    }

    @Watch('fullPermissions')
    @Watch('$route')
    redirectToSubpage() {
        if (this.$route.name !== 'controlRelationships') {
            return
        }

        if (!this.currentOrg || !this.currentEngagement) {
            return
        }

        if (this.fullPermissions.risks === null) {
            return
        }

        if (this.fullPermissions.risks === true) {
            this.$router.replace(this.riskTo)
        } else {
            ErrorHandler.failurePageOnManualError(
                createManualUnauthorizedError('You do not have sufficient privileges to view any control relationships.',
                    new PermissionOr([
                        new PermissionAnd(this.riskPermissions.map((ele : Permission) => new SinglePermission(ele))),
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
