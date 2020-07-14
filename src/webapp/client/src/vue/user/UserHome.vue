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
import UserTemplate from '@client/vue/UserTemplate.vue'

@Component({
    components: {
        UserTemplate
    }
})
export default class UserHome extends Vue {
}

</script>
