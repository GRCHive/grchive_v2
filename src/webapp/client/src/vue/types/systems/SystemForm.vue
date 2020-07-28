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
            label="Purpose"
            filled
            v-model="value.Purpose"
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
import { RawSystem } from '@client/ts/types/systems'
import * as rules from '@client/ts/frontend/formRules'

@Component
export default class SystemForm extends Vue {
    readonly rules : any = rules

    @Prop({ required: true })
    value! : RawSystem

    @Prop({ type: Boolean, default: false })
    readonly! : boolean
}

</script>
