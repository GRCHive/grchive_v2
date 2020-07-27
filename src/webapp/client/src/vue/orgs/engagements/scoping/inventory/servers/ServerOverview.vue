<template>
    <server-save-edit-dialog
        :value="currentServer"
        :parent-org-id="currentOrg.Id"
        :parent-engagement-id="currentEngagement.Id"
        edit-mode
        @save-edit="onEditServer"
    >
    </server-save-edit-dialog>
</template>

<script lang="ts">

import Vue from 'vue'
import Component from 'vue-class-component'
import ServerSaveEditDialog from '@client/vue/types/inventory/servers/ServerSaveEditDialog.vue'
import { RawServer } from '@client/ts/types/inventory'
import { RawOrganization } from '@client/ts/types/orgs'
import { RawEngagement } from '@client/ts/types/engagements'

@Component({
    components: {
        ServerSaveEditDialog,
    }
})
export default class ServerOverview extends Vue {
    get currentServer() : RawServer | null {
        return this.$store.state.inventory.rawServer
    }

    get currentOrg() : RawOrganization | null {
        return this.$store.state.org.rawOrg
    }

    get currentEngagement() : RawEngagement | null {
        return this.$store.state.engagements.rawEngagement
    }

    onEditServer(r : RawServer) {
        this.$store.commit('inventory/setRawServer', r)
    }
}

</script>
