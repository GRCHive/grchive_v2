<template>
    <v-app>
        <v-main class="max-height">
            <v-container class="max-height">
                <v-row class="max-height" justify="center" align="center">
                    <div id="errorContainer" v-if="!currentError">
                        <p class="text-h1">
                            404
                        </p>

                        <p class="text-h6">
                            Oops! You weren't support to be able to get here.
                            Click&nbsp; <a @click="goBack">here</a>&nbsp;to go back to where you were before.
                        </p>
                    </div>

                    <div id="errorContainer" v-else>
                        <p class="text-h1">
                            {{ currentError.httpCode }}
                        </p>

                        <div class="text-h6">
                            <error-wrapper-display
                                :current-error="currentError"
                                show-go-back
                            >
                            </error-wrapper-display>
                        </div>

                    </div>
                </v-row>
            </v-container>
        </v-main>
    </v-app>
</template>

<script lang="ts">

import Vue from 'vue'
import Component from 'vue-class-component'
import { Watch } from 'vue-property-decorator'
import { GrchiveErrorCodes, ErrorWrapper } from '@client/ts/error'
import ErrorWrapperDisplay from '@client/vue/ErrorWrapperDisplay.vue'

@Component({
    components: {
        ErrorWrapperDisplay,
    }
})
export default class ErrorPage extends Vue {

    @Watch('$route')
    get currentError() : ErrorWrapper | null {
        const code = this.$route.query.code
        const msg = this.$route.query.message
        const technical = this.$route.query.technical
        const httpCode = this.$route.query.httpCode

        if (code === null || code === undefined ||
            msg === null || msg === undefined ||
            technical === null || technical === undefined ||
            httpCode === null || httpCode === null
        ) {
            return null
        }

        return new ErrorWrapper(<GrchiveErrorCodes>Number(code), <string>msg, JSON.parse(atob(<string>technical)), Number(httpCode))
    }

    goBack() {
        this.$router.back()
    }

    mounted() {
        document.title = 'GRCHive Error! Oh no!'
    }
}

</script>

<style scoped>

#errorContainer {
    max-width: 800px;
    text-align: center;
}

pre {
    white-space: pre-wrap;
    word-wrap: break-word;
}

</style>
