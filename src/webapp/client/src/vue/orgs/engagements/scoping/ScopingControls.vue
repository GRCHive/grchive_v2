<template>
    <scoping-template
        page-name="Controls"
    >
        <template v-slot:content>
            <v-list-item class="px-0">
                <v-list-item-content>
                    <v-list-item-title class="text-h4">
                        Controls
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
                        :retain-focus="false"
                    >
                        <control-save-edit-dialog
                            :parent-org-id="currentOrg.Id"
                            :parent-engagement-id="currentEngagement.Id"
                            @cancel-edit="showHideNew = false"
                            @save-edit="onSaveControl"
                        >
                        </control-save-edit-dialog>
                    </v-dialog>
                </v-list-item-action>
            </v-list-item>
            <v-divider class="mb-4"></v-divider>

            <loading-container
                :loading="!allControls"
            >
                <template v-slot:default="{show}">
                    <full-height-base>
                        <control-grid
                            v-if="show"
                            :controls="allControls"
                            style="height: 100%;"
                        >
                        </control-grid>
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
import { RawControl } from '@client/ts/types/controls'
import { GrchiveApi } from '@client/ts/main'
import LoadingContainer from '@client/vue/loading/LoadingContainer.vue'
import FullHeightBase from '@client/vue/shared/FullHeightBase.vue'
import ControlSaveEditDialog from '@client/vue/types/controls/ControlSaveEditDialog.vue'
import ControlGrid from '@client/vue/types/controls/ControlGrid.vue'

@Component({
    components: {
        ScopingTemplate,
        RestrictRolePermissionButton,
        LoadingContainer,
        FullHeightBase,
        ControlSaveEditDialog,
        ControlGrid,
    }
})
export default class ScopingControls extends Vue {
    get currentOrg() : RawOrganization | null {
        return this.$store.state.org.rawOrg
    }

    get currentEngagement() : RawEngagement | null {
        return this.$store.state.engagements.rawEngagement
    }

    showHideNew: boolean = false
    allControls : RawControl[] | null = null

    get permissionsForCreate() : Permission[] {
        return [Permission.PControlsCreate, Permission.POrgUsersList, Permission.POrgUsersView]
    }

    onSaveControl(control : RawControl) {
        this.showHideNew = false
        this.allControls!.unshift(control)
    }

    @Watch('currentOrg')
    @Watch('currentEngagement')
    refreshData() {
        if (!this.currentOrg || !this.currentEngagement) {
            return
        }

        GrchiveApi.controls.listControls(this.currentOrg!.Id, this.currentEngagement!.Id).then((resp : RawControl[]) => {
            this.allControls = resp
        })
    }

    mounted() {
        this.refreshData()
    }
}

</script>
