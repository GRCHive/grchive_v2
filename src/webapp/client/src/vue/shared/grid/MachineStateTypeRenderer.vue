<template>
    <span>
        {{ typeStr }}
    </span>
</template>

<script lang="ts">

import Vue from 'vue'
import Component from 'vue-class-component'
import { MachineState } from '@client/ts/types/machineState'

@Component
export default class MachineStateInfoRenderer extends Vue {
    params : any = null

    get typeStr() : string {
        const state : MachineState = this.params.value
        if (state.Hypervisor !== null) {
            return 'Hypervisor'
        } else if (state.OperatingSystem !== null) {
            return 'Operating System'
        } else {
            return 'None'
        }
        return ''
    }

    refresh(params:  any) {
        this.params = params
        return true
    }
}

</script>
