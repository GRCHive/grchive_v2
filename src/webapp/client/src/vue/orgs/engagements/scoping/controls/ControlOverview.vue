<template>
    <control-save-edit-dialog
        :value="currentControl"
        :parent-org-id="currentOrg.Id"
        :parent-engagement-id="currentEngagement.Id"
        edit-mode
        @save-edit="onEditControl"
    >
    </control-save-edit-dialog>
</template>

<script lang="ts">

import Vue from 'vue'
import Component from 'vue-class-component'
import ControlSaveEditDialog from '@client/vue/types/controls/ControlSaveEditDialog.vue'
import { RawControl } from '@client/ts/types/controls'
import { RawOrganization } from '@client/ts/types/orgs'
import { RawEngagement } from '@client/ts/types/engagements'

@Component({
    components: {
        ControlSaveEditDialog,
    }
})
export default class ControlOverview extends Vue {
    get currentControl() : RawControl | null {
        return this.$store.state.controls.rawControl
    }

    get currentOrg() : RawOrganization | null {
        return this.$store.state.org.rawOrg
    }

    get currentEngagement() : RawEngagement | null {
        return this.$store.state.engagements.rawEngagement
    }

    onEditControl(r : RawControl) {
        this.$store.commit('controls/setRawControl', r)
    }
}

</script>
