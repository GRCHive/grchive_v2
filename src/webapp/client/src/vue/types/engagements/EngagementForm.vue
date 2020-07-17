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
        >
        </v-textarea>
    </div>
</template>

<script lang="ts">

import Vue from 'vue'
import Component from 'vue-class-component'
import { Prop } from 'vue-property-decorator'
import { RawEngagement } from '@client/ts/types/engagements'
import * as rules from '@client/ts/frontend/formRules'

@Component
export default class OrgForm extends Vue {
    readonly rules : any = rules

    @Prop({ required: true })
    value! : RawEngagement

    @Prop({ type: Boolean, default: false })
    readonly! : boolean
}

</script>
