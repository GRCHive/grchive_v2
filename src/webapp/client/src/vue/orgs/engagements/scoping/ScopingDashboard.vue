<template>
    <scoping-template
        page-name="Dashboard"
    >
        <template v-slot:content>
            <v-list-item class="px-0">
                <v-list-item-content>
                    <v-list-item-title class="text-h4">
                        Dashboard
                    </v-list-item-title>
                </v-list-item-content>
            </v-list-item>
            <v-divider class="mb-4"></v-divider>

            <loading-container
                :loading="!stats"
            >
                <template v-slot:default="{show}">
                    <div v-if="show">
                        <v-row justify="center">
                            <v-col cols="1">
                                <v-card>
                                    <div class="stat-container">
                                        <div class="d-flex justify-center icon-container">
                                            <v-icon x-large>
                                                mdi-fire
                                            </v-icon>
                                        </div>

                                        <div class="d-flex justify-center text-overline">
                                            <span class="font-weight-bold">
                                                {{ stats.NumRisks }}
                                            </span>
                                            Risks
                                        </div>
                                    </div>
                                </v-card>
                            </v-col>

                            <v-col cols="1">
                                <v-card>
                                    <div class="stat-container">
                                        <div class="d-flex justify-center icon-container">
                                            <v-icon x-large>
                                                mdi-shield-lock-outline
                                            </v-icon>
                                        </div>

                                        <div class="d-flex justify-center text-overline">
                                            <span class="font-weight-bold">
                                                {{ stats.NumControls }}
                                            </span>
                                            Controls
                                        </div>
                                    </div>
                                </v-card>
                            </v-col>

                            <v-col cols="1">
                                <v-card>
                                    <div class="stat-container">
                                        <div class="d-flex justify-center icon-container">
                                            <v-icon x-large>
                                                mdi-bank-outline
                                            </v-icon>
                                        </div>

                                        <div class="d-flex justify-center text-overline">
                                            <span class="font-weight-bold">
                                                {{ stats.NumGLAccounts }}
                                            </span>
                                            Accounts
                                        </div>

                                        <div class="d-flex justify-center text-overline">
                                            <span class="font-weight-bold">
                                                {{ stats.NumFinanciallyRelevantGLAccounts }}
                                            </span>
                                            Relevant
                                        </div>
                                    </div>
                                </v-card>
                            </v-col>

                            <v-col cols="1">
                                <v-card>
                                    <div class="stat-container">
                                        <div class="d-flex justify-center icon-container">
                                            <v-icon x-large>
                                                mdi-store
                                            </v-icon>
                                        </div>

                                        <div class="d-flex justify-center text-overline">
                                            <span class="font-weight-bold">
                                                {{ stats.NumVendors }}
                                            </span>
                                            Vendors
                                        </div>

                                        <div class="d-flex justify-center text-overline">
                                            <span class="font-weight-bold">
                                                {{ stats.NumVendorProducts }}
                                            </span>
                                            Products
                                        </div>
                                    </div>
                                </v-card>
                            </v-col>

                        </v-row>
                    </div>
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
import { RawOrganization } from '@client/ts/types/orgs'
import { RawEngagement, RawEngagementScopingStats } from '@client/ts/types/engagements'
import { GrchiveApi, ErrorHandler } from '@client/ts/main'
import LoadingContainer from '@client/vue/loading/LoadingContainer.vue'

@Component({
    components: {
        ScopingTemplate,
        LoadingContainer,
    }
})
export default class ScopingDashboard extends Vue {
    stats : RawEngagementScopingStats | null = null

    get currentOrg() : RawOrganization | null {
        return this.$store.state.org.rawOrg
    }

    get currentEngagement() : RawEngagement | null {
        return this.$store.state.engagements.rawEngagement
    }

    @Watch('currentOrg')
    @Watch('currentEngagement')
    refreshData() {
        if (!this.currentOrg || !this.currentEngagement) {
            return
        }

        GrchiveApi.engagements.getScopingStats(this.currentOrg.Id, this.currentEngagement.Id).then((resp : RawEngagementScopingStats) => {
            this.stats = resp
        }).catch((err : any) => {
            // Don't direct to failure page here so that the rest of the scoping page is still accessible.
            // If this page directs to failure then the user won't be able to access any other aspects of scoping by default
            // if the error is *only* with the dashboard stats endpoint.
            ErrorHandler.failurePopupOnError(err, {
                context: 'Failed to get scoping stats.'
            })
        })
    }

    mounted() {
        this.refreshData()
    }
}

</script>

<style scoped>

.stat-container {
    padding-top: 16px;
    padding-bottom: 16px;
}

.icon-container {
    margin-bottom: 16px;
}

>>>.text-overline {
    line-height: 1rem;
}

</style>
