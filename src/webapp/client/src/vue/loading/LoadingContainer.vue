<template>
    <div>
        <slot v-if="!loading"></slot>
        <v-row justify="center" v-else>
            <v-progress-circular size=64 indeterminate>
            </v-progress-circular>
        </v-row>
    </div>
</template>


<script lang="ts">

import Vue from 'vue'
import Component from 'vue-class-component'
import { Prop } from 'vue-property-decorator'

@Component
export default class LoadingContainer extends Vue {
    @Prop({required: true})
    loading! : boolean
}

</script>
