<template>
    <v-app>
        <div v-if="!isLoading">
            <slot name="appbar">
            </slot>

            <slot name="navbar">
            </slot>

            <v-main>
                <slot name="content">
                </slot>
            </v-main>
        </div>

        <v-main class="max-height" v-else>
            <v-container class="max-height">
                <v-row justify="center" align="center" class="max-height">
                    <v-progress-circular size=64 indeterminate>
                    </v-progress-circular>
                </v-row>
            </v-container>

        </v-main>

        <error-popup-manager></error-popup-manager>
    </v-app>
</template>

<script lang="ts">

import Vue from 'vue'
import Component from 'vue-class-component'
import ErrorPopupManager from '@client/vue/ErrorPopupManager.vue'

@Component({
    components: {
        ErrorPopupManager,
    }
})
export default class BaseTemplate extends Vue {
    get isLoading() : boolean {
        return !this.$store.getters.isFinishedLoading
    }

    mounted() {
        this.$store.dispatch('initialize')
    }
}

</script>
