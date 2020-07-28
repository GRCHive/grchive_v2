<template>
    <generic-crud-list
        label="Desktops"
        :permissions-for-create="permissionsForCreate"
        :all-data="allDesktops"
        @refresh-list="refreshData"
        @new-item="onNewDesktop"
    >
        <template v-slot:create="{ onCancel, onSave, orgId, engagementId }">
            <desktop-save-edit-dialog
                :parent-engagement-id="engagementId"
                :parent-org-id="orgId"
                @cancel-edit="onCancel"
                @save-edit="onSave"
            >
            </desktop-save-edit-dialog>
        </template>

        <template v-slot:list="{show}">
            <desktop-grid
                v-if="show"
                :desktops="allDesktops"
                style="height: 100%;"
            >
            </desktop-grid>
        </template>
    </generic-crud-list>
</template>

<script lang="ts">

import Vue from 'vue'
import Component from 'vue-class-component'
import GenericCrudList from '@client/vue/shared/GenericCrudList.vue'
import DesktopSaveEditDialog from '@client/vue/types/inventory/desktops/DesktopSaveEditDialog.vue'
import DesktopGrid from '@client/vue/types/inventory/desktops/DesktopGrid.vue'

import { Permission } from '@client/ts/types/roles'
import { GrchiveApi, ErrorHandler } from '@client/ts/main'
import { RawDesktop, InventoryType } from '@client/ts/types/inventory'

@Component({
    components: {
        GenericCrudList,
        DesktopSaveEditDialog,
        DesktopGrid
    }
})
export default class ScopingDesktopList extends Vue {
    allDesktops: RawDesktop[] | null = null

    get permissionsForCreate() : Permission[] {
        return [Permission.PDesktopsCreate]
    }

    refreshData(orgId : number | undefined, engagementId : number | undefined) {
        if (orgId === undefined || engagementId === undefined) {
            this.allDesktops = null
            return
        }
        GrchiveApi.inventory.listInventory<RawDesktop>(InventoryType.ITDesktop, orgId, engagementId).then((resp: RawDesktop[]) => {
            this.allDesktops = resp
        }).catch((err : any) => {
            ErrorHandler.failurePageOnError(err)
        })
    }

    onNewDesktop(s : RawDesktop) {
        if (!this.allDesktops) {
            return
        }
        this.allDesktops.unshift(s)
    }
}

</script>
