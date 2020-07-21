<template>
    <base-template>
        <template v-slot:appbar>
            <org-app-bar></org-app-bar>
        </template>

        <template v-slot:navbar>
            <org-nav-bar></org-nav-bar>
        </template>

        <template v-slot:content>
            <loading-container
                :loading="!currentOrg"
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
import BaseTemplate from '@client/vue/BaseTemplate.vue'
import LoadingContainer from '@client/vue/loading/LoadingContainer.vue'
import OrgAppBar from '@client/vue/orgs/OrgAppBar.vue'
import OrgNavBar from '@client/vue/orgs/OrgNavBar.vue'

@Component({
    components: {
        BaseTemplate,
        LoadingContainer,
        OrgAppBar,
        OrgNavBar,
    }
})
export default class OrgTemplate extends Vue {
    @Prop({ default: '' })
    pageName! : string

    @Watch('currentOrg')
    refreshTitle() {
        if (!this.currentOrg) {
            return
        }
        document.title = `${this.currentOrg.Name} - ${this.pageName}`
    }

    get currentOrg() : RawOrganization | null {
        return this.$store.state.org.rawOrg
    }

    @Watch('$route')
    refreshOrg() {
        const orgId : number = Number(this.$route.params.orgId)
        this.$store.dispatch('initializeCurrentResource', {
            orgId,
        })
        this.refreshTitle()
    }

    mounted() {
        this.refreshOrg()
    }
}

</script>
