<template>
    <div>
        <p>
            Oops! Something went wrong. Here's what we know:
        </p>

        <p>
            <span class="font-weight-bold">GRCHive Error Code:</span> {{ currentError.code }}
        </p>

        <p>
            <span class="font-weight-bold">GRCHive Error Message:</span> {{ currentError.message }}
        </p>

        <p>
            Not helpful? Send us the following text along with the above error code and message and we'll help you figure it out!
        </p>
        <pre>{{ currentError.technical }}</pre>

        <p>
            <a href="mailto:support@grchive.com">Contact Support.</a>
        </p>

        <p v-if="showGoBack">
            Click&nbsp; <a @click="goBack">here</a>&nbsp;to go back to where you were before.
        </p>
    </div>
</template>

<script lang="ts">

import Vue from 'vue'
import Component from 'vue-class-component'
import { Prop } from 'vue-property-decorator'
import { ErrorWrapper } from '@client/ts/error'

@Component
export default class ErrorWrapperDisplay extends Vue {
    @Prop({ required: true })
    currentError! : ErrorWrapper

    @Prop({ type: Boolean, default: false })
    showGoBack!: boolean
}

</script>
