<template>
    <loading-container
        :loading="!subaccounts"
    >
        <template v-slot:default="{show}">
            <gl-account-tree-viewer
                v-if="show"
                :accounts="subaccounts"
                :root="currentGeneralLedgerAccount"
            >
            </gl-account-tree-viewer>
        </template>
    </loading-container>
</template>

<script lang="ts">

import Vue from 'vue'
import Component from 'vue-class-component'
import { Watch } from 'vue-property-decorator'
import { RawOrganization } from '@client/ts/types/orgs'
import { RawEngagement } from '@client/ts/types/engagements'
import { RawGLAccount } from '@client/ts/types/gl'
import { GrchiveApi, ErrorHandler } from '@client/ts/main'
import LoadingContainer from '@client/vue/loading/LoadingContainer.vue'
import GlAccountTreeViewer from '@client/vue/types/gl/GlAccountTreeViewer.vue'

@Component({
    components: {
        GlAccountTreeViewer,
        LoadingContainer
    }
})
export default class GeneralLedgerAccountSubaccounts extends Vue {
    subaccounts : RawGLAccount[] | null = null

    get currentOrg() : RawOrganization | null {
        return this.$store.state.org.rawOrg
    }

    get currentEngagement() : RawEngagement | null {
        return this.$store.state.engagements.rawEngagement
    }

    get currentGeneralLedgerAccount() : RawGLAccount | null {
        return this.$store.state.gl.rawGLAccount
    }

    @Watch('currentOrg')
    @Watch('currentEngagement')
    @Watch('currentGeneralLedgerAccount')
    refreshData() {
        if (!this.currentOrg || !this.currentEngagement || !this.currentGeneralLedgerAccount) {
            this.subaccounts = null
            return
        }

        GrchiveApi.gl.listSubaccounts(this.currentOrg.Id, this.currentEngagement.Id, this.currentGeneralLedgerAccount.Id).then((resp : RawGLAccount[]) => {
            this.subaccounts = resp
        }).catch((err : any) => {
            ErrorHandler.failurePopupOnError(err, {
                context: 'Failed to get subaccounts.'
            })
        })
    }

    mounted() {
        this.refreshData()
    }
}

</script>
