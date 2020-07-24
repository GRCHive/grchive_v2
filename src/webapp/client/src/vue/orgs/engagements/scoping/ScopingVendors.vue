<template>
    <scoping-template
        page-name="Vendors"
    >
        <template v-slot:content>
            <v-list-item class="px-0">
                <v-list-item-content>
                    <v-list-item-title class="text-h4">
                        Vendors
                    </v-list-item-title>
                </v-list-item-content>

                <v-spacer></v-spacer>

                <v-list-item-action>
                    <restrict-role-permission-button
                        color="primary"
                        :permissions="permissionsForCreate"
                        :org-id="currentOrg.Id"
                        :engagement-id="currentEngagement.Id"
                        @click="showHideNew = true"
                    >
                        New
                    </restrict-role-permission-button>

                    <v-dialog
                        v-model="showHideNew"
                        persistent
                        max-width="40%"
                    >
                        <vendor-save-edit-dialog
                            :parent-org-id="currentOrg.Id"
                            :parent-engagement-id="currentEngagement.Id"
                            @cancel-edit="showHideNew = false"
                            @save-edit="onSaveVendor"
                        >
                        </vendor-save-edit-dialog>
                    </v-dialog>
                </v-list-item-action>
            </v-list-item>
            <v-divider class="mb-4"></v-divider>

            <loading-container
                :loading="!allVendors"
            >
                <template v-slot:default="{show}">
                    <full-height-base>
                        <vendor-grid
                            v-if="show"
                            :vendors="allVendors"
                            style="height: 100%;"
                        >
                        </vendor-grid>
                    </full-height-base>
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
import { Permission } from '@client/ts/types/roles'
import { RawOrganization } from '@client/ts/types/orgs'
import { RawEngagement } from '@client/ts/types/engagements'
import { RawVendor } from '@client/ts/types/vendors'
import { GrchiveApi, ErrorHandler } from '@client/ts/main'
import LoadingContainer from '@client/vue/loading/LoadingContainer.vue'
import FullHeightBase from '@client/vue/shared/FullHeightBase.vue'
import VendorSaveEditDialog from '@client/vue/types/vendors/VendorSaveEditDialog.vue'
import VendorGrid from '@client/vue/types/vendors/VendorGrid.vue'

@Component({
    components: {
        ScopingTemplate,
        RestrictRolePermissionButton,
        LoadingContainer,
        FullHeightBase,
        VendorSaveEditDialog,
        VendorGrid,
    }
})
export default class ScopingVendors extends Vue {
    get currentOrg() : RawOrganization | null {
        return this.$store.state.org.rawOrg
    }

    get currentEngagement() : RawEngagement | null {
        return this.$store.state.engagements.rawEngagement
    }

    showHideNew: boolean = false
    allVendors : RawVendor[] | null = null

    get permissionsForCreate() : Permission[] {
        return [Permission.PVendorsCreate]
    }

    onSaveVendor(r : RawVendor) {
        this.allVendors!.unshift(r)
        this.showHideNew = false
    }

    @Watch('currentOrg')
    @Watch('currentEngagement')
    refreshData() {
        if (!this.currentOrg || !this.currentEngagement) {
            return
        }

        GrchiveApi.vendors.listVendors(this.currentOrg.Id, this.currentEngagement.Id).then((resp : RawVendor[]) => {
            this.allVendors = resp
        }).catch((err : any) => {
            ErrorHandler.failurePageOnError(err)
        })
    }

    mounted() {
        this.refreshData()
    }
}

</script>
