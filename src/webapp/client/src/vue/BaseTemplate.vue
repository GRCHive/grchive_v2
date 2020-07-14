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
    </v-app>
</template>

<script lang="ts">

import Vue from 'vue'
import Component from 'vue-class-component'

@Component
export default class BaseTemplate extends Vue {
    get isLoading() : boolean {
        return !this.$store.getters.isFinishedLoading
    }

    mounted() {
        this.$store.dispatch('initialize')
    }
}

</script>
