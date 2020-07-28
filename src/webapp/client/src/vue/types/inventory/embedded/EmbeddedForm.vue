<template>
    <div>
        <inventory-form
            v-model="value.Inventory"
            :readonly="readonly"
        >
        </inventory-form>

        <v-text-field
            label="Operating System"
            filled
            v-model="value.OperatingSystem"
            :readonly="readonly"
        >
        </v-text-field>
    </div>
</template>

<script lang="ts">

import Vue from 'vue'
import Component from 'vue-class-component'
import { Prop } from 'vue-property-decorator'
import { RawEmbedded } from '@client/ts/types/inventory'
import * as rules from '@client/ts/frontend/formRules'
import InventoryForm from '@client/vue/types/inventory/InventoryForm.vue'

@Component({
    components: {
        InventoryForm,
    }
})
export default class EmbeddedForm extends Vue {
    readonly rules : any = rules

    @Prop({ required: true })
    value! : RawEmbedded

    @Prop({ type: Boolean, default: false })
    readonly! : boolean
}

</script>
