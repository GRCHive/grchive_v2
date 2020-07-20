<template>
    <div class="d-flex">
        <date-time-picker
            class="flex-grow-1 mr-4"
            label="Start"
            :value="startDate"
            @input="updateStart"
            :enable-time="enableTime"
            :readonly="readonly"
        >
        </date-time-picker>

        <date-time-picker
            class="flex-grow-1"
            label="End"
            :value="endDate"
            @input="updateEnd"
            :enable-time="enableTime"
            :readonly="readonly"
        >
        </date-time-picker>
    </div>
</template>

<script lang="ts">

import Vue from 'vue'
import Component from 'vue-class-component'
import { Prop } from 'vue-property-decorator'
import DateTimePicker from '@client/vue/shared/form/DateTimePicker.vue'

@Component({
    components: {
        DateTimePicker
    }
})
export default class DateRangePicker extends Vue {
    @Prop({ required : true })
    startDate!: string

    @Prop({ required : true })
    endDate!: string

    @Prop({ type: Boolean, default: false})
    enableTime!: boolean

    @Prop({ type: Boolean, default: false})
    readonly!: boolean

    updateStart(v : string) {
        this.$emit('update:startDate', v)
    }

    updateEnd(v : string) {
        this.$emit('update:endDate', v)
    }
}

</script>
