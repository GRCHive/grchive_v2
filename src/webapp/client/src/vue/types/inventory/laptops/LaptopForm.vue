<template>
    <div>
        <inventory-form
            v-model="value.Inventory"
            :readonly="readonly"
        >
        </inventory-form>

        <machine-state-form
            v-model="value.State"
            :readonly="readonly"
        >
        </machine-state-form>
    </div>
</template>

<script lang="ts">

import Vue from 'vue'
import Component from 'vue-class-component'
import { Prop } from 'vue-property-decorator'
import { RawLaptop } from '@client/ts/types/inventory'
import * as rules from '@client/ts/frontend/formRules'
import InventoryForm from '@client/vue/types/inventory/InventoryForm.vue'
import MachineStateForm from '@client/vue/types/machineState/MachineStateForm.vue'

@Component({
    components: {
        InventoryForm,
        MachineStateForm
    }
})
export default class LaptopForm extends Vue {
    readonly rules : any = rules

    @Prop({ required: true })
    value! : RawLaptop

    @Prop({ type: Boolean, default: false })
    readonly! : boolean
}

</script>
