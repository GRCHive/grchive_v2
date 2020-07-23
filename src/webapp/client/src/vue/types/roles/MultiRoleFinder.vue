<template>
    <v-autocomplete
        label="Roles"
        :value="value"
        @input="onInput"
        chips
        clearable
        deletable-chips
        filled
        :loading="!validRoles"
        :items="roleItems"
        :rules="rules"
        :readonly="readonly"
        multiple
    >
    </v-autocomplete>
</template>

<script lang="ts">

import Vue from 'vue'
import Component, { mixins } from 'vue-class-component'
import { VAutocomplete } from 'vuetify/lib'
import { Prop } from 'vue-property-decorator'
import { GrchiveApi } from '@client/ts/main'
import { Role } from '@client/ts/types/roles'

@Component
export default class MultiRoleFinder extends mixins(VAutocomplete) {
    @Prop({ required: true })
    orgId! : number

    validRoles : Role[] | null = null

    get roleItems() : any[] {
        if (!this.validRoles) {
            return []
        }
        return this.validRoles.map((ele : Role) => ({
            text : ele.Name,
            value: ele,
        }))
    }

    refreshValidRoles() {
        GrchiveApi.roles.listRoles(this.orgId).then((resp : Role[]) => {
            this.validRoles = resp
        })
    }

    mounted() {
        this.refreshValidRoles()
    }

    onInput(v : Role[]) {
        this.$emit('input', v)
    }
}

</script>
