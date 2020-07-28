<template>
    <database-save-edit-dialog
        :value="currentDatabase"
        :parent-org-id="currentOrg.Id"
        :parent-engagement-id="currentEngagement.Id"
        edit-mode
        @save-edit="onEditDatabase"
    >
    </database-save-edit-dialog>
</template>

<script lang="ts">

import Vue from 'vue'
import Component from 'vue-class-component'
import DatabaseSaveEditDialog from '@client/vue/types/databases/DatabaseSaveEditDialog.vue'
import { RawDatabase } from '@client/ts/types/databases'
import { RawOrganization } from '@client/ts/types/orgs'
import { RawEngagement } from '@client/ts/types/engagements'

@Component({
    components: {
        DatabaseSaveEditDialog,
    }
})
export default class DatabaseOverview extends Vue {
    get currentDatabase() : RawDatabase | null {
        return this.$store.state.databases.rawDatabase
    }

    get currentOrg() : RawOrganization | null {
        return this.$store.state.org.rawOrg
    }

    get currentEngagement() : RawEngagement | null {
        return this.$store.state.engagements.rawEngagement
    }

    onEditDatabase(r : RawDatabase) {
        this.$store.commit('databases/setRawDatabase', r)
    }
}

</script>
