<template>
    <v-autocomplete
        :label="label"
        :value="value"
        @input="onInput"
        chips
        :clearable="!readonly"
        deletable-chips
        filled
        :loading="!validRisks"
        :items="riskItems"
        :rules="rules"
        :readonly="readonly"
    >
    </v-autocomplete>
</template>

<script lang="ts">

import Vue from 'vue'
import Component, { mixins } from 'vue-class-component'
import { Prop, Watch } from 'vue-property-decorator'
import { VAutocomplete } from 'vuetify/lib'
import { RawRisk } from '@client/ts/types/risks'
import { GrchiveApi, ErrorHandler } from '@client/ts/main'

@Component
export default class SingleRiskFinder extends mixins(VAutocomplete) {
    @Prop({ required: true })
    orgId! : number

    @Prop({ required: true })
    engagementId! : number

    validRisks: RawRisk[] | null = null

    get riskItems() : any[] {
        if (!this.validRisks) {
            return []
        }
        return this.validRisks.map((ele : RawRisk) => ({
            text : `${ele.HumanId}: ${ele.Name}`,
            value: ele,
        }))
    }

    @Watch('orgId')
    @Watch('engagementId')
    refreshValidRisks() {
        GrchiveApi.risks.listRisks(this.orgId, this.engagementId).then((resp : RawRisk[]) => {
            this.validRisks = resp
        }).catch((err : any) => {
            ErrorHandler.failurePopupOnError(err, {
                context: 'Failed to list risks to select.'
            })
        })
    }

    mounted() {
        this.refreshValidRisks()
    }

    onInput(v : RawRisk) {
        this.$emit('input', v)
    }
}

</script>
