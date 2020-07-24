<template>
    <scoping-template
        :page-name="generalLedgerFullId"
        disable-default-init
    >
        <template v-slot:content>
            <loading-container
                :loading="isLoading"
            >
                <template v-slot:default="{show}">
                    <div v-if="show">
                        <gl-account-breadcrumbs
                            :acc="currentGeneralLedgerAccount"
                            class="mt-3"
                        >
                        </gl-account-breadcrumbs>
                        <v-list-item class="px-0">
                            <v-list-item-content>
                                <v-list-item-title class="text-h4">
                                    {{ generalLedgerFullId }}
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
                                        :confirmation="currentGeneralLedgerAccount.AccountId"
                                        :in-progress="deleteInProgress"
                                        @do-cancel="showHideDelete=false"
                                        @do-confirm="onDeleteAccount"
                                    >
                                        <p>
                                            You are about to delete the account <b>{{ generalLedgerFullId }}}</b>.
                                            This action is not reversible.
                                        </p>
                                    </confirmation-dialog>
                                </v-dialog>
                            </v-list-item-action>
                        </v-list-item>
                        <v-divider></v-divider>

                        <v-tabs>
                            <v-tab :to="overviewTo">Overview</v-tab>
                            <v-tab :to="subaccountsTo">Subaccounts</v-tab>
                            <restrict-role-permission-tab
                                :permissions="commentPermissions"
                                :to="commentsTo"
                                :org-id="currentOrg.Id"
                                :engagement-id="currentEngagement.Id"
                            >
                                Comments
                            </restrict-role-permission-tab>
                        </v-tabs>
                        <router-view></router-view>
                    </div>
                </template>
            </loading-container>
        </template>
    </scoping-template>
</template>

<script lang="ts">

import Vue from 'vue'
import Component from 'vue-class-component'
import { Watch } from 'vue-property-decorator'
import ScopingTemplate from '@client/vue/orgs/engagements/scoping/ScopingTemplate.vue'
import RestrictRolePermissionButton from '@client/vue/loading/RestrictRolePermissionButton.vue'
import RestrictRolePermissionTab from '@client/vue/loading/RestrictRolePermissionTab.vue'
import { Permission } from '@client/ts/types/roles'
import { RawGLAccount } from '@client/ts/types/gl'
import { RawOrganization } from '@client/ts/types/orgs'
import { GrchiveApi, ErrorHandler } from '@client/ts/main'
import { RawEngagement } from '@client/ts/types/engagements'
import ConfirmationDialog from '@client/vue/shared/ConfirmationDialog.vue'
import LoadingContainer from '@client/vue/loading/LoadingContainer.vue'
import GlAccountBreadcrumbs from '@client/vue/types/gl/GlAccountBreadcrumbs.vue'

@Component({
    components: {
        ScopingTemplate,
        RestrictRolePermissionButton,
        RestrictRolePermissionTab,
        ConfirmationDialog,
        LoadingContainer,
        GlAccountBreadcrumbs
    }
})
export default class GeneralLedgerAccountPage extends Vue {
    showHideDelete : boolean = false
    deleteInProgress: boolean = false

    get overviewTo() : any {
        return {
            name: 'glAccOverview',
            params: this.$route.params,
        }
    }

    get subaccountsTo() : any {
        return {
            name: 'glAccSubaccounts',
            params: this.$route.params,
        }
    }

    get commentsTo() : any {
        return {
            name: 'glAccComments',
            params: this.$route.params,
        }
    }

    get isLoading() : boolean {
        return !this.currentOrg || !this.currentEngagement || !this.currentGeneralLedgerAccount
    }

    get currentOrg() : RawOrganization | null {
        return this.$store.state.org.rawOrg
    }

    get currentEngagement() : RawEngagement | null {
        return this.$store.state.engagements.rawEngagement
    }

    get currentGeneralLedgerAccount() : RawGLAccount | null {
        return this.$store.state.gl.rawGLAccount
    }

    get generalLedgerFullId() : string {
        if (!this.currentGeneralLedgerAccount) {
            return 'Loading...'
        }
        return `${this.currentGeneralLedgerAccount.AccountId}: ${this.currentGeneralLedgerAccount.Name}`
    }

    get commentPermissions() : Permission[] {
        return [Permission.PGLView, Permission.PCommentsList, Permission.POrgUsersView]
    }

    get permissionsForDelete() : Permission[] {
        return [Permission.PGLDelete]
    }

    @Watch('$route')
    refreshData() {
        const orgId : number = Number(this.$route.params.orgId)
        const engId : number = Number(this.$route.params.engId)
        const glAccountId: number = Number(this.$route.params.accId)
        this.$store.dispatch('initializeCurrentResource', {
            orgId,
            engagementId: engId,
            glAccountId,
        })
    }

    mounted() {
        this.refreshData()
    }

    onDeleteAccount() {
        this.deleteInProgress = true
        GrchiveApi.gl.deleteAccount(this.currentOrg!.Id, this.currentEngagement!.Id, this.currentGeneralLedgerAccount!.Id).then(() => {
            this.$router.replace({
                name: 'glHome',
                params: this.$route.params,
            })
        }).catch((err : any) => {
            ErrorHandler.failurePopupOnError(err, {
                context: 'Failed to delete general ledger account.'
            })
        }).finally(() => {
            this.deleteInProgress = false
        })
    }
}

</script>

