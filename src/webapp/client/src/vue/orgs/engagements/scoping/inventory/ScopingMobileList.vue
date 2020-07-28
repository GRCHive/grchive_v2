<template>
    <generic-crud-list
        label="Mobile Devices"
        :permissions-for-create="permissionsForCreate"
        :all-data="allMobiles"
        @refresh-list="refreshData"
        @new-item="onNewMobile"
    >
        <template v-slot:create="{ onCancel, onSave, orgId, engagementId }">
            <mobile-save-edit-dialog
                :parent-engagement-id="engagementId"
                :parent-org-id="orgId"
                @cancel-edit="onCancel"
                @save-edit="onSave"
            >
            </mobile-save-edit-dialog>
        </template>

        <template v-slot:list="{show}">
            <mobile-grid
                v-if="show"
                :mobile="allMobiles"
                style="height: 100%;"
            >
            </mobile-grid>
        </template>
    </generic-crud-list>
</template>

<script lang="ts">

import Vue from 'vue'
import Component from 'vue-class-component'
import GenericCrudList from '@client/vue/shared/GenericCrudList.vue'
import MobileSaveEditDialog from '@client/vue/types/inventory/mobile/MobileSaveEditDialog.vue'
import MobileGrid from '@client/vue/types/inventory/mobile/MobileGrid.vue'

import { Permission } from '@client/ts/types/roles'
import { GrchiveApi, ErrorHandler } from '@client/ts/main'
import { RawMobile, InventoryType } from '@client/ts/types/inventory'

@Component({
    components: {
        GenericCrudList,
        MobileSaveEditDialog,
        MobileGrid
    }
})
export default class ScopingMobileList extends Vue {
    allMobiles: RawMobile[] | null = null

    get permissionsForCreate() : Permission[] {
        return [Permission.PMobileCreate]
    }

    refreshData(orgId : number | undefined, engagementId : number | undefined) {
        if (orgId === undefined || engagementId === undefined) {
            this.allMobiles = null
            return
        }
        GrchiveApi.inventory.listInventory<RawMobile>(InventoryType.ITMobile, orgId, engagementId).then((resp: RawMobile[]) => {
            this.allMobiles = resp
        }).catch((err : any) => {
            ErrorHandler.failurePageOnError(err)
        })
    }

    onNewMobile(s : RawMobile) {
        if (!this.allMobiles) {
            return
        }
        this.allMobiles.unshift(s)
    }
}

</script>
