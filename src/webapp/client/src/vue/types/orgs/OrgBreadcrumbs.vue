<template>
    <v-breadcrumbs
        class="pa-0"
        :items="parentOrgItems"
    >
    </v-breadcrumbs>
</template>

<script lang="ts">

import Vue from 'vue'
import Component from 'vue-class-component'
import { Watch, Prop } from 'vue-property-decorator'
import { RawOrganization } from '@client/ts/types/orgs'
import { GrchiveApi } from '@client/ts/main'

@Component
export default class OrgBreadcrumbs extends Vue {
    @Prop({ required: true} )
    org! : RawOrganization

    parentOrgs: RawOrganization[] | null = null

    get parentOrgItems() : any[] {
        if (!this.parentOrgs) {
            return []
        }

        return this.parentOrgs.reverse().map((ele : RawOrganization) => {
            return {
                text: ele.Name,
                to: {
                    name: 'orgHome',
                    params: { orgId: `${ele.Id}` },
                }
            }
        })
    }

    @Watch('org')
    refreshData() {
        GrchiveApi.orgs.getParentOrgs(this.org.Id).then((resp : RawOrganization[] | null) => {
            this.parentOrgs = resp
        })
    }

    mounted() {
        this.refreshData()
    }
}

</script>
