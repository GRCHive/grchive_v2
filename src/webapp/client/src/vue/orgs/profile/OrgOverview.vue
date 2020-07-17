<template>
    <org-save-edit-dialog
        :value="currentOrg"
        edit-mode
        @save-edit="saveOrgProfileEdit"
    >
    </org-save-edit-dialog>
</template>

<script lang="ts">

import Vue from 'vue'
import Component from 'vue-class-component'
import OrgSaveEditDialog from '@client/vue/types/orgs/OrgSaveEditDialog.vue'
import { RawOrganization } from '@client/ts/types/orgs'

@Component({
    components: {
        OrgSaveEditDialog,
    }
})
export default class OrgOverview extends Vue {
    get currentOrg() : RawOrganization | null {
        return this.$store.state.org.rawOrg
    }

    saveOrgProfileEdit(newOrgState : RawOrganization) {
        this.$store.commit('org/setRawOrg', newOrgState)
    }
}

</script>
