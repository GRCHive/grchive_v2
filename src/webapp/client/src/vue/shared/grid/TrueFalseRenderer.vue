<template>
    <div>
        <div v-if="tf">
            <v-icon
                small
                color="success"
            >
                mdi-check-circle
            </v-icon>
        </div>

        <div v-else>
            <v-icon
                small
                color="error"
            >
                mdi-close-circle
            </v-icon>
        </div>
    </div>
</template>

<script lang="ts">

import Vue from 'vue'
import Component from 'vue-class-component'

@Component
export default class TrueFalseRenderer extends Vue {
    params : any = null

    get tf() : boolean {
        return this.params!.value
    }

    refresh(params:  any) {
        this.params = params
        return true
    }
}

</script>

