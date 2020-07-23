<template>
    <div>
        <v-text-field
            label="Account ID"
            filled
            v-model="value.AccountId"
            :rules="[ rules.required ]"
            :readonly="readonly"
        >
        </v-text-field>

        <v-text-field
            label="Account Name"
            filled
            v-model="value.Name"
            :rules="[ rules.required ]"
            :readonly="readonly"
        >
        </v-text-field>

        <v-select
            label="Account Type"
            v-model="value.AccountType"
            filled
            hide-details
            :items="accountTypeItems"
            class="mb-4"
            :readonly="readonly"
        >
        </v-select>

        <v-checkbox
            v-model="value.FinanciallyRelevant"
            label="Financially Relevant"
            :readonly="readonly"
        >
        </v-checkbox>

        <single-gl-account-finder
            label="Parent Account"
            v-model="parentAccount"
            :org-id="orgId"
            :engagement-id="engagementId"
            :readonly="readonly"
            hide-details
            class="mb-4"
        >
        </single-gl-account-finder>

        <v-textarea
            label="Description"
            filled
            v-model="value.Description"
            :readonly="readonly"
            hide-details
        >
        </v-textarea>
    </div>
</template>

<script lang="ts">

import Vue from 'vue'
import Component from 'vue-class-component'
import { Prop, Watch } from 'vue-property-decorator'
import { RawGLAccount, GLAccountType } from '@client/ts/types/gl'
import { GrchiveApi } from '@client/ts/main'
import SingleGlAccountFinder from '@client/vue/types/gl/SingleGlAccountFinder.vue'
import * as rules from '@client/ts/frontend/formRules'

@Component({
    components: {
        SingleGlAccountFinder
    }
})
export default class GeneralLedgerAccountForm extends Vue {
    readonly rules : any = rules

    @Prop({ required: true })
    orgId! : number

    @Prop({ required: true })
    engagementId! : number

    @Prop({ required: true })
    value! : RawGLAccount
    parentAccount: RawGLAccount | null = null

    @Prop({ type: Boolean, default: false })
    readonly! : boolean

    @Watch('parentAccount')
    syncParentAccountToValue() {
        if (!!this.parentAccount) {
            this.value.ParentAccountId = this.parentAccount.Id
        } else {
            this.value.ParentAccountId = null
        }
    }

    @Watch('value')
    syncParentAccountFromValue() {
        if (!!this.value.ParentAccountId) {
            GrchiveApi.gl.getAccount(this.orgId, this.engagementId, this.value.ParentAccountId).then((resp : RawGLAccount | null) => {
                this.parentAccount = resp
            })
        } else {
            this.parentAccount = null
        }
    }

    mounted() {
        this.syncParentAccountFromValue()
    }

    get accountTypeItems() : any {
        return [
            {
                text: 'Asset',
                value: GLAccountType.GLATAsset,
            },
            {
                text: 'Liability',
                value: GLAccountType.GLATLiability,
            },
            {
                text: 'Equity',
                value: GLAccountType.GLATEquity,
            },
            {
                text: 'Revenue',
                value: GLAccountType.GLATRevenue,
            },
            {
                text: 'Expense',
                value: GLAccountType.GLATExpense,
            },
            {
                text: 'Contra',
                value: GLAccountType.GLATContra,
            },
        ]
    }
}

</script>
