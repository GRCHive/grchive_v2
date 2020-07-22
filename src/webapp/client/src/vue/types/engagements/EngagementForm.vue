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

        <date-range-picker
            :start-date="startTimeStr"
            :end-date="endTimeStr"
            :readonly="readonly"
            @update:startDate="updateStartTime"
            @update:endDate="updateEndTime"
            enable-time
        >
        </date-range-picker>

        <multi-role-finder
            v-if="!disableRoleEdit"
            label="Assigned Roles"
            v-model="value.Roles"
            :rules="[rules.required]"
            :readonly="readonly"
            :org-id="value.OrgId"
        >
        </multi-role-finder>
    </div>
</template>

<script lang="ts">

import Vue from 'vue'
import Component from 'vue-class-component'
import { Prop } from 'vue-property-decorator'
import { RawEngagement } from '@client/ts/types/engagements'
import DateRangePicker from '@client/vue/shared/form/DateRangePicker.vue'
import MultiRoleFinder from '@client/vue/types/roles/MultiRoleFinder.vue'
import * as rules from '@client/ts/frontend/formRules'
import formatISO from 'date-fns/formatISO'
import parseISO from 'date-fns/parseISO'

@Component({
    components: {
        DateRangePicker,
        MultiRoleFinder,
    }
})
export default class EngagementForm extends Vue {
    readonly rules : any = rules

    @Prop({ required: true })
    value! : RawEngagement

    @Prop({ type: Boolean, default: false })
    readonly! : boolean

    @Prop({ type: Boolean, default: false })
    disableRoleEdit! : boolean

    get startTimeStr() : string {
        if (!this.value.StartTime) {
            return ''
        }
        return formatISO(this.value.StartTime)
    }
    
    get endTimeStr() : string {
        if (!this.value.EndTime) {
            return ''
        }
        return formatISO(this.value.EndTime)
    }

    updateStartTime(v : string) {
        this.value.StartTime = new Date(parseISO(v))
        this.$emit('input', this.value)
    }

    updateEndTime(v : string) {
        this.value.EndTime = new Date(parseISO(v))
        this.$emit('input', this.value)
    }
}

</script>
