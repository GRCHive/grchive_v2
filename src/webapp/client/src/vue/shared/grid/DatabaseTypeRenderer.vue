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
        const typeId = this.params.data.TypeId || this.params.data.Data.TypeId
        const otherType = this.params.data.OtherType || this.params.data.Data.OtherType
        switch (typeId) {
            case DatabaseType.Other:
                return `Other (${otherType})`
            default:
                return databaseTypeToString(typeId)
        }
    }

    refresh(params:  any) {
        this.params = params
        return true
    }
}

</script>
