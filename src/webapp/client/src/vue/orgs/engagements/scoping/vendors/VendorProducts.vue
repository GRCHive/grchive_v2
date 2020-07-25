<template>
    <loading-container
        :loading="!allProducts"
    >
        <template v-slot:default="{show}">
            <v-row v-if="show">
                <v-col cols="2">
                    <v-list>
                        <v-dialog
                            v-model="showHideNew"
                            persistent
                            max-width="40%"
                        >
                            <template v-slot:activator="{on}">
                                <restrict-role-permission-container
                                    :permissions="newPermissions"
                                    :org-id="currentOrg.Id"
                                    :engagement-id="currentEngagement.Id"
                                >
                                    <v-list-item v-on="on">
                                        <v-list-item-icon>
                                            <v-icon>mdi-plus</v-icon>
                                        </v-list-item-icon>

                                        <v-list-item-content>
                                            New Product
                                        </v-list-item-content>
                                    </v-list-item>
                                </restrict-role-permission-container>
                            </template>

                            <vendor-product-save-edit-dialog
                                :parent-engagement-id="currentEngagement.Id"
                                :parent-org-id="currentOrg.Id"
                                :parent-vendor-id="currentVendor.Id"
                                @cancel-edit="showHideNew = false"
                                @save-edit="onSaveNewProduct"
                            >
                            </vendor-product-save-edit-dialog>
                        </v-dialog>

                        <v-list-item-group>
                            <v-list-item
                                v-for="(product, i) in allProducts"
                                :key="product.Id"
                                :to="createProductTo(product.Id)"
                            >
                                <v-list-item-content>
                                    <v-list-item-title>
                                        {{ product.Name }}
                                    </v-list-item-title>
                                </v-list-item-content>
                            </v-list-item>
                        </v-list-item-group>
                    </v-list>
                </v-col>

                <v-col cols="10">
                    <restrict-role-permission-container
                        :permissions="viewPermissions"
                        :org-id="currentOrg.Id"
                        :engagement-id="currentEngagement.Id"
                    >
                        <template v-slot:default="{show}">
                            <router-view v-if="show"></router-view>
                            <full-height-base v-else>
                                <div></div>
                            </full-height-base>
                        </template>
                    </restrict-role-permission-container>
                </v-col>
            </v-row>
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
import LoadingContainer from '@client/vue/loading/LoadingContainer.vue'
import VendorProductSaveEditDialog from '@client/vue/types/vendors/VendorProductSaveEditDialog.vue'
import RestrictRolePermissionContainer from '@client/vue/loading/RestrictRolePermissionContainer.vue'
import FullHeightBase from '@client/vue/shared/FullHeightBase.vue'

@Component({
    components: {
        LoadingContainer,
        VendorProductSaveEditDialog,
        RestrictRolePermissionContainer,
        FullHeightBase
    }
})
export default class VendorProducts extends Vue {
    showHideNew: boolean = false
    allProducts : RawVendorProduct[] | null = null

    get currentVendor() : RawVendor | null {
        return this.$store.state.vendors.rawVendor
    }

    get currentOrg() : RawOrganization | null {
        return this.$store.state.org.rawOrg
    }

    get currentEngagement() : RawEngagement | null {
        return this.$store.state.engagements.rawEngagement
    }

    get currentVendorProduct() : RawVendorProduct | null {
        return this.$store.state.vendors.rawVendorProduct
    }

    get newPermissions() : Permission[] {
        return [Permission.PVendorProductsCreate]
    }

    get viewPermissions() : Permission[] {
        return [Permission.PVendorProductsView]
    }

    // This is needed to sync changes from the VendorProductOverview page and this page since the allProducts array
    // isn't available for edit by the VendorProductOverview page.
    @Watch('currentVendorProduct')
    updateCurrentVendorProduct(newVal : RawVendorProduct | null, oldVal : RawVendorProduct | null) {
        if (!this.allProducts) {
            return
        }

        if (!this.currentVendorProduct) {
            // Need to detect a delete. We do this using the query parameter that gets set when something is deleted. Hacky....
            if (!!this.$route.query.isDelete && !!oldVal) {
                let idx : number = this.allProducts.findIndex((ele : RawVendorProduct) => ele.Id == oldVal.Id)
                if (idx == -1) {
                    return
                }
                this.allProducts.splice(idx, 1)
            }
        } else {
            let idx : number = this.allProducts.findIndex((ele : RawVendorProduct) => ele.Id == this.currentVendorProduct!.Id)
            if (idx == -1) {
                return
            }

            Vue.set(this.allProducts, idx, this.currentVendorProduct)
        }
    }

    @Watch('currentVendor')
    @Watch('currentOrg')
    @Watch('currentEngagement')
    refreshData() {
        if (!this.currentVendor || !this.currentOrg || !this.currentEngagement) {
            this.allProducts = null
            return
        }

        GrchiveApi.vendors.listVendorProducts(this.currentOrg.Id, this.currentEngagement.Id, this.currentVendor.Id).then((resp : RawVendorProduct[]) => {
            this.allProducts = resp
        }).catch((err : any) => {
            ErrorHandler.failurePageOnError(err)
        })
    }

    mounted() {
        this.refreshData()
    }

    onSaveNewProduct(p : RawVendorProduct) {
        this.showHideNew = false
        if (!this.allProducts) {
            return
        }
        this.allProducts.unshift(p)
    }

    createProductTo(pId : number) : any {
        return {
            name: 'vendorProductOverview',
            params: {
                ...this.$route.params,
                productId: `${pId}`
            }
        }
    }
}

</script>
