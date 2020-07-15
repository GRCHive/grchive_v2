<template>
    <div>
        <v-text-field
            label="Name"
            filled
            v-model="value.Name"
            :readonly="readonly"
            :rules="[ rules.required, rules.createMaxLength(256) ]"
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
import { RawOrganization } from '@client/ts/types/orgs'
import * as rules from '@client/ts/frontend/formRules'

@Component
export default class OrgForm extends Vue {
    readonly rules : any = rules

    @Prop({ required: true })
    value! : RawOrganization

    @Prop({ default: false })
    readonly! : boolean
}

</script>
