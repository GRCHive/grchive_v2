<template>
    <generic-crud-list
        label="Laptops"
        :permissions-for-create="permissionsForCreate"
        :all-data="allLaptops"
        @refresh-list="refreshData"
        @new-item="onNewLaptop"
    >
        <template v-slot:create="{ onCancel, onSave, orgId, engagementId }">
            <laptop-save-edit-dialog
                :parent-engagement-id="engagementId"
                :parent-org-id="orgId"
                @cancel-edit="onCancel"
                @save-edit="onSave"
            >
            </laptop-save-edit-dialog>
        </template>

        <template v-slot:list="{show}">
            <laptop-grid
                v-if="show"
                :laptops="allLaptops"
                style="height: 100%;"
            >
            </laptop-grid>
        </template>
    </generic-crud-list>
</template>

<script lang="ts">

import Vue from 'vue'
import Component from 'vue-class-component'
import GenericCrudList from '@client/vue/shared/GenericCrudList.vue'
import LaptopSaveEditDialog from '@client/vue/types/inventory/laptops/LaptopSaveEditDialog.vue'
import LaptopGrid from '@client/vue/types/inventory/laptops/LaptopGrid.vue'

import { Permission } from '@client/ts/types/roles'
import { GrchiveApi, ErrorHandler } from '@client/ts/main'
import { RawLaptop, InventoryType } from '@client/ts/types/inventory'

@Component({
    components: {
        GenericCrudList,
        LaptopSaveEditDialog,
        LaptopGrid
    }
})
export default class ScopingLaptopList extends Vue {
    allLaptops: RawLaptop[] | null = null

    get permissionsForCreate() : Permission[] {
        return [Permission.PLaptopsCreate]
    }

    refreshData(orgId : number | undefined, engagementId : number | undefined) {
        if (orgId === undefined || engagementId === undefined) {
            this.allLaptops = null
            return
        }
        GrchiveApi.inventory.listInventory<RawLaptop>(InventoryType.ITLaptop, orgId, engagementId).then((resp: RawLaptop[]) => {
            this.allLaptops = resp
        }).catch((err : any) => {
            ErrorHandler.failurePageOnError(err)
        })
    }

    onNewLaptop(s : RawLaptop) {
        if (!this.allLaptops) {
            return
        }
        this.allLaptops.unshift(s)
    }
}

</script>
