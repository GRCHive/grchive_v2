<template>
    <generic-nav-bar
        :nav-links="navLinks"
    >
        <v-list-item
            dense
            :to="{name: 'orgHome', params: `${backParams}` }"
            link
            color="secondary"
            id="navBarHeader"
        >
            <v-list-item-icon>
                <v-icon>mdi-keyboard-return</v-icon>
            </v-list-item-icon>
            <v-list-item-content>
                <v-list-item-title>
                    Back to Organization
                </v-list-item-title>
            </v-list-item-content>
        </v-list-item>
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
    get backParams() : any {
        if (!this.currentOrg) {
            return {}
        }
        return {
            orgId: this.currentOrg.Id
        }
    }

    get currentOrg() : RawOrganization | null {
        return this.$store.state.org.rawOrg
    }

    get currentEngagement() : RawEngagement | null {
        return this.$store.state.engagements.rawEngagement
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
                params: {
                    orgId: this.currentOrg!.Id,
                    engId: this.currentEngagement!.Id,
                }
            },
            {
                title: 'Workflow',
                icon: 'mdi-sitemap',
                children: [
                    {
                        title: 'Scoping',
                        icon: 'mdi-crosshairs',
                        disabled: true,
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
