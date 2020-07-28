<template>
    <system-save-edit-dialog
        :value="currentSystem"
        :parent-org-id="currentOrg.Id"
        :parent-engagement-id="currentEngagement.Id"
        edit-mode
        @save-edit="onEditSystem"
    >
    </system-save-edit-dialog>
</template>

<script lang="ts">

import Vue from 'vue'
import Component from 'vue-class-component'
import SystemSaveEditDialog from '@client/vue/types/systems/SystemSaveEditDialog.vue'
import { RawSystem } from '@client/ts/types/systems'
import { RawOrganization } from '@client/ts/types/orgs'
import { RawEngagement } from '@client/ts/types/engagements'

@Component({
    components: {
        SystemSaveEditDialog,
    }
})
export default class SystemOverview extends Vue {
    get currentSystem() : RawSystem | null {
        return this.$store.state.systems.rawSystem
    }

    get currentOrg() : RawOrganization | null {
        return this.$store.state.org.rawOrg
    }

    get currentEngagement() : RawEngagement | null {
        return this.$store.state.engagements.rawEngagement
    }

    onEditSystem(r : RawSystem) {
        this.$store.commit('systems/setRawSystem', r)
    }
}

</script>
