<template>
    <vendor-save-edit-dialog
        :value="currentVendor"
        :parent-org-id="currentOrg.Id"
        :parent-engagement-id="currentEngagement.Id"
        edit-mode
        @save-edit="onEditVendor"
    >
    </vendor-save-edit-dialog>
</template>

<script lang="ts">

import Vue from 'vue'
import Component from 'vue-class-component'
import VendorSaveEditDialog from '@client/vue/types/vendors/VendorSaveEditDialog.vue'
import { RawVendor } from '@client/ts/types/vendors'
import { RawOrganization } from '@client/ts/types/orgs'
import { RawEngagement } from '@client/ts/types/engagements'

@Component({
    components: {
        VendorSaveEditDialog,
    }
})
export default class VendorOverview extends Vue {
    get currentVendor() : RawVendor | null {
        return this.$store.state.vendors.rawVendor
    }

    get currentOrg() : RawOrganization | null {
        return this.$store.state.org.rawOrg
    }

    get currentEngagement() : RawEngagement | null {
        return this.$store.state.engagements.rawEngagement
    }

    onEditVendor(r : RawVendor) {
        this.$store.commit('vendors/setRawVendor', r)
    }
}

</script>
