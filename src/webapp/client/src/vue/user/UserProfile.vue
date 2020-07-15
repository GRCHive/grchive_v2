<template>
    <user-template
        page-name="Organizations"
        :relevant-user="currentUser"
    >
        <template v-slot:content>
            <v-list-item class="px-0">
                <v-list-item-content>
                    <v-list-item-title class="text-h4">
                        My Profile
                    </v-list-item-title>
                </v-list-item-content>
            </v-list-item>
            <v-divider class="mb-4"></v-divider>

            <v-form onSubmit="return false;" v-model="formValid">
                <v-text-field
                    label="Full Name"
                    filled
                    v-model="workingUser.FullName"
                    :readonly="!canEdit"
                    :rules="[ rules.required, rules.createMaxLength(512) ]"
                >
                </v-text-field>

                <v-text-field
                    label="Email"
                    readonly
                    filled
                    :value="workingUser.Email"
                >
                </v-text-field>

                <edit-state-buttons
                    v-model="canEdit"
                    :edit-pending="editPending"
                    @save-edit="saveEdit"
                    @cancel-edit="cancelEdit"
                >
                </edit-state-buttons>
            </v-form>

            <v-divider class="mb-4"></v-divider>

            <span>
                Email Verification:

                <template v-if="workingUser.EmailVerified">
                    <v-icon color="success">mdi-check-circle</v-icon>
                </template>

                <template v-else>
                    <v-icon color="error">mdi-close-circle</v-icon>
                    <v-btn color="warning" :disabled="verificationSent" @click="sendVerification">
                        Resend Verification Email
                    </v-btn>
                </template>
            </span>
        </template>
    </user-template>
</template>

<script lang="ts">

import Vue from 'vue'
import Component from 'vue-class-component'
import { Watch } from 'vue-property-decorator'
import { RawUser } from '@client/ts/types/users'
import { GrchiveApi } from '@client/ts/api/client'
import EditStateButtons from '@client/vue/shared/EditStateButtons.vue'
import UserTemplate from '@client/vue/UserTemplate.vue'
import * as rules from '@client/ts/frontend/formRules'

// TODO: #9mdmj0
// Support viewing other user profile pages.
@Component({
    components: {
        UserTemplate,
        EditStateButtons,
    }
})
export default class UserProfile extends Vue {
    readonly rules : any = rules
    canEdit: boolean = false
    editPending : boolean = false

    // Email verification. Only allow user to send the verification email once
    // every time they come to this page.
    verificationSent : boolean = false

    // Need to store a working copy of the user data that we use to edit.
    // Only dump it into the store once we finish editing.
    workingUser : RawUser | null = null

    sendVerification() {
        this.verificationSent = true
        GrchiveApi.user.resendEmailVerification().catch((err : any) => {
            this.$store.commit('errors/addApiError', err)
        })
    }

    get currentUser() : RawUser | null {
        return this.$store.state.user.rawUser
    }

    @Watch('currentUser')
    syncWorkingCopy() {
        this.workingUser = JSON.parse(JSON.stringify(this.currentUser))
    }

    saveEdit() {
        if (!this.workingUser) {
            return
        }

        this.editPending = true
        GrchiveApi.user.updateUser(this.workingUser!).then(() => {
            this.$store.commit('user/setRawUser', this.workingUser!)
        }).catch((err : any) => {
            this.$store.commit('errors/addApiError', err)
        }).finally(() => {
            this.editPending = false
        })
    }

    cancelEdit() {
        this.syncWorkingCopy()
    }
}

</script>
