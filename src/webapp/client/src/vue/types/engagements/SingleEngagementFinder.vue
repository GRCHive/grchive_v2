<template>
    <v-autocomplete
        :label="label"
        :value="value"
        filled
        :rules="rules"
        :readonly="readonly"
        :loading="!validEngagements"
        :items="engagementItems"
        :clearable="clearable && !readonly"
        @input="onInput"
    >
    </v-autocomplete>
</template>

<script lang="ts">

import Vue from 'vue'
import Component, { mixins } from 'vue-class-component'
import { Watch, Prop } from 'vue-property-decorator'
import { VAutocomplete } from 'vuetify/lib'
import { RawEngagement } from '@client/ts/types/engagements'
import { GrchiveApi, ErrorHandler } from '@client/ts/main'

@Component
export default class SingleEngagementFinder extends mixins(VAutocomplete) {
    @Prop({ required: true })
    orgId! : number

    validEngagements: RawEngagement[] | null = null

    @Watch('orgId')
    refreshValidEngagements() {
        GrchiveApi.engagements.listOrgEngagements(this.orgId).then((resp : RawEngagement[]) => {
            this.validEngagements = resp
        }).catch((err : any) => {
            ErrorHandler.failurePopupOnError(err, {
                context: 'Failed to get all available engagements.'
            })
        })
    }

    mounted() {
        this.refreshValidEngagements()
    }

    onInput(v : RawEngagement[]) {
        this.$emit('input', v)
    }

    get engagementItems() : any[] {
        if (!this.validEngagements) {
            return []
        }
        return this.validEngagements.map((ele : RawEngagement) => ({
            text : ele.Name,
            value: ele,
        })
    }
}

</script>
