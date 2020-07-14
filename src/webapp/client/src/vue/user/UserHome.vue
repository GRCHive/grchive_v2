<template>
    <user-template>
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

    @Watch('currentUser')
    setTitle() {
        if (!this.currentUser) {
            return
        }
        document.title = `Welcome Back, ${this.currentUser.FullName}!`
    }

    mounted() {
        this.setTitle()
    }
}

</script>
