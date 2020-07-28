<template>
    <laptop-save-edit-dialog
        :value="currentLaptop"
        :parent-org-id="currentOrg.Id"
        :parent-engagement-id="currentEngagement.Id"
        edit-mode
        @save-edit="onEditLaptop"
    >
    </laptop-save-edit-dialog>
</template>

<script lang="ts">

import Vue from 'vue'
import Component from 'vue-class-component'
import LaptopSaveEditDialog from '@client/vue/types/inventory/laptops/LaptopSaveEditDialog.vue'
import { RawLaptop } from '@client/ts/types/inventory'
import { RawOrganization } from '@client/ts/types/orgs'
import { RawEngagement } from '@client/ts/types/engagements'

@Component({
    components: {
        LaptopSaveEditDialog,
    }
})
export default class LaptopOverview extends Vue {
    get currentLaptop() : RawLaptop | null {
        return this.$store.state.inventory.rawLaptop
    }

    get currentOrg() : RawOrganization | null {
        return this.$store.state.org.rawOrg
    }

    get currentEngagement() : RawEngagement | null {
        return this.$store.state.engagements.rawEngagement
    }

    onEditLaptop(r : RawLaptop) {
        this.$store.commit('inventory/setRawLaptop', r)
    }
}

</script>
