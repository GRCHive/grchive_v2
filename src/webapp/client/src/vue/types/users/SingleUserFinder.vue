<template>
    <v-autocomplete
        :label="label"
        :value="value"
        @input="onInput"
        chips
        clearable
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
import { Prop } from 'vue-property-decorator'
import { VAutocomplete } from 'vuetify/lib'
import { RawUser } from '@client/ts/types/users'
import { GrchiveApi } from '@client/ts/main'

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
        })
    }

    refreshValidUsers() {
        GrchiveApi.orgs.getUsersInOrg(this.orgId).then((resp : RawUser[] | null) => {
            this.validUsers = resp
        })
    }

    mounted() {
        this.refreshValidUsers()
    }

    onInput(v : RawUser[]) {
        this.$emit('input', v)
    }
}


</script>
