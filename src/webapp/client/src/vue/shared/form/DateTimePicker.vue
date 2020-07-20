<template>
    <div class="flatpickr" ref="dateTime">
        <v-text-field
            :label="label"
            filled
            data-input
        >
        </v-text-field>
    </div>
</template>

<script lang="ts">

import Vue from 'vue'
import Component from 'vue-class-component'
import { Prop } from 'vue-property-decorator'
import { v4 as uuidv4 } from 'uuid'
import formatISO from 'date-fns/formatISO'
import flatpickr from 'flatpickr'
import 'flatpickr/dist/themes/material_blue.css'

@Component
export default class DateTimePicker extends Vue {
    @Prop({ required : true })
    value!: string

    @Prop({ default: 'Date Time'})
    label!: string

    @Prop({ type: Boolean, default: false})
    enableTime!: boolean

    fpInstance: any

    $refs!: {
        dateTime: HTMLElement
    }

    updateValue() {
        if (this.fpInstance.selectedDates.length == 0) {
            this.$emit('input', '')
        } else {
            this.$emit('input', formatISO(this.fpInstance.selectedDates[0]))
        }
    }

    mounted() {
        this.fpInstance = flatpickr(this.$refs.dateTime, {
            dateFormat: this.enableTime ? 
                'F J, Y at G:i:S K' :
                'F J, Y',
            wrap: true,
            enableTime: this.enableTime,
        })
        this.fpInstance.config.onValueUpdate.push(this.updateValue)
    }

}

</script>
