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
                    :disabled="selectedRisks.length == 0"
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
                            Select Risk
                        </v-card-title>

                        <v-form v-model="formValid">
                            <single-risk-finder
                                v-model="linkRisk"
                                :org-id="currentOrg.Id"
                                :engagement-id="currentEngagement.Id"
                                :rules="[rules.required]"
                            >
                            </single-risk-finder>
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
            :loading="!relatedRisks"
        >
            <template v-slot:default="{show}">
                <full-height-base>
                    <risk-grid
                        v-if="show"
                        v-model="selectedRisks"
                        :risks="rawRisks"
                        :rels="relatedRisks"
                        :selectable="canSelect"
                        style="height: 100%;"
                    >
                    </risk-grid>
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
import { RawRisk } from '@client/ts/types/risks'
import { GrchiveApi, ErrorHandler } from '@client/ts/main'
import { RelationshipWrapper } from '@client/ts/types/relationships'
import { RawOrganization } from '@client/ts/types/orgs'
import { RawEngagement } from '@client/ts/types/engagements'
import { Permission } from '@client/ts/types/roles'
import SingleRiskFinder from '@client/vue/types/risks/SingleRiskFinder.vue'
import RiskGrid from '@client/vue/types/risks/RiskGrid.vue'
import LoadingContainer from '@client/vue/loading/LoadingContainer.vue'
import FullHeightBase from '@client/vue/shared/FullHeightBase.vue'
import RestrictRolePermissionButton from '@client/vue/loading/RestrictRolePermissionButton.vue'
import * as rules from '@client/ts/frontend/formRules'

@Component({
    components: {
        RiskGrid,
        LoadingContainer,
        FullHeightBase,
        RestrictRolePermissionButton,
        SingleRiskFinder
    }
})
export default class RiskRelationshipsDisplay extends Vue {
    readonly rules : any = rules
    relatedRisks : RelationshipWrapper<RawRisk>[] | null = null
    showHideNew: boolean = false
    linkInProgress: boolean = false
    linkRisk : RawRisk | null = null
    formValid: boolean = false

    canSelect: boolean = false
    selectedRisks: RawRisk[] = []

    get currentOrg() : RawOrganization | null {
        return this.$store.state.org.rawOrg
    }

    get currentEngagement() : RawEngagement | null {
        return this.$store.state.engagements.rawEngagement
    }

    get currentResourceId() : ResourceId {
        return createResourceIdFromRouteParams(this.$route.params)
    }

    get rawRisks() : RawRisk[] {
        if (!this.relatedRisks) {
            return []
        }
        return this.relatedRisks.map((ele : RelationshipWrapper<RawRisk>) => ele.Data)
    }

    get permissionsForCreate() : Permission[] {
        return [Permission.PRelRisksControlsCreate]
    }

    get permissionsForDelete() : Permission[] {
        return [Permission.PRelRisksControlsDelete]
    }

    @Watch('currentResourceId')
    refreshData() {
        GrchiveApi.relationships.listRelationships<RawRisk>(this.currentResourceId, ResourceType.RTRisk).then((resp : RelationshipWrapper<RawRisk>[]) => {
            this.relatedRisks = resp
        }).catch((err : any) => {
            ErrorHandler.failurePageOnError(err)
        })
    }

    onCancel() {
        this.showHideNew = false
        this.linkRisk = null
    }

    onLink() {
        if (!this.linkRisk) {
            return
        }

        this.linkInProgress = true
        GrchiveApi.relationships.createRelationship<RawRisk>(this.currentResourceId, ResourceType.RTRisk, {
            orgId: this.currentOrg!.Id,
            engagementId: this.currentEngagement!.Id,
            riskId: this.linkRisk.Id,
        }).then((resp : RelationshipWrapper<RawRisk>) => {
            if (!this.relatedRisks) {
                return
            }
            this.relatedRisks.unshift(resp)
            this.onCancel()
        }).catch((err : any) => {
            ErrorHandler.failurePopupOnError(err, {
                context: 'Failed to create a relationship to the selected risk.'
            })
        }).finally(() => {
            this.linkInProgress = false
        })
    }

    mounted() {
        this.refreshData()
    }

    deleteSelected() {
        if (this.selectedRisks.length == 0) {
            return
        }
        this.linkInProgress = true
        for (let c of this.selectedRisks) {
            GrchiveApi.relationships.deleteRelationship(this.currentResourceId, ResourceType.RTRisk, c.Id).then(() => {
                if (!this.relatedRisks) {
                    return
                }
                let idx = this.relatedRisks.findIndex((ele : RelationshipWrapper<RawRisk>) => ele.Data.Id == c.Id)
                if (idx == -1) {
                    return
                }
                this.relatedRisks.splice(idx, 1)
            }).catch((err : any) => {
                ErrorHandler.failurePopupOnError(err, {
                    context: 'Failed to delete a risk-control relationship.'
                })
            }).finally(() => {
                this.linkInProgress = false
            })
        }
    }
}

</script>
