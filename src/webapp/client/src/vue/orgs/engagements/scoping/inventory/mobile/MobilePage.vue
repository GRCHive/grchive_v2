<template>
    <scoping-template
        :page-name="mobileName"
        disable-default-init
    >
        <template v-slot:content>
            <v-list-item class="px-0">
                <v-list-item-content>
                    <v-list-item-title class="text-h4">
                        {{ mobileName }}
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
                            :confirmation="mobileName"
                            :in-progress="deleteInProgress"
                            @do-cancel="showHideDelete=false"
                            @do-confirm="onDeleteMobile"
                        >
                            <p>
                                You are about to delete the mobile <b>{{ mobileName }}</b>.
                                This action is not reversible.
                            </p>
                        </confirmation-dialog>
                    </v-dialog>
                </v-list-item-action>
            </v-list-item>
            <v-divider></v-divider>

            <loading-container
                :loading="isLoading"
            >
                <template v-slot:default="{show}">
                    <div v-if="show">
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
import { RawMobile, InventoryType } from '@client/ts/types/inventory'
import { RawOrganization } from '@client/ts/types/orgs'
import { GrchiveApi, ErrorHandler } from '@client/ts/main'
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
export default class MobilePage extends Vue {
    showHideDelete : boolean = false
    deleteInProgress: boolean = false

    get isLoading() : boolean {
        return !this.currentOrg || !this.currentEngagement || !this.currentMobile
    }

    get currentOrg() : RawOrganization | null {
        return this.$store.state.org.rawOrg
    }

    get currentEngagement() : RawEngagement | null {
        return this.$store.state.engagements.rawEngagement
    }

    get currentMobile() : RawMobile | null {
        return this.$store.state.inventory.rawMobile
    }

    get mobileName() : string {
        if (!this.currentMobile) {
            return ''
        }
        return this.currentMobile.Inventory.Name
    }

    get permissionsForDelete() : Permission[] {
        return [Permission.PMobileDelete]
    }

    get overviewTo() : any {
        return {
            name: 'mobileOverview',
            params: this.$route.params,
        }
    }

    get commentsTo() : any {
        return {
            name: 'mobileComments',
            params: this.$route.params,
        }
    }

    get commentPermissions() : Permission[] {
        return [Permission.PMobileView, Permission.PCommentsList, Permission.POrgUsersView]
    }

    onDeleteMobile() {
        this.deleteInProgress = true
        GrchiveApi.inventory.deleteInventory(InventoryType.ITMobile, this.currentOrg!.Id, this.currentEngagement!.Id, this.currentMobile!.Id).then(() => {
            this.$router.replace({
                name: 'scopingMobile',
                params: this.$route.params,
            })
        }).catch((err : any) => {
            ErrorHandler.failurePopupOnError(err, {
                context: 'Failed to delete mobile.'
            })
        }).finally(() => {
            this.deleteInProgress = false
        })
    }

    @Watch('$route')
    refreshData() {
        const orgId : number = Number(this.$route.params.orgId)
        const engId : number = Number(this.$route.params.engId)
        const mobileId : number = Number(this.$route.params.mobileId)
        this.$store.dispatch('initializeCurrentResource', {
            orgId,
            engagementId: engId,
            mobileId,
        })
    }

    mounted() {
        this.refreshData()
    }
}

</script>
