<template>
    <scoping-template
        page-name="Systems"
    >
        <template v-slot:content>
            <v-list-item class="px-0">
                <v-list-item-content>
                    <v-list-item-title class="text-h4">
                        Systems
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
                        <system-save-edit-dialog
                            :parent-org-id="currentOrg.Id"
                            :parent-engagement-id="currentEngagement.Id"
                            @cancel-edit="showHideNew=false"
                            @save-edit="onSaveSystem"
                        >
                        </system-save-edit-dialog>
                    </v-dialog>
                </v-list-item-action>
            </v-list-item>
            <v-divider class="mb-4"></v-divider>

            <loading-container
                :loading="!allSystems"
            >
                <template v-slot:default="{show}">
                    <full-height-base>
                        <system-grid
                            v-if="show"
                            :systems="allSystems"
                            style="height: 100%;"
                        >
                        </system-grid>
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
import { RawSystem } from '@client/ts/types/systems'
import { GrchiveApi, ErrorHandler } from '@client/ts/main'
import LoadingContainer from '@client/vue/loading/LoadingContainer.vue'
import FullHeightBase from '@client/vue/shared/FullHeightBase.vue'
import SystemSaveEditDialog from '@client/vue/types/systems/SystemSaveEditDialog.vue'
import SystemGrid from '@client/vue/types/systems/SystemGrid.vue'

@Component({
    components: {
        ScopingTemplate,
        RestrictRolePermissionButton,
        LoadingContainer,
        FullHeightBase,
        SystemSaveEditDialog,
        SystemGrid
    }
})
export default class ScopingSystems extends Vue {
    get currentOrg() : RawOrganization | null {
        return this.$store.state.org.rawOrg
    }

    get currentEngagement() : RawEngagement | null {
        return this.$store.state.engagements.rawEngagement
    }

    showHideNew: boolean = false
    allSystems : RawSystem[] | null = null

    get permissionsForCreate() : Permission[] {
        return [Permission.PSystemsCreate]
    }

    onSaveSystem(risk : RawSystem) {
        this.allSystems!.unshift(risk)
        this.showHideNew = false
    }

    @Watch('currentOrg')
    @Watch('currentEngagement')
    refreshData() {
        if (!this.currentOrg || !this.currentEngagement) {
            return
        }

        GrchiveApi.systems.listSystems(this.currentOrg.Id, this.currentEngagement.Id).then((resp : RawSystem[]) => {
            this.allSystems = resp
        }).catch((err : any) => {
            ErrorHandler.failurePageOnError(err)
        })
    }

    mounted() {
        this.refreshData()
    }
}

</script>
