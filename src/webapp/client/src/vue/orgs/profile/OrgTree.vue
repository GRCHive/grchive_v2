<template>
    <org-tree-viewer
        :orgs="suborgs"
        :current-org-id="currentOrg.Id"
        :root-org="currentOrg"
        allow-add-suborgs
    >
    </org-tree-viewer>
</template>

<script lang="ts">

import Vue from 'vue'
import Component from 'vue-class-component'
import OrgTreeViewer from '@client/vue/types/orgs/OrgTreeViewer.vue'
import { Watch } from 'vue-property-decorator'
import { GrchiveApi, ErrorHandler } from '@client/ts/main'
import { RawOrganization } from '@client/ts/types/orgs'

@Component({
    components: {
        OrgTreeViewer,
    }
})
export default class OrgTree extends Vue {
    suborgs : RawOrganization[] | null = null

    get currentOrg() : RawOrganization | null {
        return this.$store.state.org.rawOrg
    }

    @Watch('currentOrg')
    refreshData() {
        if (!this.currentOrg) {
            return
        }

        GrchiveApi.orgs.getSuborgs(this.currentOrg!.Id).then((resp : RawOrganization[]) => {
            this.suborgs = resp
        }).catch((err : any) => {
            ErrorHandler.failurePopupOnError(err, {
                context: 'Failed to get organization tree.'
            })
        })
    }

    mounted() {
        this.refreshData()
    }
}

</script>
