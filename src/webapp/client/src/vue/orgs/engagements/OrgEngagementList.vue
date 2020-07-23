<template>
    <org-template
        page-name="Engagements"
    >
        <template v-slot:content>
            <v-list-item class="px-0">
                <v-list-item-content>
                    <v-list-item-title class="text-h4">
                        {{ currentOrg.Name }} Engagements
                    </v-list-item-title>
                </v-list-item-content>

                <v-spacer></v-spacer>

                <v-list-item-action>
                    <restrict-role-permission-button
                        color="primary"
                        :permissions="permissionsForCreate"
                        :org-id="currentOrg.Id"
                        @click="showHideNew = true"
                    >
                        New
                    </restrict-role-permission-button>

                    <v-dialog
                        v-model="showHideNew"
                        persistent
                        max-width="40%"
                        :retain-focus="false"
                    >
                        <engagement-save-edit-dialog
                            :parent-org-id="currentOrg.Id"
                            @cancel-edit="showHideNew = false"
                            @save-edit="saveNewEngagement"
                        >
                        </engagement-save-edit-dialog>
                    </v-dialog>
                </v-list-item-action>
            </v-list-item>
            <v-divider class="mb-4"></v-divider>

            <loading-container
                :loading="!allEngagements"
            >
                <template v-slot:default="{show}">
                    <full-height-base>
                        <engagement-grid
                            v-if="show"
                            :engagements="allEngagements"
                            style="height: 100%;"
                        >
                        </engagement-grid>
                    </full-height-base>
                </template>
            </loading-container>
        </template>
    </org-template>

</template>

<script lang="ts">

import Vue from 'vue'
import Component from 'vue-class-component'
import { Watch } from 'vue-property-decorator'
import { RawOrganization } from '@client/ts/types/orgs'
import { GrchiveApi } from '@client/ts/main'
import { RawEngagement } from '@client/ts/types/engagements'
import { Permission } from '@client/ts/types/roles'
import OrgTemplate from '@client/vue/orgs/OrgTemplate.vue'
import LoadingContainer from '@client/vue/loading/LoadingContainer.vue'
import RestrictRolePermissionButton from '@client/vue/loading/RestrictRolePermissionButton.vue'
import EngagementSaveEditDialog from '@client/vue/types/engagements/EngagementSaveEditDialog.vue'
import EngagementGrid from '@client/vue/types/engagements/EngagementGrid.vue'
import FullHeightBase from '@client/vue/shared/FullHeightBase.vue'

@Component({
    components: {
        OrgTemplate,
        LoadingContainer,
        RestrictRolePermissionButton,
        EngagementSaveEditDialog,
        EngagementGrid,
        FullHeightBase
    }
})
export default class OrgEngagementList extends Vue {
    get currentOrg() : RawOrganization | null {
        return this.$store.state.org.rawOrg
    }

    allEngagements : RawEngagement[] | null = null
    showHideNew : boolean = false

    get permissionsForCreate() : Permission[] {
        return [Permission.POrgEngagementCreate, Permission.POrgEngagementList]
    }

    saveNewEngagement(e : RawEngagement) {
        this.showHideNew = false
        if (!this.allEngagements) {
            return
        }
        this.allEngagements.unshift(e)
    }

    @Watch('currentOrg')
    refreshEngagements() {
        if (!this.currentOrg) {
            this.allEngagements = null
            return
        }
        GrchiveApi.engagements.listOrgEngagements(this.currentOrg.Id).then((resp : RawEngagement[]) => {
            this.allEngagements = resp
        })
    }

    mounted() {
        this.refreshEngagements()
    }
}

</script>
