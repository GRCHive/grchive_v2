<template>
    <v-breadcrumbs
        class="pa-0"
        :items="parentAccItems"
    >
    </v-breadcrumbs>
</template>

<script lang="ts">

import Vue from 'vue'
import Component from 'vue-class-component'
import { Watch, Prop } from 'vue-property-decorator'
import { RawGLAccount } from '@client/ts/types/gl'
import { GrchiveApi, ErrorHandler } from '@client/ts/main'
import { RawOrganization } from '@client/ts/types/orgs'
import { RawEngagement } from '@client/ts/types/engagements'

@Component
export default class GlAccountBreadcrumbs extends Vue {
    @Prop({ required: true} )
    acc! : RawGLAccount

    parentAccs: RawGLAccount[] | null = null

    get currentOrg() : RawOrganization | null {
        return this.$store.state.org.rawOrg
    }

    get currentEngagement() : RawEngagement | null {
        return this.$store.state.engagements.rawEngagement
    }

    get parentAccItems() : any[] {
        if (!this.parentAccs) {
            return []
        }

        return this.parentAccs.slice().reverse().map((ele : RawGLAccount) => {
            return {
                text: `${ele.AccountId}: ${ele.Name}`,
                to: {
                    name: 'glAccHome',
                    params: {
                        ...this.$route.params,
                        accId: `${ele.Id}`
                    },
                }
            }
        })
    }

    get parentAccountId() : number | null {
        return this.acc.ParentAccountId
    }

    @Watch('currentOrg')
    @Watch('currentEngagement')
    @Watch('parentAccountId')
    refreshData() {
        if (!this.currentOrg || !this.currentEngagement) {
            this.parentAccs = null
            return
        }

        GrchiveApi.gl.listParentAccounts(this.currentOrg.Id, this.currentEngagement.Id, this.acc.Id).then((resp : RawGLAccount[]) => {
            this.parentAccs = resp
        }).catch((err : any) => {
            ErrorHandler.failurePopupOnError(err, {
                context: 'Failed to get parent accounts.'
            })
        })
    }

    mounted() {
        this.refreshData()
    }
}

</script>
