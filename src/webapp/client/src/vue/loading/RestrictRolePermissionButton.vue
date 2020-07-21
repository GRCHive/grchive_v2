<template>
    <div style="position: relative;">
        <v-btn
            :icon="icon"
            :color="color"
            @mousedown.stop
            @mouseup.stop
            @click.stop="onClick"
            :disabled="!hasPermissions"
        >
            <slot v-bind:show="!!hasPermissions"></slot>
        </v-btn>

        <v-overlay
            absolute
            :value="!hasPermissions"
        >
            <div style="z-index: 10;">
                <v-row justify="center" v-if="isLoading">
                    <v-progress-circular size=16 indeterminate>
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
            </div>
        </v-overlay>
    </div>
</template>

<script lang="ts">

import RestrictRolePermissionBase from '@client/vue/loading/RestrictRolePermissionBase.vue'
import Component, { mixins } from 'vue-class-component'
import { VBtn } from 'vuetify/lib'

@Component
export default class RestrictRolePermissionButton extends mixins(VBtn, RestrictRolePermissionBase) {
    onClick() {
        this.$emit('click')
    }
}

</script>
