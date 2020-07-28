<template>
    <span>
        {{ typeStr }}
    </span>
</template>

<script lang="ts">

import Vue from 'vue'
import Component from 'vue-class-component'
import { DatabaseType, databaseTypeToString } from '@client/ts/types/databases'

@Component
export default class DatabaseTypeRenderer extends Vue {
    params : any = null

    get typeStr() : string {
        switch (this.params.data.TypeId) {
            case DatabaseType.Other:
                return `Other (${this.params.data.OtherType})`
            default:
                return databaseTypeToString(this.params.data.TypeId)
        }
    }

    refresh(params:  any) {
        this.params = params
        return true
    }
}

</script>
