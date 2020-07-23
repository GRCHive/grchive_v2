<template>
    <div>
        <v-dialog
            persistent
            max-width="60%"
            v-model="showHideMoreInfo"
        >
            <v-card>
                <v-card-title>
                    Error Information
                </v-card-title>
                <v-divider></v-divider>

                <error-wrapper-display
                    :current-error="currentError"
                    class="mx-4"
                >
                </error-wrapper-display>

                <v-card-actions>
                    <v-btn
                        color="error"
                        @click="showHideMoreInfo = false"
                    >
                        Close
                    </v-btn>

                    <v-spacer></v-spacer>

                    <v-btn
                        color="success"
                        href="mailto:support@grchive.com"
                    >
                        Help!
                    </v-btn>
                </v-card-actions>
            </v-card>
        </v-dialog>

        <v-snackbar
            :value="!!currentError"
            multi-line
            bottom
            :timeout="-1"
        >
            {{ currentErrorDisplayText }}

            <template v-slot:action="{ attrs }">
                <v-btn icon @click="prevError" :disabled="!canGoPrev">
                    <v-icon>
                        mdi-chevron-left
                    </v-icon>
                </v-btn>

                <span class="text-overline" style="flex-shrink: 0;">
                    {{ displayedIndex + 1 }} / {{ allErrors.length }}
                </span>

                <v-btn icon @click="nextError" :disabled="!canGoNext">
                    <v-icon>
                        mdi-chevron-right
                    </v-icon>
                </v-btn>

                <v-btn
                    color="info"
                    v-bind="attrs"
                    @click="showHideMoreInfo = true"
                >
                    Info
                </v-btn>

                <v-btn
                    color="error"
                    v-bind="attrs"
                    @click="closeCurrentError"
                    class="ml-4"
                >
                    Close
                </v-btn>
            </template>
        </v-snackbar>
    </div>
</template>

<script lang="ts">

import Vue from 'vue'
import Component from 'vue-class-component'
import { Watch } from 'vue-property-decorator'
import { ErrorWrapper } from '@client/ts/error'
import ErrorWrapperDisplay from '@client/vue/ErrorWrapperDisplay.vue'

const errorDurationSeconds = 50000

@Component({
    components: {
        ErrorWrapperDisplay,
    }
})
export default class ErrorPopupManager extends Vue {
    handledErrors : Set<string> = new Set<string>()
    displayedIndex : number = 0
    showHideMoreInfo:  boolean = false

    get allErrors() : ErrorWrapper[] {
        return this.$store.state.errors.relevantErrors
    }

    get currentError() : ErrorWrapper | null {
        if (this.displayedIndex < 0 || this.displayedIndex >= this.allErrors.length) {
            return null
        }
        return this.allErrors[this.displayedIndex]
    }

    get currentErrorDisplayText() : string {
        if (!this.currentError) {
            return ''
        }

        return `Error Code: ${this.currentError.code}. ${this.currentError.context}`
    }

    get canGoNext() : boolean {
        return this.displayedIndex < this.allErrors.length - 1
    }

    nextError() {
        this.displayedIndex += 1
    }

    get canGoPrev() : boolean {
        return this.displayedIndex > 0
    }

    prevError() {
        this.displayedIndex -= 1
    }

    closeCurrentError() {
        if (!this.currentError) {
            return
        }
        this.$store.commit('errors/removeError', this.currentError.displayId)
    }

    @Watch('allErrors')
    startErrorClearTimers() {
        for (let e of this.allErrors) {
            if (this.handledErrors.has(e.displayId)) {
                continue
            }
            this.handledErrors.add(e.displayId)

            setTimeout(() => {
                this.$store.commit('errors/removeError', e.displayId)
            }, errorDurationSeconds * 1000)
        }
    }
}

</script>

<style scoped>

>>>.v-snack__wrapper {
    max-width: 80%;
}

</style>
