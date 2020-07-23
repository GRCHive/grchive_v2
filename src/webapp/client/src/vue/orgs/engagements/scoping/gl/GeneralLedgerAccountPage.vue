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
import { GrchiveApi } from '@client/ts/main'
import { RawEngagement } from '@client/ts/types/engagements'
import ConfirmationDialog from '@client/vue/shared/ConfirmationDialog.vue'
import LoadingContainer from '@client/vue/loading/LoadingContainer.vue'

@Component({
    components: {
        ScopingTemplate,
        RestrictRolePermissionButton,
        RestrictRolePermissionTab,
        ConfirmationDialog,
        LoadingContainer
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
        GrchiveApi.gl.deleteAccount(this.currentOrg!.Id, this.currentEngagement!.Id, this.currentGeneralLedgerAccount!.Id).then((resp : void | null) => {
            if (resp !== null) {
                this.$router.replace({
                    name: 'glHome',
                    params: this.$route.params,
                })
            }
        }).finally(() => {
            this.deleteInProgress = false
        })
    }
}

</script>

