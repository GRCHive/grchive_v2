<template>
    <span>
        {{ dateTimeStr }}
    </span>
</template>

<script lang="ts">

import Vue from 'vue'
import Component from 'vue-class-component'
import { standardFormatTime } from '@client/ts/time'

@Component
export default class SimpleDateTimeRenderer extends Vue {
    params : any = null

    get dateTimeStr() : string {
        return standardFormatTime(this.params.value)
    }

    refresh(params:  any) {
        this.params = params
        return true
    }
}

</script>
