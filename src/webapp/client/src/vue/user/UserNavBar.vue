<template>
    <generic-nav-bar
        :nav-links="navLinks"
    >
    </generic-nav-bar>
</template>

<script lang="ts">

import Vue from 'vue'
import Component from 'vue-class-component'
import { NavLink } from '@client/ts/frontend/navLink'
import GenericNavBar from '@client/vue/navbar/GenericNavBar.vue'

@Component({
    components: {
        GenericNavBar
    }
})
export default class UserNavBar extends Vue {

    get navLinks() : NavLink[] {
        return [
            {
                title: 'Organizations',
                icon: 'mdi-account-group',
                path: 'userHome',
            },
            {
                title: 'Notifications',
                icon: 'mdi-bell',
                disabled: true,
            },
            {
                title: 'Profile',
                icon: 'mdi-account-circle',
                path: 'userProfile',
            }
        ]
    }
}

</script>
