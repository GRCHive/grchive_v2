<template>
    <div>
        <v-text-field
            label="Identifier"
            filled
            v-model="value.HumanId"
            :readonly="readonly"
            :rules="[ rules.required, rules.createMaxLength(256) ]"
        >
        </v-text-field>

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

        <div class="d-flex align-center mb-4">
            <span class="font-weight-bold">Severity:</span>

            <v-rating
                v-model="value.Severity"
                :readonly="readonly"
                color="error"
                size="30"
            >
            </v-rating>
        </div>

        <v-textarea
            label="Severity Justification"
            filled
            v-model="value.SeverityJustification"
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
import { RawRisk } from '@client/ts/types/risks'
import * as rules from '@client/ts/frontend/formRules'

@Component
export default class RiskForm extends Vue {
    readonly rules : any = rules

    @Prop({ required: true })
    value! : RawRisk

    @Prop({ type: Boolean, default: false })
    readonly! : boolean
}

</script>
