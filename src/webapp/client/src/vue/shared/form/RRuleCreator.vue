<template>
    <div>
        <v-select
            v-model="freq"
            :items="freqOptions"
            filled
            hide-details
            label="Frequency"
            :readonly="readonly"
            class="mb-4"
        >
        </v-select>

        <v-text-field
            label="Interval"
            outlined
            type="number"
            v-model="interval"
            :readonly="readonly"
            :rules = "[ rules.numeric, rules.geq(0) ]"
            class="mb-4"
        >
        </v-text-field>

        <date-time-picker
            :value="dtStartStr"
            @input="onChangeDtStart"
            label="Starting From"
            enable-time
            :readonly="readonly"
        >
        </date-time-picker>
    </div>
</template>

<script lang="ts">

import Vue from 'vue'
import Component from 'vue-class-component'
import { Watch, Prop } from 'vue-property-decorator'
import { RRule } from 'rrule'
import { zonedTimeToUtc } from 'date-fns-tz'
import DateTimePicker from '@client/vue/shared/form/DateTimePicker.vue'
import formatISO from 'date-fns/formatISO'
import parseISO from 'date-fns/parseISO'
import * as rules from '@client/ts/frontend/formRules'

@Component({
    components: {
        DateTimePicker,
    }
})
export default class RRuleCreator extends Vue {
    readonly rules: any = rules

    @Prop()
    value! : string

    @Prop({type: Boolean})
    readonly!: boolean

    freq : any = RRule.DAILY
    interval : number = 1
    dtStart : Date = new Date()

    get dtStartStr() : string {
        return formatISO(this.dtStart)        
    }

    onChangeDtStart(v : string) {
        this.dtStart = parseISO(v)
    }

    get freqOptions() : any[] {
        return [
            {
                text: 'Daily',
                value: RRule.DAILY,
            },
            {
                text: 'Weekly',
                value: RRule.WEEKLY,
            },
            {
                text: 'Monthly',
                value: RRule.MONTHLY,
            },
            {
                text: 'Yearly',
                value: RRule.YEARLY,
            },
        ]
    }

    @Watch('value')
    syncFromValue() {
        this.freq = RRule.DAILY
        this.interval = 1
        this.dtStart = new Date()

        if (this.value != '') {
            const rule = RRule.fromString(this.value)
            this.freq = rule.origOptions.freq

            if (!!rule.origOptions.interval) {
                this.interval = rule.origOptions.interval
            }

            if (!!rule.origOptions.dtstart) {
                this.dtStart = rule.origOptions.dtstart
            }
        }
    }

    get workingRRule() : RRule {
        // This conversion might not be necessary since the DateTimePicker seems to set the UTC time properly?
        let utcTime : Date = zonedTimeToUtc(this.dtStart, Intl.DateTimeFormat().resolvedOptions().timeZone)
        return new RRule({
            freq: this.freq,
            interval: this.interval,
            dtstart: utcTime,
        })
    }

    @Watch('workingRRule')
    syncToValue() {
        this.$emit('input', this.workingRRule.toString())
    }

    mounted() {
        this.syncFromValue()
    }
}

</script>
