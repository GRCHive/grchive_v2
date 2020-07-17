<template>
    <v-list-item
        :to="{name: `${item.path}` }"
        link
        :color="item.disabled ? `secondary` : `primary`"
        :disabled="item.disabled"
        :two-line="item.disabled"
        v-if="!item.hidden"
    >
        <v-list-item-icon v-if="item.icon != ''">
            <v-icon>{{ item.icon }}</v-icon>
        </v-list-item-icon>

        <v-list-item-content>
            <v-list-item-title>
                {{ item.title }}
            </v-list-item-title>
            <v-list-item-subtitle v-if="item.disabled">
                Coming Soon.
            </v-list-item-subtitle>
        </v-list-item-content>
    </v-list-item>

</template>

<script lang="ts">

import Vue from 'vue'
import Component from 'vue-class-component'
import { Prop } from 'vue-property-decorator'
import { NavLink } from '@client/ts/frontend/navLink'

@Component
export default class GenericNavBarItem extends Vue {
    @Prop()
    item! : NavLink
}

</script>
