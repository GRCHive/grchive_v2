<template>
    <v-btn
        :icon="icon"
        :color="color"
        @mousedown.stop
        @mouseup.stop
        @click.stop="onClick"
        :disabled="!hasPermissions"
    >
        <slot v-bind:show="!!hasPermissions"></slot>
        <v-overlay
            absolute
            :value="!hasPermissions"
        >
            <v-row justify="center" v-if="isLoading">
                <v-progress-circular size=16 indeterminate>
                </v-progress-circular>
            </v-row>

            <v-tooltip bottom v-else>
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
    </v-btn>
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
