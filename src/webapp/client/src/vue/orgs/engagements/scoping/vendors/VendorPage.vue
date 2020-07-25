<template>
    <scoping-template
        :page-name="vendorName"
        disable-default-init
    >
        <template v-slot:content>
            <v-list-item class="px-0">
                <v-list-item-content>
                    <v-list-item-title class="text-h4">
                        {{ vendorName }}
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
                            :confirmation="vendorName"
                            :in-progress="deleteInProgress"
                            @do-cancel="showHideDelete=false"
                            @do-confirm="onDeleteVendor"
                        >
                            <p>
                                You are about to delete the vendor <b>{{ vendorName }}</b>.
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
                            <v-tab :to="productsTo">Products</v-tab>
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
import { RawVendor } from '@client/ts/types/vendors'
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
export default class VendorPage extends Vue {
    showHideDelete : boolean = false
    deleteInProgress: boolean = false

    get isLoading() : boolean {
        return !this.currentOrg || !this.currentEngagement || !this.currentVendor
    }

    get currentOrg() : RawOrganization | null {
        return this.$store.state.org.rawOrg
    }

    get currentEngagement() : RawEngagement | null {
        return this.$store.state.engagements.rawEngagement
    }

    get currentVendor() : RawVendor | null {
        return this.$store.state.vendors.rawVendor
    }

    get vendorName() : string {
        if (!this.currentVendor) {
            return ''
        }
        return this.currentVendor.Name
    }

    get permissionsForDelete() : Permission[] {
        return [Permission.PVendorsDelete]
    }

    get overviewTo() : any {
        return {
            name: 'vendorOverview',
            params: this.$route.params,
        }
    }

    get productsTo() : any {
        return {
            name: 'vendorProducts',
            params: this.$route.params,
            permissions: [Permission.PVendorProductsList]
        }
    }

    get commentsTo() : any {
        return {
            name: 'vendorComments',
            params: this.$route.params,
        }
    }

    get commentPermissions() : Permission[] {
        return [Permission.PVendorsView, Permission.PCommentsList, Permission.POrgUsersView]
    }

    onDeleteVendor() {
        this.deleteInProgress = true
        GrchiveApi.vendors.deleteVendor(this.currentOrg!.Id, this.currentEngagement!.Id, this.currentVendor!.Id).then(() => {
            this.$router.replace({
                name: 'scopingVendors',
                params: this.$route.params,
            })
        }).catch((err : any) => {
            ErrorHandler.failurePopupOnError(err, {
                context: 'Failed to delete vendor.'
            })
        }).finally(() => {
            this.deleteInProgress = false
        })
    }

    @Watch('$route')
    refreshData() {
        const orgId : number = Number(this.$route.params.orgId)
        const engId : number = Number(this.$route.params.engId)
        const vendorId : number = Number(this.$route.params.vendorId)
        this.$store.dispatch('initializeCurrentResource', {
            orgId,
            engagementId: engId,
            vendorId,
        })
    }

    mounted() {
        this.refreshData()
    }
}

</script>

