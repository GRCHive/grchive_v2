<template>
    <desktop-save-edit-dialog
        :value="currentDesktop"
        :parent-org-id="currentOrg.Id"
        :parent-engagement-id="currentEngagement.Id"
        edit-mode
        @save-edit="onEditDesktop"
    >
    </desktop-save-edit-dialog>
</template>

<script lang="ts">

import Vue from 'vue'
import Component from 'vue-class-component'
import DesktopSaveEditDialog from '@client/vue/types/inventory/desktops/DesktopSaveEditDialog.vue'
import { RawDesktop } from '@client/ts/types/inventory'
import { RawOrganization } from '@client/ts/types/orgs'
import { RawEngagement } from '@client/ts/types/engagements'

@Component({
    components: {
        DesktopSaveEditDialog,
    }
})
export default class DesktopOverview extends Vue {
    get currentDesktop() : RawDesktop | null {
        return this.$store.state.inventory.rawDesktop
    }

    get currentOrg() : RawOrganization | null {
        return this.$store.state.org.rawOrg
    }

    get currentEngagement() : RawEngagement | null {
        return this.$store.state.engagements.rawEngagement
    }

    onEditDesktop(r : RawDesktop) {
        this.$store.commit('inventory/setRawDesktop', r)
    }
}

</script>
