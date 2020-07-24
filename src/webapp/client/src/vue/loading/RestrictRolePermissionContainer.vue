<template>
    <div style="position: relative;">
        <slot v-bind:show="!!hasPermissions"></slot>
        <v-overlay
            absolute
            :value="!hasPermissions"
        >
            <v-row justify="center" v-if="isLoading">
                <v-progress-circular size=64 indeterminate>
                </v-progress-circular>
            </v-row>

            <v-tooltip max-width="400px" bottom v-else>
                <template v-slot:activator="{ on, attrs }">
                    <v-icon
                        color="error"
                        v-bind="attrs"
                        v-on="on"
                    >
                        mdi-lock
                    </v-icon>
                </template>
                <span>{{ tooltipStr }}</span>
            </v-tooltip>
        </v-overlay>
    </div>
</template>

<script lang="ts">

import RestrictRolePermissionBase from '@client/vue/loading/RestrictRolePermissionBase.vue'
import Component from 'vue-class-component'

@Component
export default class RestrictRolePermissionContainer extends RestrictRolePermissionBase {
}

</script>

