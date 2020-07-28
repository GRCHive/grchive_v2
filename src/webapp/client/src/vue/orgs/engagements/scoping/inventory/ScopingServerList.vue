<template>
    <generic-crud-list
        label="Servers"
        :permissions-for-create="permissionsForCreate"
        :all-data="allServers"
        @refresh-list="refreshData"
        @new-item="onNewServer"
    >
        <template v-slot:create="{ onCancel, onSave, orgId, engagementId }">
            <server-save-edit-dialog
                :parent-engagement-id="engagementId"
                :parent-org-id="orgId"
                @cancel-edit="onCancel"
                @save-edit="onSave"
            >
            </server-save-edit-dialog>
        </template>

        <template v-slot:list="{show}">
            <server-grid
                v-if="show"
                :servers="allServers"
                style="height: 100%;"
            >
            </server-grid>
        </template>
    </generic-crud-list>
</template>

<script lang="ts">

import Vue from 'vue'
import Component from 'vue-class-component'
import GenericCrudList from '@client/vue/shared/GenericCrudList.vue'
import ServerSaveEditDialog from '@client/vue/types/inventory/servers/ServerSaveEditDialog.vue'
import ServerGrid from '@client/vue/types/inventory/servers/ServerGrid.vue'

import { Permission } from '@client/ts/types/roles'
import { GrchiveApi, ErrorHandler } from '@client/ts/main'
import { RawServer, InventoryType } from '@client/ts/types/inventory'

@Component({
    components: {
        GenericCrudList,
        ServerSaveEditDialog,
        ServerGrid
    }
})
export default class ScopingServerList extends Vue {
    allServers: RawServer[] | null = null

    get permissionsForCreate() : Permission[] {
        return [Permission.PServersCreate]
    }

    refreshData(orgId : number | undefined, engagementId : number | undefined) {
        if (orgId === undefined || engagementId === undefined) {
            this.allServers = null
            return
        }
        GrchiveApi.inventory.listInventory<RawServer>(InventoryType.ITServer, orgId, engagementId).then((resp: RawServer[]) => {
            this.allServers = resp
        }).catch((err : any) => {
            ErrorHandler.failurePageOnError(err)
        })
    }

    onNewServer(s : RawServer) {
        if (!this.allServers) {
            return
        }
        this.allServers.unshift(s)
    }
}

</script>
