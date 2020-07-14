<template>
    <base-template>
        <template v-slot:appbar>
            <user-app-bar></user-app-bar>
        </template>

        <template v-slot:navbar>
            <user-nav-bar></user-nav-bar>
        </template>

        <template v-slot:content>
            <slot name="content">
            </slot>
        </template>
    </base-template>
</template>

<script lang="ts">

import Vue from 'vue'
import Component from 'vue-class-component'
import BaseTemplate from '@client/vue/BaseTemplate.vue'
import UserAppBar from '@client/vue/user/UserAppBar.vue'
import UserNavBar from '@client/vue/user/UserNavBar.vue'

@Component({
    components: {
        BaseTemplate,
        UserAppBar,
        UserNavBar
    }
})
export default class UserTemplate extends Vue {
}

</script>
