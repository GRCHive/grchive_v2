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
                                <stat-display
                                    icon="mdi-fire"
                                    :stats="riskStats"
                                >
                                </stat-display>
                            </v-col>

                            <v-col cols="1">
                                <stat-display
                                    icon="mdi-shield-lock-outline"
                                    :stats="controlStats"
                                >
                                </stat-display>
                            </v-col>

                            <v-col cols="1">
                                <stat-display
                                    icon="mdi-bank-outline"
                                    :stats="glStats"
                                >
                                </stat-display>
                            </v-col>

                            <v-col cols="1">
                                <stat-display
                                    icon="mdi-store"
                                    :stats="vendorStats"
                                >
                                </stat-display>
                            </v-col>

                            <v-col cols="1">
                                <stat-display
                                    icon="mdi-application"
                                    :stats="systemStats"
                                >
                                </stat-display>
                            </v-col>

                            <v-col cols="1">
                                <stat-display
                                    icon="mdi-database"
                                    :stats="databaseStats"
                                >
                                </stat-display>
                            </v-col>
                        </v-row>

                        <v-row justify="center">
                            <v-col cols="1">
                                <stat-display
                                    icon="mdi-server"
                                    :stats="serverStats"
                                >
                                </stat-display>
                            </v-col>

                            <v-col cols="1">
                                <stat-display
                                    icon="mdi-monitor"
                                    :stats="desktopStats"
                                >
                                </stat-display>
                            </v-col>

                            <v-col cols="1">
                                <stat-display
                                    icon="mdi-laptop"
                                    :stats="laptopStats"
                                >
                                </stat-display>
                            </v-col>

                            <v-col cols="1">
                                <stat-display
                                    icon="mdi-cellphone"
                                    :stats="mobileStats"
                                >
                                </stat-display>
                            </v-col>

                            <v-col cols="1">
                                <stat-display
                                    icon="mdi-chip"
                                    :stats="embeddedStats"
                                >
                                </stat-display>
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
import { StatInfo } from '@client/ts/stats'
import { InventoryType } from '@client/ts/types/inventory'
import LoadingContainer from '@client/vue/loading/LoadingContainer.vue'
import StatDisplay from '@client/vue/shared/StatDisplay.vue'

@Component({
    components: {
        ScopingTemplate,
        LoadingContainer,
        StatDisplay,
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

    get riskStats() : StatInfo[] {
        return [{
            value: this.stats!.NumRisks,
            text: 'Risks'
        }]
    }

    get controlStats() : StatInfo[] {
        return [{
            value: this.stats!.NumControls,
            text: 'Risks'
        }]
    }

    get glStats() : StatInfo[] {
        return [{
            value: this.stats!.NumGLAccounts,
            text: 'Accounts'
        }, {
            value: this.stats!.NumFinanciallyRelevantGLAccounts,
            text: 'Relevant'
        }]
    }

    get vendorStats() : StatInfo[] {
        return [{
            value: this.stats!.NumVendors,
            text: 'Vendors'
        }, {
            value: this.stats!.NumVendorProducts,
            text: 'Products'
        }]
    }

    get systemStats() : StatInfo[] {
        return [{
            value: this.stats!.NumSystems,
            text: 'Systems'
        }]
    }

    get databaseStats() : StatInfo[] {
        return [{
            value: this.stats!.NumDatabases,
            text: 'Databases'
        }]
    }

    get serverStats() : StatInfo[] {
        return [{
            value: this.stats!.NumInventory[InventoryType.ITServer],
            text: 'Servers'
        }]
    }

    get desktopStats() : StatInfo[] {
        return [{
            value: this.stats!.NumInventory[InventoryType.ITDesktop],
            text: 'Desktops'
        }]
    }

    get laptopStats() : StatInfo[] {
        return [{
            value: this.stats!.NumInventory[InventoryType.ITLaptop],
            text: 'Laptops'
        }]
    }

    get mobileStats() : StatInfo[] {
        return [{
            value: this.stats!.NumInventory[InventoryType.ITMobile],
            text: 'Mobile'
        }]
    }

    get embeddedStats() : StatInfo[] {
        return [{
            value: this.stats!.NumInventory[InventoryType.ITEmbedded],
            text: 'Embedded'
        }]
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
