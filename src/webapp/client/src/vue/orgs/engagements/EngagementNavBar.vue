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
export default class EngagementNavBar extends Vue {
    get backParams() : NavLink {
        let ret : NavLink = {
            title: 'Organization',
            icon: '',
            path: 'orgHome'
        }
        if (!!this.currentOrg) {
            ret.params = {
                orgId: this.currentOrg.Id
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
                title: 'Overview',
                icon: 'mdi-home',
                path: 'orgSingleEngagement'
                params: this.baseParams,
            },
            {
                title: 'Workflow',
                icon: 'mdi-sitemap',
                children: [
                    {
                        title: 'Scoping',
                        icon: 'mdi-crosshairs',
                        path: 'scopingDashboard'
                        params: this.baseParams,
                    },
                    {
                        title: 'PBC',
                        icon: 'mdi-share-variant',
                        disabled: true,
                    },
                ]
            },
            {
                title: 'Automation',
                icon: 'mdi-robot',
                disabled: true,
            },
            {
                title: 'CA/CM',
                icon: 'mdi-refresh',
                disabled: true,
            },
            {
                title: 'Roles',
                icon: 'mdi-security',
                disabled: true,
            },
            {
                title: 'Users',
                icon: 'mdi-account-group',
                disabled: true,
            },
            {
                title: 'Audit Trail',
                icon: 'mdi-shoe-print',
                disabled: true,
            },
            {
                title: 'Settings',
                icon: 'mdi-cog',
                disabled: true,
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
