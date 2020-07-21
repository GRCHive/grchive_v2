<template>
    <v-rating
        :value="rating"
        readonly
        color="error"
        size="20"
    >
    </v-rating>
</template>

<script lang="ts">
import Vue from 'vue'
import Component from 'vue-class-component'

@Component
export default class RatingRenderer extends Vue {
    params : any = null

    get rating() : number {
        return this.params!.value
    }

    refresh(params:  any) {
        this.params = params
        return true
    }
}
</script>
