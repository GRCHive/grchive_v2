<template>
    <span>
        {{ freqStr }}
    </span>
</template>

<script lang="ts">

import Vue from 'vue'
import Component from 'vue-class-component'
import { ControlFrequencyType } from '@client/ts/types/controls'
import { RRule } from 'rrule'
import Case from 'case'

@Component
export default class ControlFrequencyRenderer extends Vue {
    params : any = null

    get freqStr() : string {
        switch (this.params.data.FreqType) {
            case ControlFrequencyType.CFTAdHoc:
                return 'Ad-Hoc'
            case ControlFrequencyType.CFTInterval:
                return Case.title(RRule.fromString(this.params.data.FreqInterval).toText())
            case ControlFrequencyType.CFTOther:
                return this.params.data.FreqOther
        }
        return ''
    }

    refresh(params:  any) {
        this.params = params
        return true
    }
}

</script>
