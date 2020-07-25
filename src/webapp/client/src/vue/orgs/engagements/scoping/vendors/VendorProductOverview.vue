<template>
    <loading-container
        :loading="!currentVendorProduct"
    >
        <template v-slot:default="{show}">
            <template v-if="show">
                <v-list-item>
                    <v-list-item-content>
                        <v-list-item-title class="text-h4">
                            {{ currentVendorProduct.Name }}
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
                                :confirmation="currentVendorProduct.Name"
                                :in-progress="deleteInProgress"
                                @do-cancel="showHideDelete=false"
                                @do-confirm="onDeleteVendorProduct"
                            >
                                <p>
                                    You are about to delete the vendor product <b>{{ currentVendorProduct.Name }}</b>.
                                    This action is not reversible.
                                </p>
                            </confirmation-dialog>
                        </v-dialog>
                    </v-list-item-action>
                </v-list-item>
                <v-divider></v-divider>

                <vendor-product-save-edit-dialog
                    :value="currentVendorProduct"
                    edit-mode
                    :parent-engagement-id="currentEngagement.Id"
                    :parent-org-id="currentOrg.Id"
                    :parent-vendor-id="currentVendor.Id"
                    @save-edit="onEditProduct"
                >
                </vendor-product-save-edit-dialog>
            </template>
        </template>
    </loading-container>
</template>

<script lang="ts">

import Vue from 'vue'
import Component from 'vue-class-component'
import { Watch } from 'vue-property-decorator'
import { RawVendor, RawVendorProduct } from '@client/ts/types/vendors'
import { RawOrganization } from '@client/ts/types/orgs'
import { RawEngagement } from '@client/ts/types/engagements'
import { GrchiveApi, ErrorHandler } from '@client/ts/main'
import { Permission } from '@client/ts/types/roles'
import VendorProductSaveEditDialog from '@client/vue/types/vendors/VendorProductSaveEditDialog.vue'
import LoadingContainer from '@client/vue/loading/LoadingContainer.vue'
import RestrictRolePermissionButton from '@client/vue/loading/RestrictRolePermissionButton.vue'
import ConfirmationDialog from '@client/vue/shared/ConfirmationDialog.vue'

@Component({
    components: {
        VendorProductSaveEditDialog,
        LoadingContainer,
        RestrictRolePermissionButton,
        ConfirmationDialog
    }
})
export default class VendorProductOverview extends Vue {
    showHideDelete: boolean = false
    deleteInProgress : boolean = false

    get currentVendor() : RawVendor | null {
        return this.$store.state.vendors.rawVendor
    }

    get currentVendorProduct() : RawVendorProduct | null {
        return this.$store.state.vendors.rawVendorProduct
    }

    get currentOrg() : RawOrganization | null {
        return this.$store.state.org.rawOrg
    }

    get currentEngagement() : RawEngagement | null {
        return this.$store.state.engagements.rawEngagement
    }

    onEditProduct(p : RawVendorProduct) {
        this.$store.commit('vendors/setRawVendorProduct', p)
    }

    get permissionsForDelete() : Permission[] {
        return [Permission.PVendorProductsDelete]
    }

    onDeleteVendorProduct() {
        this.deleteInProgress = true
        GrchiveApi.vendors.deleteVendorProduct(this.currentOrg!.Id, this.currentEngagement!.Id, this.currentVendor!.Id, this.currentVendorProduct!.Id).then(() => {
            this.$router.replace({
                name: 'vendorProducts',
                params: this.$route.params,
                query: {
                    isDelete: 'yes',
                }
            })
        }).catch((err : any) => {
            ErrorHandler.failurePopupOnError(err, {
                context: 'Failed to delete vendor product.'
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
        const productId : number = Number(this.$route.params.productId)
        this.$store.dispatch('initializeCurrentResource', {
            orgId,
            engagementId: engId,
            vendorId,
            vendorProductId: productId
        })
    }

    mounted() {
        this.refreshData()
    }
}

</script>
