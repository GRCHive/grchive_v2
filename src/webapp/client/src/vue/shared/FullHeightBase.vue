<template>
    <div ref="baseDiv" :style="style">
        <slot></slot>
    </div>
</template>

<script lang="ts">

import Vue from 'vue'
import Component from 'vue-class-component'

@Component
export default class FullHeightBase extends Vue {
    $refs! : {
        baseDiv: HTMLElement
    }

    divHeight : number = 0

    get style() : any {
        return {
            height: `${this.divHeight}px`,
            overflow: 'auto',
        }
    }

    refreshDimensions() {
        const bb = this.$refs.baseDiv.getBoundingClientRect()
        this.divHeight = window.innerHeight - bb.top
    }

    mounted() {
        this.refreshDimensions()
        window.addEventListener('resize', this.refreshDimensions)
    }
}

</script>
