<template>
    <v-alert
        type="warning"
        :value="showBanner"
        @input="closeBanner"
        dismissible
        class="pa-2 ma-0"
        width="100%"
        tile
    >
        <template v-slot:prepend>
            <v-icon>mdi-exclamation</v-icon>
        </template>

        <v-row align="center">
            <v-col class="pl-4 py-0 flex-grow-1 flex-shrink-0">
                Please verify your email address to access all available functionality.
            </v-col>

            <v-col class="py-0 pr-2 flex-grow-0 flex-shrink-1">
                <v-btn small outlined @click="sendVerification" v-if="!sentEmail">
                    Resend Verification Email
                </v-btn>

                <v-btn small outlined disabled v-else>
                    Email Sent!
                </v-btn>
            </v-col>
        </v-row>
    </v-alert>
</template>

<script lang="ts">

import Vue from 'vue'
import Component from 'vue-class-component'
import { GrchiveApi } from '@client/ts/api/client'

@Component
export default class VerifyEmailBanner extends Vue {
    sentEmail : boolean = false

    get showBanner() : boolean {
        return this.$store.getters['appLayout/showEmailVerificationBanner']
    }

    closeBanner() {
        this.$store.commit('appLayout/closeVerificationBanner')
    }

    sendVerification() {
        this.sentEmail = true
        GrchiveApi.user.resendEmailVerification().catch((err : any) => {
            this.$store.commit('errors/addApiError', err)
        })
    }
}

</script>
