<template>
    <div>
        <v-select
            v-model="machineType"
            label="Machine State Type"
            :items="machineTypeItems"
            filled
            :value="machineType"
            :readonly="readonly"
        >
        </v-select>

        <v-text-field
            label="Operating System"
            filled
            v-model="value.OperatingSystem"
            :readonly="readonly"
            v-if="value.OperatingSystem !== null"
        >
        </v-text-field>

        <v-text-field
            label="Hypervisor"
            filled
            v-model="value.Hypervisor"
            :readonly="readonly"
            v-if="value.Hypervisor !== null"
        >
        </v-text-field>
    </div>
</template>

<script lang="ts">

import Vue from 'vue'
import Component from 'vue-class-component'
import { Watch, Prop } from 'vue-property-decorator'
import { MachineState, MachineStateType } from '@client/ts/types/machineState'

@Component
export default class MachineStateForm extends Vue {
    @Prop({ required: true })
    value! : MachineState
    initialSyncDone: boolean = false

    @Prop({ type: Boolean, default: false })
    readonly! : boolean

    machineType: MachineStateType = MachineStateType.None
    get machineTypeItems() : any[] {
        return [
            {
                text: 'None',
                value: MachineStateType.None,
            },
            {
                text: 'Hypervisor',
                value: MachineStateType.Hypervisor,
            },
            {
                text: 'Operating System',
                value: MachineStateType.OperatingSystem,
            },
        ]
    }

    @Watch('value')
    syncMachineType() {
        this.initialSyncDone = false
        if (this.value.Hypervisor !== null) {
            this.machineType = MachineStateType.Hypervisor
        } else if (this.value.OperatingSystem !== null) {
            this.machineType = MachineStateType.OperatingSystem
        } else {
            this.machineType = MachineStateType.None
        }

        Vue.nextTick(() => {
            this.initialSyncDone = true
        })
    }

    @Watch('machineType')
    onChangeMachineType(newVal : MachineStateType, oldVal : MachineStateType) {
        if (newVal == oldVal || !this.initialSyncDone) {
            return
        }

        this.machineType = newVal
        switch (newVal) {
            case MachineStateType.None:
                this.value.Hypervisor = null
                this.value.OperatingSystem = null
                break
            case MachineStateType.Hypervisor:
                this.value.Hypervisor = ''
                this.value.OperatingSystem = null
                break
            case MachineStateType.OperatingSystem:
                this.value.Hypervisor = null
                this.value.OperatingSystem = ''
                break
        }
    }

    mounted() {
        this.syncMachineType()
    }
}

</script>
