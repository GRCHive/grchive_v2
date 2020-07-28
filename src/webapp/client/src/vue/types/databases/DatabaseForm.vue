<template>
    <div>
        <v-text-field
            label="Name"
            filled
            v-model="value.Name"
            :readonly="readonly"
            :rules="[ rules.required ]"
        >
        </v-text-field>

        <v-textarea
            label="Description"
            filled
            v-model="value.Description"
            :readonly="readonly"
            hide-details
            class="mb-4"
        >
        </v-textarea>

        <v-autocomplete
            v-model="value.TypeId"
            label="Database Type"
            filled
            :items="databaseItems"
            :readonly="readonly"
            hide-details
            class="mb-4"
        >
        </v-autocomplete>

        <v-text-field
            label="Other"
            filled
            v-model="value.OtherType"
            :readonly="readonly"
            v-if="value.TypeId == DatabaseType.Other"
        >
        </v-text-field>

        <v-text-field
            label="Version"
            filled
            v-model="value.Version"
            :readonly="readonly"
        >
        </v-text-field>
    </div>
</template>

<script lang="ts">

import Vue from 'vue'
import Component from 'vue-class-component'
import { Prop } from 'vue-property-decorator'
import { RawDatabase, DatabaseType } from '@client/ts/types/databases'
import * as rules from '@client/ts/frontend/formRules'

@Component
export default class DatabaseForm extends Vue {
    readonly rules : any = rules
    readonly DatabaseType = DatabaseType

    @Prop({ required: true })
    value! : RawDatabase

    @Prop({ type: Boolean, default: false })
    readonly! : boolean

    get databaseItems() : any[] {
        return [
            {
                text: 'Other',
                value: DatabaseType.Other,
            },
            {
                text: 'PostgreSQL',
                value: DatabaseType.PostgreSQL,
            },
            {
                text: 'MySQL',
                value: DatabaseType.MySQL,
            },
            {
                text: 'MariaDB',
                value: DatabaseType.MariaDB,
            },
            {
                text: 'SAP HAN',
                value: DatabaseType.SAPHana,
            },
            {
                text: 'IBM DB2',
                value: DatabaseType.IBMDb2,
            },
            {
                text: 'Oracle',
                value: DatabaseType.Oracle,
            },
            {
                text: 'Microsoft SQL',
                value: DatabaseType.MsSql,
            },
            {
                text: 'Microsoft Access',
                value: DatabaseType.MsAccess,
            },
            {
                text: 'SQLite',
                value: DatabaseType.SQLite,
            },
            {
                text: 'H2',
                value: DatabaseType.H2,
            },
        ]
    }
}

</script>
