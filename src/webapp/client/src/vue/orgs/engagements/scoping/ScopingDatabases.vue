<template>
    <scoping-template
        page-name="Databases"
    >
        <template v-slot:content>
            <v-list-item class="px-0">
                <v-list-item-content>
                    <v-list-item-title class="text-h4">
                        Databases
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
                        <database-save-edit-dialog
                            :parent-org-id="currentOrg.Id"
                            :parent-engagement-id="currentEngagement.Id"
                            @cancel-edit="showHideNew=false"
                            @save-edit="onSaveDatabase"
                        >
                        </database-save-edit-dialog>
                    </v-dialog>
                </v-list-item-action>
            </v-list-item>
            <v-divider class="mb-4"></v-divider>

            <loading-container
                :loading="!allDatabases"
            >
                <template v-slot:default="{show}">
                    <full-height-base>
                        <database-grid
                            v-if="show"
                            :databases="allDatabases"
                            style="height: 100%;"
                        >
                        </database-grid>
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
import { RawDatabase } from '@client/ts/types/databases'
import { GrchiveApi, ErrorHandler } from '@client/ts/main'
import LoadingContainer from '@client/vue/loading/LoadingContainer.vue'
import FullHeightBase from '@client/vue/shared/FullHeightBase.vue'
import DatabaseSaveEditDialog from '@client/vue/types/databases/DatabaseSaveEditDialog.vue'
import DatabaseGrid from '@client/vue/types/databases/DatabaseGrid.vue'

@Component({
    components: {
        ScopingTemplate,
        RestrictRolePermissionButton,
        LoadingContainer,
        FullHeightBase,
        DatabaseSaveEditDialog,
        DatabaseGrid
    }
})
export default class ScopingDatabases extends Vue {
    get currentOrg() : RawOrganization | null {
        return this.$store.state.org.rawOrg
    }

    get currentEngagement() : RawEngagement | null {
        return this.$store.state.engagements.rawEngagement
    }

    showHideNew: boolean = false
    allDatabases : RawDatabase[] | null = null

    get permissionsForCreate() : Permission[] {
        return [Permission.PDatabasesCreate]
    }

    onSaveDatabase(risk : RawDatabase) {
        this.allDatabases!.unshift(risk)
        this.showHideNew = false
    }

    @Watch('currentOrg')
    @Watch('currentEngagement')
    refreshData() {
        if (!this.currentOrg || !this.currentEngagement) {
            return
        }

        GrchiveApi.databases.listDatabases(this.currentOrg.Id, this.currentEngagement.Id).then((resp : RawDatabase[]) => {
            this.allDatabases = resp
        }).catch((err : any) => {
            ErrorHandler.failurePageOnError(err)
        })
    }

    mounted() {
        this.refreshData()
    }
}

</script>
