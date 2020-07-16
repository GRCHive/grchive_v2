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
export default class OrgNavBar extends Vue {

    get navLinks() : NavLink[] {
        return [
            {
                title: 'Profile',
                icon: 'mdi-home',
                path: 'orgHome',
            },
            {
                title: 'Engagements',
                icon: 'mdi-briefcase',
                disabled: true,
            },
            {
                title: 'Roles',
                icon: 'mdi-security',
                disabled: true,
            },
            {
                title: 'Users',
                icon: 'mdi-account-group',
                disabled: true,
            },
            {
                title: 'Audit Trail',
                icon: 'mdi-shoe-print',
                disabled: true,
            },
            {
                title: 'Settings',
                icon: 'mdi-cog',
                disabled: true,
            },
        ]
    }
}

</script>

