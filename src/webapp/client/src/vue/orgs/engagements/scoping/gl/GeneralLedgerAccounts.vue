<template>
    <div>
        <v-list-item class="px-0">
            <v-list-item-content>
                <v-list-item-title class="text-h6 font-weight-bold">
                    Chart of Accounts
                </v-list-item-title>
            </v-list-item-content>

            <v-spacer></v-spacer>

            <restrict-role-permission-button
                :permissions="newAccountPermissions"
                :org-id="currentOrg.Id"
                :engagement-id="currentEngagement.Id"
                color="primary"
                @click="showHideNew = true"
            >
                New
            </restrict-role-permission-button>

            <v-dialog
                v-model="showHideNew"
                persistent
                max-width="40%"
            >
                <general-ledger-account-save-edit-dialog
                    :parent-engagement-id="currentEngagement.Id"
                    :parent-org-id="currentOrg.Id"
                    @cancel-edit="showHideNew = false"
                    @save-edit="onNewAccount"
                    :key="dialogKey"
                >
                </general-ledger-account-save-edit-dialog>
            </v-dialog>

        </v-list-item>
        <loading-container
            :loading="!accounts"
        >
            <template v-slot:default="{show}">
                <gl-account-tree-viewer
                    v-if="show"
                    :accounts="accounts"
                >
                </gl-account-tree-viewer>
            </template>
        </loading-container>
    </div>
</template>

<script lang="ts">

import Vue from 'vue'
import Component from 'vue-class-component'
import { Watch } from 'vue-property-decorator'
import { RawOrganization } from '@client/ts/types/orgs'
import { RawEngagement } from '@client/ts/types/engagements'
import { RawGLAccount } from '@client/ts/types/gl'
import { GrchiveApi } from '@client/ts/main'
import { Permission } from '@client/ts/types/roles'
import LoadingContainer from '@client/vue/loading/LoadingContainer.vue'
import RestrictRolePermissionButton from '@client/vue/loading/RestrictRolePermissionButton.vue'
import GeneralLedgerAccountSaveEditDialog from '@client/vue/types/gl/GeneralLedgerAccountSaveEditDialog.vue'
import GlAccountTreeViewer from '@client/vue/types/gl/GlAccountTreeViewer.vue'

@Component({
    components: {
        LoadingContainer,
        RestrictRolePermissionButton,
        GeneralLedgerAccountSaveEditDialog,
        GlAccountTreeViewer
    }
})
export default class GeneralLedgerAccounts extends Vue {
    get currentOrg() : RawOrganization | null {
        return this.$store.state.org.rawOrg
    }

    get currentEngagement() : RawEngagement | null {
        return this.$store.state.engagements.rawEngagement
    }

    get newAccountPermissions() : Permission[] {
        return [Permission.PGLCreate]
    }

    accounts : RawGLAccount[] | null = null
    showHideNew: boolean = false
    dialogKey : number = 1

    onNewAccount(acc : RawGLAccount) {
        this.showHideNew = false
        if (!this.accounts) {
            return
        }
        this.accounts.unshift(acc)

        // Need to force the dialog to regenerate to repull the list of eligible parent accounts.
        this.dialogKey += 1
    }

    @Watch('currentOrg')
    @Watch('currentEngagement')
    refreshData() {
        if (!this.currentOrg || !this.currentEngagement) {
            return
        }

        GrchiveApi.gl.listAccounts(this.currentOrg!.Id, this.currentEngagement!.Id).then((resp : RawGLAccount[] | null) => {
            this.accounts = resp
        })
    }

    mounted() {
        this.refreshData()
    }
}

</script>
