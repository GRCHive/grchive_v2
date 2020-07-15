<template>
    <user-template
        page-name="Organizations"
        :relevant-user="currentUser"
    >
        <template v-slot:content>
            <h1>Hello World</h1>
        </template>
    </user-template>
</template>

<script lang="ts">

import Vue from 'vue'
import Component from 'vue-class-component'
import { Watch } from 'vue-property-decorator'
import UserTemplate from '@client/vue/UserTemplate.vue'
import { RawUser } from '@client/ts/types/users'

@Component({
    components: {
        UserTemplate
    }
})
export default class UserHome extends Vue {
    get currentUser() : RawUser | null {
        return this.$store.state.user.rawUser
    }
}

</script>
