<template>
    <v-tab
        :to="finalTo"
    >
        <v-row justify="center" v-if="isLoading">
            <v-progress-circular size=16 indeterminate>
            </v-progress-circular>
        </v-row>
        <div v-else>
            <slot></slot>

            <v-overlay
                absolute
                :value="!hasPermissions"
            >
                <div>
                    <v-tooltip max-width="400px" bottom>
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
    </v-tab>

</template>

<script lang="ts">

import RestrictRolePermissionBase from '@client/vue/loading/RestrictRolePermissionBase.vue'
import Component, { mixins } from 'vue-class-component'
import { VTab } from 'vuetify/lib'

@Component
export default class RestrictRolePermissionTab extends mixins(VTab, RestrictRolePermissionBase) {

    // This is sort of a hack because the 'disabled' property on the vuetify tab doesn't seem to work.
    get finalTo() : any {
        if (this.isLoading || !this.hasPermissions) {
            return {
                path: '#',
            }
        } else {
            //@ts-ignore
            return this.to
        }
    }
}

</script>
