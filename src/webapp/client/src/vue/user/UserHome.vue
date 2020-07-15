<template>
    <user-template
        page-name="Organizations"
        :relevant-user="currentUser"
    >
        <template v-slot:content>
            <loading-container
                :loading="!orgs"
            >

                <v-list-item class="px-0">
                    <v-list-item-content>
                        <v-list-item-title class="text-h4">
                            My Organizations
                        </v-list-item-title>
                    </v-list-item-content>

                    <v-spacer></v-spacer>

                    <v-list-item-action>
                        <v-dialog
                            v-model="showHideNewOrg"
                            max-width="40%"
                            persistent
                        >
                            <template v-slot:activator="{on}">
                                <v-btn
                                    color="primary"
                                    v-on="on"
                                >
                                    New
                                </v-btn>
                            </template>

                            <org-save-edit-dialog
                                @cancel-edit="showHideNewOrg = false"
                                @save-edit="saveNewOrg"
                            >
                            </org-save-edit-dialog>
                        </v-dialog>
                    </v-list-item-action>

                </v-list-item>
                <v-divider class="mb-4"></v-divider>

                <v-treeview
                >
                </v-treeview>
            </loading-container>
        </template>
    </user-template>
</template>

<script lang="ts">

import Vue from 'vue'
import Component from 'vue-class-component'
import { Watch } from 'vue-property-decorator'
import { RawUser } from '@client/ts/types/users'
import { RawOrganization } from '@client/ts/types/orgs'
import { GrchiveApi } from '@client/ts/main'
import UserTemplate from '@client/vue/UserTemplate.vue'
import LoadingContainer from '@client/vue/loading/LoadingContainer.vue'
import OrgSaveEditDialog from '@client/vue/types/orgs/OrgSaveEditDialog.vue'

@Component({
    components: {
        UserTemplate,
        OrgSaveEditDialog,
        LoadingContainer
    }
})
export default class UserHome extends Vue {
    orgs : RawOrganization[] | null = null
    showHideNewOrg : boolean = false

    saveNewOrg(org : RawOrganization) {
        this.showHideNewOrg = false
        if (!this.orgs) {
            return
        }
        this.orgs.unshift(org)
    }

    get currentUser() : RawUser | null {
        return this.$store.state.user.rawUser
    }

    get userOrganizationItems() : any[] {
        if (!this.orgs) {
            return []
        }
        // id: unique key
        // name: Text to display
        // children: children in tree.
        return this.orgs.map((ele : RawOrganization) => ({
            id: ele.Id,
            name: ele.Name,
        }))
    }

    refreshData() {
        GrchiveApi.user.getUserOrgs().then((resp : RawOrganization[] | null) => {
            this.orgs = resp
        })
    }

    mounted() {
        this.refreshData()
    }
}

</script>
