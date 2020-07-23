<template>
    <base-app-bar>
        <template v-slot:left-items>
            <v-dialog
                v-model="showHideSelector"
                persistent
                max-width="40%"
            >
                <template v-slot:activator="{on}">
                    <v-text-field
                        dense
                        :value="engagementName"
                        hide-details
                        solo
                        flat
                        v-on="on"
                    >
                        <template v-slot:append>
                            <v-icon small>
                                mdi-chevron-down
                            </v-icon>
                        </template>
                    </v-text-field>
                </template>

                <v-card>
                    <v-card-title>
                        Select Engagement
                    </v-card-title>

                    <engagement-grid
                        :engagements="allEngagements"
                        style="height: 600px;"
                        @change-engagement="showHideSelector = false"
                    >
                    </engagement-grid>

                    <v-card-actions>
                        <v-btn
                            color="error"
                            @click="showHideSelector = false"
                        >
                            Cancel
                        </v-btn>
                    </v-card-actions>
                </v-card>
            </v-dialog>
        </template>
    </base-app-bar>
</template>

<script lang="ts">

import Vue from 'vue'
import Component from 'vue-class-component'
import { Watch } from 'vue-property-decorator'
import BaseAppBar from '@client/vue/BaseAppBar.vue'
import { RawOrganization } from '@client/ts/types/orgs'
import { RawEngagement } from '@client/ts/types/engagements'
import { GrchiveApi } from '@client/ts/main'
import EngagementGrid from '@client/vue/types/engagements/EngagementGrid.vue'

@Component({
    components: {
        BaseAppBar,
        EngagementGrid
    }
})
export default class EngagementAppBar extends Vue {
    get currentOrg() : RawOrganization | null {
        return this.$store.state.org.rawOrg
    }

    get currentEngagement() : RawEngagement | null {
        return this.$store.state.engagements.rawEngagement
    }

    get engagementName() : string {
        if (!this.currentEngagement) {
            return ''
        }
        return this.currentEngagement.Name
    }

    allEngagements : RawEngagement[] | null = null
    showHideSelector: boolean = false

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

<style scoped>

>>>.v-toolbar__content > .v-input {
    flex-grow: 0 !important;
}

>>>.v-toolbar__content > .v-input .v-input__slot {
    cursor: pointer;
}

>>>.v-toolbar__content > .v-input .v-input__slot input {
    cursor: pointer;
    caret-color: transparent;
}

>>>.v-toolbar__content > .v-input .v-input__slot i {
    cursor: pointer;
}

</style>
