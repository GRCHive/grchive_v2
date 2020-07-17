<template>
    <org-template
        page-name="Profile"
    >
        <template v-slot:content>
            <v-list-item class="px-0">
                <v-list-item-content>
                    <v-list-item-title class="text-h4">
                        Organization Profile
                    </v-list-item-title>
                </v-list-item-content>
            </v-list-item>
            <v-divider class="mb-4"></v-divider>

            <v-tabs>
                <v-tab>Overview</v-tab>
                <v-tab-item>
                    <loading-container
                        :loading="!currentOrg"
                    >
                        <org-save-edit-dialog
                            :value="currentOrg"
                            edit-mode
                            @save-edit="saveOrgProfileEdit"
                        >
                        </org-save-edit-dialog>
                    </loading-container>
                </v-tab-item>

                <v-tab>Org Tree</v-tab>
                <v-tab-item>
                    <loading-container
                        :loading="!suborgs || !currentOrg"
                    >
                        <org-tree-viewer
                            :orgs="suborgs"
                            :current-org-id="currentOrgId"
                            allow-add-suborgs
                        >
                        </org-tree-viewer>
                    </loading-container>
                </v-tab-item>
            </v-tabs>
        </template>
    </org-template>
</template>

<script lang="ts">

import Vue from 'vue'
import Component from 'vue-class-component'
import { Watch } from 'vue-property-decorator'
import OrgTemplate from '@client/vue/orgs/OrgTemplate.vue'
import OrgSaveEditDialog from '@client/vue/types/orgs/OrgSaveEditDialog.vue'
import OrgTreeViewer from '@client/vue/types/orgs/OrgTreeViewer.vue'
import LoadingContainer from '@client/vue/loading/LoadingContainer.vue'
import { RawOrganization } from '@client/ts/types/orgs'
import { GrchiveApi } from '@client/ts/main'

@Component({
    components: {
        OrgTemplate,
        OrgSaveEditDialog,
        OrgTreeViewer,
        LoadingContainer,
    }
})
export default class OrgProfile extends Vue {
    suborgs : RawOrganization[] | null = null

    get currentOrg() : RawOrganization | null {
        return this.$store.state.org.rawOrg
    }

    get currentOrgId() : number {
        if (!this.currentOrg) {
            return -1
        }
        return this.currentOrg.Id
    }

    @Watch('currentOrg')
    refreshData() {
        if (!this.currentOrg) {
            return
        }

        GrchiveApi.orgs.getSuborgs(this.currentOrg!.Id).then((resp : RawOrganization[] | null) => {
            this.suborgs = resp
        })
    }

    mounted() {
        this.refreshData()
    }

    saveOrgProfileEdit(newOrgState : RawOrganization) {
        this.$store.commit('org/setRawOrg', newOrgState)
    }
}

</script>
