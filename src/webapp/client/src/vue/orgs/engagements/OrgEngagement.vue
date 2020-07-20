<template>
    <engagement-template
        page-name="Overview"
    >
        <template v-slot:content>
            <v-list-item class="px-0">
                <v-list-item-content>
                    <v-list-item-title class="text-h4">
                        {{ currentEngagement.Name }}
                    </v-list-item-title>
                </v-list-item-content>

                <v-spacer></v-spacer>
            </v-list-item>
            <v-divider class="mb-4"></v-divider>

            <engagement-save-edit-dialog
                :value="currentEngagement"
                :parent-org-id="currentOrg.Id"
                @save-edit="updateEngagement"
                edit-mode
            >
            </engagement-save-edit-dialog>
        </template>
    </engagement-template>
</template>

<script lang="ts">

import Vue from 'vue'
import Component from 'vue-class-component'
import EngagementTemplate from '@client/vue/orgs/engagements/EngagementTemplate.vue'
import { RawEngagement } from '@client/ts/types/engagements'
import { RawOrganization } from '@client/ts/types/orgs'
import EngagementSaveEditDialog from '@client/vue/types/engagements/EngagementSaveEditDialog.vue'

@Component({
    components: {
        EngagementTemplate,
        EngagementSaveEditDialog,
    }
})
export default class OrgEngagement extends Vue {
    get currentEngagement() : RawEngagement | null {
        return this.$store.state.engagements.rawEngagement
    }

    get currentOrg() : RawOrganization | null {
        return this.$store.state.org.rawOrg
    }

    updateEngagement(e : RawEngagement) {
        this.$store.commit('engagements/setRawEngagement', e)
    }
}

</script>
