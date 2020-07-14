<template>
    <base-app-bar>
    </base-app-bar>
</template>

<script lang="ts">

import Vue from 'vue'
import Component from 'vue-class-component'
import BaseAppBar from '@client/vue/BaseAppBar.vue'

@Component({
    components: {
        BaseAppBar
    }
})
export default class UserAppBar extends Vue {
}

</script>
