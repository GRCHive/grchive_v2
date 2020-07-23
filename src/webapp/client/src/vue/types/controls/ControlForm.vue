<template>
    <div>
        <v-text-field
            label="Identifier"
            filled
            v-model="value.HumanId"
            :readonly="readonly"
            :rules="[ rules.required, rules.createMaxLength(256) ]"
        >
        </v-text-field>

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

        <v-select
            filled
            label="Control Type"
            v-model="value.ControlType"
            :items="controlTypeItems"
            :readonly="readonly"
            hide-details
            class="mb-4"
        >
        </v-select>

        <single-user-finder
            v-model="workingUser"
            filled
            label="Control Owner"
            :readonly="readonly"
            hide-details
            class="mb-4"
            :loading="loadingWorkingUser"
            :org-id="orgId"
        >
        </single-user-finder>

        <v-checkbox
            label="Is Manual"
            v-model="value.IsManual"
            :readonly="readonly"
            hide-details
            class="mb-4"
        >
        </v-checkbox>

        <v-select
            filled
            label="Frequency Type"
            v-model="value.FreqType"
            :items="controlFreqTypeItems"
            :readonly="readonly"
            class="mb-4"
            hide-details
        >
        </v-select>
    
        <v-text-field
            v-if="value.FreqType == ControlFrequencyType.CFTOther"
            label="Other Frequency"
            filled
            v-model="value.FreqOther"
            :readonly="readonly"
            hide-details
            class="mb-4"
        >
        </v-text-field>
            
        <r-rule-creator
            v-if="value.FreqType == ControlFrequencyType.CFTInterval"
            v-model="value.FreqInterval"
            :readonly="readonly"
        >
        </r-rule-creator>

        <div class="d-flex align-center my-4">
            <span class="font-weight-bold">Likelihood of Failure:</span>

            <v-rating
                v-model="value.Likelihood"
                :readonly="readonly"
                color="error"
                size="30"
            >
            </v-rating>
        </div>

        <v-textarea
            label="Likelihood Justification"
            filled
            v-model="value.LikelihoodJustification"
            :readonly="readonly"
            hide-details
        >
        </v-textarea>
    </div>
</template>

<script lang="ts">

import Vue from 'vue'
import Component from 'vue-class-component'
import { Watch, Prop } from 'vue-property-decorator'
import { RawControl, ControlType, ControlFrequencyType } from '@client/ts/types/controls'
import { RawUser } from '@client/ts/types/users'
import { GrchiveApi } from '@client/ts/main'
import RRuleCreator from '@client/vue/shared/form/RRuleCreator.vue'
import * as rules from '@client/ts/frontend/formRules'
import SingleUserFinder from '@client/vue/types/users/SingleUserFinder.vue'

@Component({
    components: {
        RRuleCreator,
        SingleUserFinder,
    }
})
export default class ControlForm extends Vue {
    readonly rules : any = rules
    readonly ControlFrequencyType : any = ControlFrequencyType

    @Prop({ required: true })
    value! : RawControl
    workingUser : RawUser | null = null
    loadingWorkingUser: boolean = false

    @Prop({ type: Boolean, default: false })
    readonly! : boolean

    @Prop({  required: true })
    orgId! : number

    @Watch('value')
    syncToWorkingUser() {
        if (!this.value.OwnerId) {
            this.workingUser = null
        } else {
            this.loadingWorkingUser = true
            GrchiveApi.orgs.getOrgUser(this.orgId, this.value.OwnerId).then((resp : RawUser) => {
                this.workingUser = resp
            }).finally(() => {
                this.loadingWorkingUser = false
            })
        }
    }

    @Watch('workingUser')
    syncFromWorkingUser() {
        if (!this.workingUser) {
            this.value.OwnerId = null
        } else {
            this.value.OwnerId = this.workingUser.Id
        }
    }

    get controlTypeItems() : any[] {
        return [
            {
                text: 'Access',
                value: ControlType.CTAccess,
            },
            {
                text: 'Authorization',
                value: ControlType.CTAuthorization,
            },
            {
                text: 'Configuration',
                value: ControlType.CTConfiguration,
            },
            {
                text: 'GITC',
                value: ControlType.CTGitc,
            },
            {
                text: 'Interface',
                value: ControlType.CTInterface,
            },
            {
                text: 'Management Review',
                value: ControlType.CTManagementReview,
            },
            {
                text: 'Reconciliation',
                value: ControlType.CTReconciliation,
            },
            {
                text: 'Report/IPE',
                value: ControlType.CTReport,
            },
        ]
    }

    get controlFreqTypeItems() : any[] {
        return [
            {
                text: 'Ad-Hoc',
                value: ControlFrequencyType.CFTAdHoc,
            },
            {
                text: 'Interval',
                value: ControlFrequencyType.CFTInterval,
            },
            {
                text: 'Other',
                value: ControlFrequencyType.CFTOther,
            },
        ]
    }
}

</script>
