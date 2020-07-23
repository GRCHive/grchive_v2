<template>
    <scoping-template
        page-name="Risks"
    >
        <template v-slot:content>
            <v-list-item class="px-0">
                <v-list-item-content>
                    <v-list-item-title class="text-h4">
                        Risks
                    </v-list-item-title>
                </v-list-item-content>

                <v-spacer></v-spacer>

                <v-list-item-action>
                    <restrict-role-permission-button
                        color="primary"
                        :permissions="permissionsForCreate"
                        :org-id="currentOrg.Id"
                        :engagement-id="currentEngagement.Id"
                        @click="showHideNew = true"
                    >
                        New
                    </restrict-role-permission-button>

                    <v-dialog
                        v-model="showHideNew"
                        persistent
                        max-width="40%"
                    >
                        <risk-save-edit-dialog
                            :parent-org-id="currentOrg.Id"
                            :parent-engagement-id="currentEngagement.Id"
                            @cancel-edit="showHideNew=false"
                            @save-edit="onSaveRisk"
                        >
                        </risk-save-edit-dialog>
                    </v-dialog>
                </v-list-item-action>
            </v-list-item>
            <v-divider class="mb-4"></v-divider>

            <loading-container
                :loading="!allRisks"
            >
                <template v-slot:default="{show}">
                    <full-height-base>
                        <risk-grid
                            v-if="show"
                            :risks="allRisks"
                            style="height: 100%;"
                        >
                        </risk-grid>
                    </full-height-base>
                </template>
            </loading-container>
        </template>
    </scoping-template>
</template>

<script lang="ts">

import Vue from 'vue'
import Component from 'vue-class-component'
import { Watch } from 'vue-property-decorator'
import ScopingTemplate from '@client/vue/orgs/engagements/scoping/ScopingTemplate.vue'
import RestrictRolePermissionButton from '@client/vue/loading/RestrictRolePermissionButton.vue'
import { Permission } from '@client/ts/types/roles'
import { RawOrganization } from '@client/ts/types/orgs'
import { RawEngagement } from '@client/ts/types/engagements'
import { RawRisk } from '@client/ts/types/risks'
import { GrchiveApi } from '@client/ts/main'
import RiskSaveEditDialog from '@client/vue/types/risks/RiskSaveEditDialog.vue'
import LoadingContainer from '@client/vue/loading/LoadingContainer.vue'
import RiskGrid from '@client/vue/types/risks/RiskGrid.vue'
import FullHeightBase from '@client/vue/shared/FullHeightBase.vue'

@Component({
    components: {
        ScopingTemplate,
        RestrictRolePermissionButton,
        RiskSaveEditDialog,
        LoadingContainer,
        RiskGrid,
        FullHeightBase
    }
})
export default class ScopingRisks extends Vue {
    get currentOrg() : RawOrganization | null {
        return this.$store.state.org.rawOrg
    }

    get currentEngagement() : RawEngagement | null {
        return this.$store.state.engagements.rawEngagement
    }

    showHideNew: boolean = false
    allRisks : RawRisk[] | null = null

    get permissionsForCreate() : Permission[] {
        return [Permission.PRisksCreate]
    }

    onSaveRisk(risk : RawRisk) {
        this.allRisks!.unshift(risk)
        this.showHideNew = false
    }

    @Watch('currentOrg')
    @Watch('currentEngagement')
    refreshData() {
        if (!this.currentOrg || !this.currentEngagement) {
            return
        }

        GrchiveApi.risks.listRisks(this.currentOrg.Id, this.currentEngagement.Id).then((resp : RawRisk[]) => {
            this.allRisks = resp
        })
    }

    mounted() {
        this.refreshData()
    }
}

</script>
