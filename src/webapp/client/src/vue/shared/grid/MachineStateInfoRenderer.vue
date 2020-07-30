<template>
    <span>
        {{ infoStr }}
    </span>
</template>

<script lang="ts">

import Vue from 'vue'
import Component from 'vue-class-component'
import { MachineState } from '@client/ts/types/machineState'

@Component
export default class MachineStateTypeRenderer extends Vue {
    params : any = null

    get infoStr() : string {
        const state : MachineState = this.params.value
        if (state.Hypervisor !== null) {
            return state.Hypervisor
        } else if (state.OperatingSystem !== null) {
            return state.OperatingSystem
        } else {
            return 'N/A'
        }
        return ''
    }

    refresh(params:  any) {
        this.params = params
        return true
    }
}

</script>
