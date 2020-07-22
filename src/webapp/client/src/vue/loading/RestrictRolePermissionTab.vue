<template>
    <div style="position: relative; display: flex;">
        <v-tab
            :to="to"
        >
            <slot></slot>
        </v-tab>

        <v-overlay
            absolute
            :value="isLoading || !hasPermissions"
        >
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
        </v-overlay>
    </div>
</template>

<script lang="ts">

import RestrictRolePermissionBase from '@client/vue/loading/RestrictRolePermissionBase.vue'
import Component, { mixins } from 'vue-class-component'
import { VTab } from 'vuetify/lib'

@Component
export default class RestrictRolePermissionTab extends mixins(VTab, RestrictRolePermissionBase) {
}

</script>
