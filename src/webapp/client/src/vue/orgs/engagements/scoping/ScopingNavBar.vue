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
import { Permission } from '@client/ts/types/roles'
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
                path: 'scopingDashboard',
                params: this.baseParams,
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
                        permissions: [Permission.PRisksList],
                    },
                    {
                        title: 'Controls',
                        icon: 'mdi-shield-lock-outline',
                        path: 'scopingControls',
                        params: this.baseParams,
                        permissions: [Permission.PControlsList],
                    },
                ],
            },
            {
                title: 'Vendors',
                icon: 'mdi-store',
                path: 'scopingVendors',
                params: this.baseParams,
                permissions: [Permission.PVendorsList],
            },
            {
                title: 'General Ledger',
                icon: 'mdi-bank-outline',
                path: 'glHome',
                params: this.baseParams,
                permissions: [Permission.PGLList],
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
                        path: 'scopingDatabases',
                        params: this.baseParams,
                        permissions: [Permission.PDatabasesList],
                    },
                    {
                        title: 'Inventory',
                        icon: 'mdi-warehouse',
                        path: 'scopingInventory',
                        params: this.baseParams,
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
