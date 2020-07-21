<template>
    <generic-nav-bar
        :nav-links="navLinks"
        :back-link="backParams"
    >
    </generic-nav-bar>
</template>

<script lang="ts">

import Vue from 'vue'
import Component from 'vue-class-component'
import { NavLink } from '@client/ts/frontend/navLink'
import { RawOrganization } from '@client/ts/types/orgs'
import { RawEngagement } from '@client/ts/types/engagements'
import GenericNavBar from '@client/vue/navbar/GenericNavBar.vue'

@Component({
    components: {
        GenericNavBar
    }
})
export default class ScopingNavBar extends Vue {
    get backParams() : NavLink {
        let ret : NavLink = {
            title: 'Engagement',
            icon: '',
            path: 'orgSingleEngagement'
        }
        if (!!this.currentOrg && !!this.currentEngagement) {
            ret.params = {
                orgId: this.currentOrg.Id,
                engId: this.currentEngagement.Id
            }
        }
        return ret
    }

    get currentOrg() : RawOrganization | null {
        return this.$store.state.org.rawOrg
    }

    get currentEngagement() : RawEngagement | null {
        return this.$store.state.engagements.rawEngagement
    }

    get baseParams() : any {
        return {
            orgId: this.currentOrg!.Id,
            engId: this.currentEngagement!.Id,
        }
    }

    get navLinks() : NavLink[] {
        if (!this.currentOrg || !this.currentEngagement) {
            return []
        }

        return [
            {
                title: 'Dashboard',
                icon: 'mdi-view-dashboard',
                disabled: true,
            },
            {
                title: 'Relationships',
                icon: 'mdi-transit-connection-variant',
                children: [
                    {
                        title: 'Processes',
                        icon: 'mdi-graph-outline',
                        disabled: true,
                    },
                    {
                        title: 'Networks',
                        icon: 'mdi-wan',
                        disabled: true,
                    },
                ],
            },
            {
                title: 'RCM',
                icon: 'mdi-grid',
                children: [
                    {
                        title: 'Risks',
                        icon: 'mdi-fire',
                        path: 'scopingRisks',
                        params: this.baseParams,
                    },
                    {
                        title: 'Controls',
                        icon: 'mdi-shield-lock-outline',
                        disabled: true,
                    },
                ],
            },
            {
                title: 'Vendors',
                icon: 'mdi-store',
                disabled: true,
            },
            {
                title: 'General Ledger',
                icon: 'mdi-bank-outline',
                disabled: true,
            },
            {
                title: 'IT',
                icon: 'mdi-web',
                children: [
                    {
                        title: 'Systems',
                        icon: 'mdi-application',
                        disabled: true,
                    },
                    {
                        title: 'Databases',
                        icon: 'mdi-database',
                        disabled: true,
                    },
                    {
                        title: 'Servers',
                        icon: 'mdi-server-network',
                        disabled: true,
                    },
                ],
            },
        ]
    }
}

</script>

<style scoped>

#navBarHeader {
    background-color: #DCDCDC;
}

</style>
