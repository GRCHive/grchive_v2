<template>
    <v-autocomplete
        :label="label"
        :value="value"
        @input="onInput"
        chips
        :clearable="!readonly"
        deletable-chips
        filled
        :loading="!validUsers"
        :items="userItems"
        :rules="rules"
        :readonly="readonly"
    >
    </v-autocomplete>
</template>

<script lang="ts">

import Vue from 'vue'
import Component, { mixins } from 'vue-class-component'
import { Prop, Watch } from 'vue-property-decorator'
import { VAutocomplete } from 'vuetify/lib'
import { RawUser } from '@client/ts/types/users'
import { GrchiveApi, ErrorHandler } from '@client/ts/main'

@Component
export default class SingleUserFinder extends mixins(VAutocomplete) {
    @Prop({ required: true })
    orgId! : number

    validUsers: RawUser[] | null = null

    get userItems() : any[] {
        if (!this.validUsers) {
            return []
        }
        return this.validUsers.map((ele : RawUser) => ({
            text : `${ele.FullName} (${ele.Email})`,
            value: ele,
        }))
    }

    @Watch('orgId')
    refreshValidUsers() {
        GrchiveApi.orgs.getUsersInOrg(this.orgId).then((resp : RawUser[]) => {
            this.validUsers = resp
        }).catch((err : any) => {
            ErrorHandler.failurePopupOnError(err, {
                context: 'Failed to list userse to select.'
            })
        })
    }

    mounted() {
        this.refreshValidUsers()
    }

    onInput(v : RawUser) {
        this.$emit('input', v)
    }
}


</script>
