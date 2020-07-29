<template>
    <div>
        <v-list-item class="px-0">
            <v-list-item-action v-if="canSelect" class="mr-4">
                <restrict-role-permission-button
                    color="error"
                    :permissions="permissionsForDelete"
                    :org-id="currentOrg.Id"
                    :engagement-id="currentEngagement.Id"
                    @click="deleteSelected"
                    :loading="linkInProgress"
                    :disabled="selectedControls.length == 0"
                >
                    Delete
                </restrict-role-permission-button>

            </v-list-item-action>

            <v-list-item-action>
                <v-btn
                    color="secondary"
                    v-if="!canSelect"
                    @click="canSelect = true"
                >
                    Manage
                </v-btn>

                <v-btn
                    color="secondary"
                    v-if="canSelect"
                    @click="canSelect = false"
                >
                    Cancel
                </v-btn>
            </v-list-item-action>

            <v-spacer></v-spacer>

            <v-list-item-action>
                <restrict-role-permission-button
                    color="primary"
                    :permissions="permissionsForCreate"
                    :org-id="currentOrg.Id"
                    :engagement-id="currentEngagement.Id"
                    @click="showHideNew = true"
                >
                    Link
                </restrict-role-permission-button>

                <v-dialog
                    v-model="showHideNew"
                    persistent
                    max-width="40%"
                >
                    <v-card>
                        <v-card-title>
                            Select Control
                        </v-card-title>

                        <v-form v-model="formValid">
                            <single-control-finder
                                v-model="linkControl"
                                :org-id="currentOrg.Id"
                                :engagement-id="currentEngagement.Id"
                                :rules="[rules.required]"
                            >
                            </single-control-finder>
                        </v-form>

                        <v-card-actions>
                            <v-btn
                                color="error"
                                @click="onCancel"
                                :loading="linkInProgress"
                            >
                                Cancel
                            </v-btn>

                            <v-spacer></v-spacer>

                            <v-btn
                                color="primary"
                                @click="onLink"
                                :loading="linkInProgress"
                                :disabled="!formValid"
                            >
                                Link
                            </v-btn>
                        </v-card-actions>
                    </v-card>
                </v-dialog>
            </v-list-item-action>
        </v-list-item>

        <loading-container
            :loading="!relatedControls"
        >
            <template v-slot:default="{show}">
                <full-height-base>
                    <control-grid
                        v-if="show"
                        v-model="selectedControls"
                        :controls="rawControls"
                        :rels="relatedControls"
                        :selectable="canSelect"
                        style="height: 100%;"
                    >
                    </control-grid>
                </full-height-base>
            </template>
        </loading-container>
    </div>
</template>

<script lang="ts">

import Vue from 'vue'
import Component from 'vue-class-component'
import { Watch } from 'vue-property-decorator'
import { createResourceIdFromRouteParams } from '@client/ts/frontend/routeUtility'
import { ResourceId, ResourceType } from '@client/ts/types/resource'
import { RawControl } from '@client/ts/types/controls'
import { GrchiveApi, ErrorHandler } from '@client/ts/main'
import { RelationshipWrapper } from '@client/ts/types/relationships'
import { RawOrganization } from '@client/ts/types/orgs'
import { RawEngagement } from '@client/ts/types/engagements'
import { Permission } from '@client/ts/types/roles'
import SingleControlFinder from '@client/vue/types/controls/SingleControlFinder.vue'
import ControlGrid from '@client/vue/types/controls/ControlGrid.vue'
import LoadingContainer from '@client/vue/loading/LoadingContainer.vue'
import FullHeightBase from '@client/vue/shared/FullHeightBase.vue'
import RestrictRolePermissionButton from '@client/vue/loading/RestrictRolePermissionButton.vue'
import * as rules from '@client/ts/frontend/formRules'

@Component({
    components: {
        ControlGrid,
        LoadingContainer,
        FullHeightBase,
        RestrictRolePermissionButton,
        SingleControlFinder
    }
})
export default class ControlRelationshipsDisplay extends Vue {
    readonly rules : any = rules
    relatedControls : RelationshipWrapper<RawControl>[] | null = null
    showHideNew: boolean = false
    linkInProgress: boolean = false
    linkControl : RawControl | null = null
    formValid: boolean = false

    canSelect: boolean = false
    selectedControls: RawControl[] = []

    get currentOrg() : RawOrganization | null {
        return this.$store.state.org.rawOrg
    }

    get currentEngagement() : RawEngagement | null {
        return this.$store.state.engagements.rawEngagement
    }

    get currentResourceId() : ResourceId {
        return createResourceIdFromRouteParams(this.$route.params)
    }

    get rawControls() : RawControl[] {
        if (!this.relatedControls) {
            return []
        }
        return this.relatedControls.map((ele : RelationshipWrapper<RawControl>) => ele.Data)
    }

    get permissionsForCreate() : Permission[] {
        return [Permission.PRelRisksControlsCreate]
    }

    get permissionsForDelete() : Permission[] {
        return [Permission.PRelRisksControlsDelete]
    }

    @Watch('currentResourceId')
    refreshData() {
        GrchiveApi.relationships.listRelationships<RawControl>(this.currentResourceId, ResourceType.RTControl).then((resp : RelationshipWrapper<RawControl>[]) => {
            this.relatedControls = resp
        }).catch((err : any) => {
            ErrorHandler.failurePageOnError(err)
        })
    }

    onCancel() {
        this.showHideNew = false
        this.linkControl = null
    }

    onLink() {
        if (!this.linkControl) {
            return
        }

        this.linkInProgress = true
        GrchiveApi.relationships.createRelationship<RawControl>(this.currentResourceId, ResourceType.RTControl, {
            orgId: this.currentOrg!.Id,
            engagementId: this.currentEngagement!.Id,
            controlId: this.linkControl.Id,
        }).then((resp : RelationshipWrapper<RawControl>) => {
            if (!this.relatedControls) {
                return
            }
            this.relatedControls.unshift(resp)
            this.onCancel()
        }).catch((err : any) => {
            ErrorHandler.failurePopupOnError(err, {
                context: 'Failed to create a relationship to the selected control.'
            })
        }).finally(() => {
            this.linkInProgress = false
        })
    }

    mounted() {
        this.refreshData()
    }

    deleteSelected() {
        if (this.selectedControls.length == 0) {
            return
        }
        this.linkInProgress = true
        for (let c of this.selectedControls) {
            GrchiveApi.relationships.deleteRelationship(this.currentResourceId, ResourceType.RTControl, c.Id).then(() => {
                if (!this.relatedControls) {
                    return
                }
                let idx = this.relatedControls.findIndex((ele : RelationshipWrapper<RawControl>) => ele.Data.Id == c.Id)
                if (idx == -1) {
                    return
                }
                this.relatedControls.splice(idx, 1)
            }).catch((err : any) => {
                ErrorHandler.failurePopupOnError(err, {
                    context: 'Failed to delete a control-risk relationship.'
                })
            }).finally(() => {
                this.linkInProgress = false
            })
        }
    }
}

</script>
