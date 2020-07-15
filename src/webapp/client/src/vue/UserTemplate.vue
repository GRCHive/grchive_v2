<template>
    <base-template>
        <template v-slot:appbar>
            <user-app-bar></user-app-bar>
        </template>

        <template v-slot:navbar>
            <user-nav-bar></user-nav-bar>
        </template>

        <template v-slot:content>
            <div class="mx-4" v-if="!!relevantUser">
                <slot name="content">
                </slot>
            </div>

            <v-row justify="center" v-else>
                <v-progress-circular size=64 indeterminate>
                </v-progress-circular>
            </v-row>
        </template>
    </base-template>
</template>

<script lang="ts">

import Vue from 'vue'
import Component from 'vue-class-component'
import { Watch, Prop } from 'vue-property-decorator'
import BaseTemplate from '@client/vue/BaseTemplate.vue'
import UserAppBar from '@client/vue/user/UserAppBar.vue'
import UserNavBar from '@client/vue/user/UserNavBar.vue'
import { RawUser } from '@client/ts/types/users'

@Component({
    components: {
        BaseTemplate,
        UserAppBar,
        UserNavBar
    }
})
export default class UserTemplate extends Vue {
    @Prop({ default: '' })
    pageName! : string

    @Prop({ required: true })
    relevantUser! : RawUser | null

    @Watch('relevantUser')
    refreshTitle() {
        if (!this.relevantUser) {
            return
        }

        document.title = `${this.relevantUser.FullName} - ${this.pageName}`
    }

    mounted() {
        this.refreshTitle()
    }
}

</script>
