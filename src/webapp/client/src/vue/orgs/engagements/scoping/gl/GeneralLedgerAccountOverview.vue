<template>
    <general-ledger-account-save-edit-dialog
        :value="currentGeneralLedgerAccount"
        edit-mode
        :parent-engagement-id="currentEngagement.Id"
        :parent-org-id="currentOrg.Id"
        @save-edit="onEdit"
    >
    </general-ledger-account-save-edit-dialog>
</template>

<script lang="ts">

import Vue from 'vue'
import Component from 'vue-class-component'
import { RawGLAccount } from '@client/ts/types/gl'
import { RawOrganization } from '@client/ts/types/orgs'
import { RawEngagement } from '@client/ts/types/engagements'
import GeneralLedgerAccountSaveEditDialog from '@client/vue/types/gl/GeneralLedgerAccountSaveEditDialog.vue'

@Component({
    components: {
        GeneralLedgerAccountSaveEditDialog
    }
})
export default class GeneralLedgerAccountOverview extends Vue {
    get currentGeneralLedgerAccount() : RawGLAccount | null {
        return this.$store.state.gl.rawGLAccount
    }

    get currentOrg() : RawOrganization | null {
        return this.$store.state.org.rawOrg
    }

    get currentEngagement() : RawEngagement | null {
        return this.$store.state.engagements.rawEngagement
    }

    onEdit(r : RawGLAccount) {
        this.$store.commit('gl/setRawGeneralLedger', r)
    }
}

</script>
