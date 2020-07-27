<template>
    <div>
        <v-text-field
            label="Name"
            filled
            v-model="value.Name"
            :readonly="readonly"
            :rules="[ rules.required ]"
        >
        </v-text-field>

        <v-textarea
            label="Description"
            filled
            v-model="value.Description"
            :readonly="readonly"
            hide-details
            class="mb-4"
        >
        </v-textarea>

        <v-text-field
            label="Brand"
            filled
            v-model="value.Brand"
            :readonly="readonly"
        >
        </v-text-field>

        <v-text-field
            label="Model"
            filled
            v-model="value.Model"
            :readonly="readonly"
        >
        </v-text-field>
    </div>
</template>

<script lang="ts">

import Vue from 'vue'
import Component from 'vue-class-component'
import { Prop } from 'vue-property-decorator'
import { RawInventory } from '@client/ts/types/inventory'
import * as rules from '@client/ts/frontend/formRules'

@Component
export default class InventoryForm extends Vue {
    readonly rules : any = rules

    @Prop({ required: true })
    value! : RawInventory

    @Prop({ type: Boolean, default: false })
    readonly! : boolean
}

</script>
