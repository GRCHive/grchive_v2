<template>
    <div>
        <slot v-bind:show="!!hasPermissions"></slot>
        <v-overlay
            absolute
            :value="!hasPermissions"
        >
            <v-row justify="center" v-if="isLoading">
                <v-progress-circular size=64 indeterminate>
                </v-progress-circular>
            </v-row>

            <div v-else>
                {{ tooltipStr }} 
            </div>
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

