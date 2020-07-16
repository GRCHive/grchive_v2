<template>
    <base-template>
        <template v-slot:appbar>
        </template>

        <template v-slot:navbar>
        </template>

        <template v-slot:content>
            <loading-container
                :loading="!currentOrg"
            >
                <div class="mx-4">
                    <slot name="content">
                    </slot>
                </div>
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

@Component({
    components: {
        BaseTemplate,
        LoadingContainer,
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

    refreshOrg() {
        const orgId : number = Number(this.$route.params.orgId)
        this.$store.dispatch('org/initializeOrgStore', orgId)
    }

    mounted() {
        this.refreshOrg()
    }
}

</script>
