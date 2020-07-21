<template>
    <base-template>
        <template v-slot:appbar>
            <engagement-app-bar>
            </engagement-app-bar>
        </template>

        <template v-slot:navbar>
            <scoping-nav-bar>
            </scoping-nav-bar>
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
import Component, {mixins} from 'vue-class-component'
import EngagementTemplate from '@client/vue//orgs/engagements/EngagementTemplate.vue'
import LoadingContainer from '@client/vue/loading/LoadingContainer.vue'
import EngagementAppBar from '@client/vue/orgs/engagements/EngagementAppBar.vue'
import ScopingNavBar from '@client/vue/orgs/engagements/scoping/ScopingNavBar.vue'
import BaseTemplate from '@client/vue/BaseTemplate.vue'

@Component({
    components: {
        BaseTemplate,
        LoadingContainer,
        EngagementAppBar,
        ScopingNavBar
    }
})
export default class ScopingTemplate extends mixins(EngagementTemplate) {
}

</script>
