<template>
    <scoping-template
        :page-name="riskId"
    >
        <template v-slot:content>
            <v-list-item class="px-0">
                <v-list-item-content>
                    <v-list-item-title class="text-h4">
                        {{ riskId }}: {{ riskName }}
                    </v-list-item-title>
                </v-list-item-content>

                <v-spacer></v-spacer>

                <v-list-item-action>
                    <restrict-role-permission-button
                        color="error"
                        :permissions="permissionsForDelete"
                        :org-id="currentOrg.Id"
                        :engagement-id="currentEngagement.Id"
                        @click="showHideDelete = true"
                    >
                        Delete
                    </restrict-role-permission-button>

                    <v-dialog
                        v-model="showHideDelete"
                        persistent
                        max-width="40%"
                    >
                        <confirmation-dialog
                            :confirmation="riskId"
                            :in-progress="deleteInProgress"
                            @do-cancel="showHideDelete=false"
                            @do-confirm="onDeleteRisk"
                        >
                            <p>
                                You are about to delete the risk <b>{{ riskId }} : {{ riskName }}</b>.
                                This action is not reversible.
                            </p>
                        </confirmation-dialog>
                    </v-dialog>
                </v-list-item-action>
            </v-list-item>
            <v-divider></v-divider>

            <v-tabs>
                <v-tab :to="overviewTo">Overview</v-tab>
            </v-tabs>
            <router-view></router-view>
        </template>
    </scoping-template>
</template>

<script lang="ts">

import Vue from 'vue'
import Component from 'vue-class-component'
import ScopingTemplate from '@client/vue/orgs/engagements/scoping/ScopingTemplate.vue'
import RestrictRolePermissionButton from '@client/vue/loading/RestrictRolePermissionButton.vue'
import { Permission } from '@client/ts/types/roles'
import { RawRisk } from '@client/ts/types/risks'
import { RawOrganization } from '@client/ts/types/orgs'
import { GrchiveApi } from '@client/ts/main'
import { RawEngagement } from '@client/ts/types/engagements'
import ConfirmationDialog from '@client/vue/shared/ConfirmationDialog.vue'

@Component({
    components: {
        ScopingTemplate,
        RestrictRolePermissionButton,
        ConfirmationDialog,
    }
})
export default class RiskPage extends Vue {
    showHideDelete : boolean = false
    deleteInProgress: boolean = false

    get currentOrg() : RawOrganization | null {
        return this.$store.state.org.rawOrg
    }

    get currentEngagement() : RawEngagement | null {
        return this.$store.state.engagements.rawEngagement
    }

    get currentRisk() : RawRisk | null {
        return this.$store.state.risks.rawRisk
    }

    get riskId() : string {
        if (!this.currentRisk) {
            return ''
        }
        return this.currentRisk.HumanId
    }

    get riskName() : string {
        if (!this.currentRisk) {
            return ''
        }
        return this.currentRisk.Name
    }

    get permissionsForDelete() : Permission[] {
        return [Permission.PRisksDelete]
    }

    get overviewTo() : any {
        return {
            name: 'riskOverview',
            params: this.$route.params,
        }
    }

    onDeleteRisk() {
        this.deleteInProgress = true
        GrchiveApi.risks.deleteRisk(this.currentOrg!.Id, this.currentEngagement!.Id, this.currentRisk!.Id).then((resp : void | null) => {
            if (resp !== null) {
                this.$router.replace({
                    name: 'scopingRisks',
                    params: this.$route.params,
                })
            }
        }).finally(() => {
            this.deleteInProgress = false
        })
    }

    refreshData() {
        const orgId : number = Number(this.$route.params.orgId)
        const engId : number = Number(this.$route.params.engId)
        const riskId : number = Number(this.$route.params.riskId)
        this.$store.dispatch('initializeCurrentResource', {
            orgId,
            engagementId: engId,
            riskId,
        })
    }

    mounted() {
        this.refreshData()
    }
}

</script>

