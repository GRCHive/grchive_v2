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

        <v-text-field
            label="URL"
            filled
            v-model="value.Url"
            :readonly="readonly"
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
    </div>
</template>

<script lang="ts">

import Vue from 'vue'
import Component from 'vue-class-component'
import { Prop } from 'vue-property-decorator'
import { RawVendorProduct } from '@client/ts/types/vendors'
import * as rules from '@client/ts/frontend/formRules'

@Component
export default class VendorProductForm extends Vue {
    readonly rules : any = rules

    @Prop({ required: true })
    value! : RawVendorProduct

    @Prop({ type: Boolean, default: false })
    readonly! : boolean
}

</script>
