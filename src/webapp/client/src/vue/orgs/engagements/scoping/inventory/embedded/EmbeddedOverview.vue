<template>
    <embedded-save-edit-dialog
        :value="currentEmbedded"
        :parent-org-id="currentOrg.Id"
        :parent-engagement-id="currentEngagement.Id"
        edit-mode
        @save-edit="onEditEmbedded"
    >
    </embedded-save-edit-dialog>
</template>

<script lang="ts">

import Vue from 'vue'
import Component from 'vue-class-component'
import EmbeddedSaveEditDialog from '@client/vue/types/inventory/embedded/EmbeddedSaveEditDialog.vue'
import { RawEmbedded } from '@client/ts/types/inventory'
import { RawOrganization } from '@client/ts/types/orgs'
import { RawEngagement } from '@client/ts/types/engagements'

@Component({
    components: {
        EmbeddedSaveEditDialog,
    }
})
export default class EmbeddedOverview extends Vue {
    get currentEmbedded() : RawEmbedded | null {
        return this.$store.state.inventory.rawEmbedded
    }

    get currentOrg() : RawOrganization | null {
        return this.$store.state.org.rawOrg
    }

    get currentEngagement() : RawEngagement | null {
        return this.$store.state.engagements.rawEngagement
    }

    onEditEmbedded(r : RawEmbedded) {
        this.$store.commit('inventory/setRawEmbedded', r)
    }
}

</script>
