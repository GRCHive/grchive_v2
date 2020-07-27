<template>
    <div>
        <inventory-form
            v-model="value.Inventory"
            :readonly="readonly"
        >
        </inventory-form>

        <v-text-field
            label="Physical Location"
            filled
            v-model="value.PhysicalLocation"
            :readonly="readonly"
        >
        </v-text-field>

        <v-text-field
            label="Operating System"
            filled
            v-model="value.OperatingSystem"
            :readonly="readonly"
        >
        </v-text-field>

        <v-text-field
            label="Hypervisor"
            filled
            v-model="value.Hypervisor"
            :readonly="readonly"
        >
        </v-text-field>

        <v-text-field
            label="Static External IP"
            filled
            v-model="value.StaticExternalIp"
            :readonly="readonly"
        >
        </v-text-field>
    </div>
</template>

<script lang="ts">

import Vue from 'vue'
import Component from 'vue-class-component'
import { Prop } from 'vue-property-decorator'
import { RawServer } from '@client/ts/types/inventory'
import * as rules from '@client/ts/frontend/formRules'
import InventoryForm from '@client/vue/types/inventory/InventoryForm.vue'

@Component({
    components: {
        InventoryForm,
    }
})
export default class ServerForm extends Vue {
    readonly rules : any = rules

    @Prop({ required: true })
    value! : RawServer

    @Prop({ type: Boolean, default: false })
    readonly! : boolean
}

</script>
