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
        const freqType = this.params.data.FreqType || this.params.data.Data.FreqType
        const freqInterval = this.params.data.FreqInterval || this.params.data.Data.FreqInterval
        const freqOther = this.params.data.FreqOther || this.params.data.Data.FreqOther

        switch (freqType) {
            case ControlFrequencyType.CFTAdHoc:
                return 'Ad-Hoc'
            case ControlFrequencyType.CFTInterval:
                return Case.title(RRule.fromString(freqInterval).toText())
            case ControlFrequencyType.CFTOther:
                return freqOther
        }
        return ''
    }

    refresh(params:  any) {
        this.params = params
        return true
    }
}

</script>
