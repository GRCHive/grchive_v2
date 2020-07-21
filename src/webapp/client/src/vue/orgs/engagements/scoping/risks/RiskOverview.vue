<template>
    <risk-save-edit-dialog
        :value="currentRisk"
        :parent-org-id="currentOrg.Id"
        :parent-engagement-id="currentEngagement.Id"
        edit-mode
        @save-edit="onEditRisk"
    >
    </risk-save-edit-dialog>
</template>

<script lang="ts">

import Vue from 'vue'
import Component from 'vue-class-component'
import RiskSaveEditDialog from '@client/vue/types/risks/RiskSaveEditDialog.vue'
import { RawRisk } from '@client/ts/types/risks'
import { RawOrganization } from '@client/ts/types/orgs'
import { RawEngagement } from '@client/ts/types/engagements'

@Component({
    components: {
        RiskSaveEditDialog,
    }
})
export default class RiskOverview extends Vue {
    get currentRisk() : RawRisk | null {
        return this.$store.state.risks.rawRisk
    }

    get currentOrg() : RawOrganization | null {
        return this.$store.state.org.rawOrg
    }

    get currentEngagement() : RawEngagement | null {
        return this.$store.state.engagements.rawEngagement
    }

    onEditRisk(r : RawRisk) {
        this.$store.commit('risks/setRawRisk', r)
    }
}

</script>
