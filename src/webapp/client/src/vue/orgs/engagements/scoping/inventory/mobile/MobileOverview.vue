<template>
    <mobile-save-edit-dialog
        :value="currentMobile"
        :parent-org-id="currentOrg.Id"
        :parent-engagement-id="currentEngagement.Id"
        edit-mode
        @save-edit="onEditMobile"
    >
    </mobile-save-edit-dialog>
</template>

<script lang="ts">

import Vue from 'vue'
import Component from 'vue-class-component'
import MobileSaveEditDialog from '@client/vue/types/inventory/mobile/MobileSaveEditDialog.vue'
import { RawMobile } from '@client/ts/types/inventory'
import { RawOrganization } from '@client/ts/types/orgs'
import { RawEngagement } from '@client/ts/types/engagements'

@Component({
    components: {
        MobileSaveEditDialog,
    }
})
export default class MobileOverview extends Vue {
    get currentMobile() : RawMobile | null {
        return this.$store.state.inventory.rawMobile
    }

    get currentOrg() : RawOrganization | null {
        return this.$store.state.org.rawOrg
    }

    get currentEngagement() : RawEngagement | null {
        return this.$store.state.engagements.rawEngagement
    }

    onEditMobile(r : RawMobile) {
        this.$store.commit('inventory/setRawMobile', r)
    }
}

</script>
