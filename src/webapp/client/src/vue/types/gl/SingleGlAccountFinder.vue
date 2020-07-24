<template>
    <v-autocomplete
        :label="label"
        :value="value"
        @input="onInput"
        chips
        :clearable="!readonly"
        deletable-chips
        filled
        :loading="!validAccounts"
        :items="accItems"
        :rules="rules"
        :readonly="readonly"
        :hide-details="hideDetails"
    >
    </v-autocomplete>
</template>

<script lang="ts">

import Vue from 'vue'
import Component, { mixins } from 'vue-class-component'
import { Prop, Watch } from 'vue-property-decorator'
import { VAutocomplete } from 'vuetify/lib'
import { RawGLAccount } from '@client/ts/types/gl'
import { GrchiveApi } from '@client/ts/main'

@Component
export default class SingleGlAccountFinder extends mixins(VAutocomplete) {
    @Prop({ required: true })
    orgId! : number

    @Prop({ required: true })
    engagementId! : number

    validAccounts: RawGLAccount[] | null = null

    @Watch('orgId')
    @Watch('engagementId')
    refreshValidAccounts() {
        GrchiveApi.gl.listAccounts(this.orgId, this.engagementId).then((resp : RawGLAccount[]) => {
            this.validAccounts = resp
        })
    }

    mounted() {
        this.refreshValidAccounts()
    }

    onInput(v : RawGLAccount[]) {
        this.$emit('input', v)
    }

    get accItems() : any[] {
        if (!this.validAccounts) {
            return []
        }
        return this.validAccounts.map((ele : RawGLAccount) => ({
            text : `${ele.AccountId}: (${ele.Name})`,
            value: ele,
        }))
    }
}


</script>
