<template>
    <generic-crud-list
        label="Embedded Devices"
        :permissions-for-create="permissionsForCreate"
        :all-data="allEmbeddeds"
        @refresh-list="refreshData"
        @new-item="onNewEmbedded"
    >
        <template v-slot:create="{ onCancel, onSave, orgId, engagementId }">
            <embedded-save-edit-dialog
                :parent-engagement-id="engagementId"
                :parent-org-id="orgId"
                @cancel-edit="onCancel"
                @save-edit="onSave"
            >
            </embedded-save-edit-dialog>
        </template>

        <template v-slot:list="{show}">
            <embedded-grid
                v-if="show"
                :embedded="allEmbeddeds"
                style="height: 100%;"
            >
            </embedded-grid>
        </template>
    </generic-crud-list>
</template>

<script lang="ts">

import Vue from 'vue'
import Component from 'vue-class-component'
import GenericCrudList from '@client/vue/shared/GenericCrudList.vue'
import EmbeddedSaveEditDialog from '@client/vue/types/inventory/embedded/EmbeddedSaveEditDialog.vue'
import EmbeddedGrid from '@client/vue/types/inventory/embedded/EmbeddedGrid.vue'

import { Permission } from '@client/ts/types/roles'
import { GrchiveApi, ErrorHandler } from '@client/ts/main'
import { RawEmbedded, InventoryType } from '@client/ts/types/inventory'

@Component({
    components: {
        GenericCrudList,
        EmbeddedSaveEditDialog,
        EmbeddedGrid
    }
})
export default class ScopingEmbeddedList extends Vue {
    allEmbeddeds: RawEmbedded[] | null = null

    get permissionsForCreate() : Permission[] {
        return [Permission.PEmbeddedCreate]
    }

    refreshData(orgId : number | undefined, engagementId : number | undefined) {
        if (orgId === undefined || engagementId === undefined) {
            this.allEmbeddeds = null
            return
        }
        GrchiveApi.inventory.listInventory<RawEmbedded>(InventoryType.ITEmbedded, orgId, engagementId).then((resp: RawEmbedded[]) => {
            this.allEmbeddeds = resp
        }).catch((err : any) => {
            ErrorHandler.failurePageOnError(err)
        })
    }

    onNewEmbedded(s : RawEmbedded) {
        if (!this.allEmbeddeds) {
            return
        }
        this.allEmbeddeds.unshift(s)
    }
}

</script>
