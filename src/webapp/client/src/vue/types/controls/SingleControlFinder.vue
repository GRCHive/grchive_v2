<template>
    <v-autocomplete
        :label="label"
        :value="value"
        @input="onInput"
        chips
        :clearable="!readonly"
        deletable-chips
        filled
        :loading="!validControls"
        :items="controlItems"
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
import { RawControl } from '@client/ts/types/controls'
import { GrchiveApi, ErrorHandler } from '@client/ts/main'

@Component
export default class SingleControlFinder extends mixins(VAutocomplete) {
    @Prop({ required: true })
    orgId! : number

    @Prop({ required: true })
    engagementId! : number

    validControls: RawControl[] | null = null

    get controlItems() : any[] {
        if (!this.validControls) {
            return []
        }
        return this.validControls.map((ele : RawControl) => ({
            text : `${ele.HumanId}: ${ele.Name}`,
            value: ele,
        }))
    }

    @Watch('orgId')
    @Watch('engagementId')
    refreshValidControls() {
        GrchiveApi.controls.listControls(this.orgId, this.engagementId).then((resp : RawControl[]) => {
            this.validControls = resp
        }).catch((err : any) => {
            ErrorHandler.failurePopupOnError(err, {
                context: 'Failed to list controls to select.'
            })
        })
    }

    mounted() {
        this.refreshValidControls()
    }

    onInput(v : RawControl) {
        this.$emit('input', v)
    }
}

</script>
