<template>
    <base-template>
        <template v-slot:appbar>
            <engagement-app-bar>
            </engagement-app-bar>
        </template>

        <template v-slot:navbar>
            <engagement-nav-bar>
            </engagement-nav-bar>
        </template>

        <template v-slot:content>
            <loading-container
                :loading="loading"
            >
                <template v-slot:default="{show}">
                    <div class="mx-4" v-if="show">
                        <slot name="content">
                        </slot>
                    </div>
                </template>
            </loading-container>
        </template>
    </base-template>

</template>

<script lang="ts">

import Vue from 'vue'
import Component from 'vue-class-component'
import { Watch, Prop } from 'vue-property-decorator'
import { RawOrganization } from '@client/ts/types/orgs'
import { RawEngagement } from '@client/ts/types/engagements'
import BaseTemplate from '@client/vue/BaseTemplate.vue'
import LoadingContainer from '@client/vue/loading/LoadingContainer.vue'
import EngagementAppBar from '@client/vue/orgs/engagements/EngagementAppBar.vue'
import EngagementNavBar from '@client/vue/orgs/engagements/EngagementNavBar.vue'

@Component({
    components: {
        BaseTemplate,
        LoadingContainer,
        EngagementAppBar,
        EngagementNavBar
    }
})
export default class EngagementTemplate extends Vue {
    @Prop({ required: true })
    pageName! : string

    @Prop({ type: Boolean, default: false })
    disableDefaultInit!: boolean

    get loading() : boolean {
        return !this.currentOrg ||!this.currentEngagement
    }

    get currentOrg() : RawOrganization | null {
        return this.$store.state.org.rawOrg
    }

    get currentEngagement() : RawEngagement | null {
        return this.$store.state.engagements.rawEngagement
    }

    @Watch('currentOrg')
    @Watch('currentEngagement')
    @Watch('pageName')
    refreshTitle() {
        if (!this.currentOrg || !this.currentEngagement) {
            return
        }
        document.title = `${this.currentOrg.Name} - ${this.currentEngagement.Name} - ${this.pageName}`
    }

    @Watch('$route')
    refreshEngagement() {
        if (this.disableDefaultInit) {
            return
        }

        const orgId : number = Number(this.$route.params.orgId)
        const engId : number = Number(this.$route.params.engId)
        this.$store.dispatch('initializeCurrentResource', {
            orgId,
            engagementId: engId,
        })
        this.refreshTitle()
    }

    mounted() {
        this.refreshEngagement()
    }

}

</script>
