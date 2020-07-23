<template>
    <scoping-template
        page-name="General Ledger"
    >
        <template v-slot:content>
            <v-list-item class="px-0">
                <v-list-item-content>
                    <v-list-item-title class="text-h4">
                        General Ledger
                    </v-list-item-title>
                </v-list-item-content>

                <v-spacer></v-spacer>
            </v-list-item>
            <v-divider class="mb-4"></v-divider>

            <v-tabs>
                <v-tab :to="accountsTo">Accounts</v-tab>
                <restrict-role-permission-tab
                    :permissions="commentPermissions"
                    :to="commentsTo"
                    :org-id="currentOrg.Id"
                    :engagement-id="currentEngagement.Id"
                >
                    Comments
                </restrict-role-permission-tab>
            </v-tabs>
            <router-view></router-view>
        </template>
    </scoping-template>
</template>

<script lang="ts">

import Vue from 'vue'
import Component from 'vue-class-component'
import { Watch } from 'vue-property-decorator'
import ScopingTemplate from '@client/vue/orgs/engagements/scoping/ScopingTemplate.vue'
import { Permission } from '@client/ts/types/roles'
import { RawOrganization } from '@client/ts/types/orgs'
import { RawEngagement } from '@client/ts/types/engagements'
import { RawGLAccount } from '@client/ts/types/gl'
import LoadingContainer from '@client/vue/loading/LoadingContainer.vue'
import FullHeightBase from '@client/vue/shared/FullHeightBase.vue'
import RestrictRolePermissionTab from '@client/vue/loading/RestrictRolePermissionTab.vue'

@Component({
    components: {
        ScopingTemplate,
        RestrictRolePermissionTab,
        LoadingContainer,
        FullHeightBase
    }
})
export default class ScopingGeneralLedger extends Vue {
    get currentOrg() : RawOrganization | null {
        return this.$store.state.org.rawOrg
    }

    get currentEngagement() : RawEngagement | null {
        return this.$store.state.engagements.rawEngagement
    }

    get commentPermissions() : Permission[] {
        return [Permission.PGLList, Permission.PCommentsList, Permission.POrgUsersView]
    }

    get accountsTo() : any {
        return {
            name: 'glAccounts',
            params: this.$route.params,
        }
    }

    get commentsTo() : any {
        return {
            name: 'glComments',
            params: this.$route.params,
        }
    }

}

</script>
