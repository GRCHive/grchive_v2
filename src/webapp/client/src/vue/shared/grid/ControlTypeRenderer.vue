<template>
    <span>
        {{ typeStr }}
    </span>
</template>

<script lang="ts">

import Vue from 'vue'
import Component from 'vue-class-component'
import { controlTypeToString } from '@client/ts/types/controls'

@Component
export default class ControlTypeRenderer extends Vue {
    params : any = null

    get typeStr() : string {
        return controlTypeToString(this.params!.value)
    }

    refresh(params:  any) {
        this.params = params
        return true
    }
}

</script>
