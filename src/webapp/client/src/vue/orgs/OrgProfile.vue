<template>
    <org-template
        page-name="Profile"
    >
        <template v-slot:content>
            <org-breadcrumbs
                :org="currentOrg"
                class="mt-3"
            >
            </org-breadcrumbs>

            <v-list-item class="px-0">
                <v-list-item-content>
                    <v-list-item-title class="text-h4">
                        {{ currentOrg.Name }} Profile
                    </v-list-item-title>
                </v-list-item-content>
            </v-list-item>
            <v-divider></v-divider>

            <v-tabs>
                <v-tab :to="overviewTo">Overview</v-tab>
                <v-tab :to="treeTo">Org Tree</v-tab>
            </v-tabs>
            <router-view></router-view>
        </template>
    </org-template>
</template>

<script lang="ts">

import Vue from 'vue'
import Component from 'vue-class-component'
import { RawOrganization } from '@client/ts/types/orgs'
import OrgTemplate from '@client/vue/orgs/OrgTemplate.vue'
import OrgBreadcrumbs from '@client/vue/types/orgs/OrgBreadcrumbs.vue'

@Component({
    components: {
        OrgBreadcrumbs,
        OrgTemplate,
    }
})
export default class OrgProfile extends Vue {
    get currentOrg() : RawOrganization | null {
        return this.$store.state.org.rawOrg
    }

    get overviewTo() : any {
        return {
            path: 'overview',
            params: { orgId: this.currentOrg!.Id },
        }
    }

    get treeTo() : any {
        return {
            path: 'tree',
            params: { orgId: this.currentOrg!.Id },
        }
    }
}

</script>
