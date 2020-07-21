<template>
    <div>
        <v-text-field
            :label="label"
            :value="valueStr"
            filled
            readonly
            v-if="readonly"
        >
        </v-text-field>

        <div class="flatpickr" ref="dateTime" v-else>
            <v-text-field
                :label="label"
                :value="valueStr"
                filled
                data-input
            >
            </v-text-field>
        </div>
    </div>
</template>

<script lang="ts">

import Vue from 'vue'
import Component from 'vue-class-component'
import { Watch, Prop } from 'vue-property-decorator'
import { v4 as uuidv4 } from 'uuid'
import formatISO from 'date-fns/formatISO'
import parseISO from 'date-fns/parseISO'
import flatpickr from 'flatpickr'
import 'flatpickr/dist/themes/material_blue.css'
import { standardFormatTime } from '@client/ts/time'

@Component
export default class DateTimePicker extends Vue {
    @Prop({ required : true })
    value!: string

    @Prop({ default: 'Date Time'})
    label!: string

    @Prop({ type: Boolean, default: false})
    enableTime!: boolean

    @Prop({ type: Boolean, default: false})
    readonly!: boolean

    fpInstance: any
    valueStr : string = ''

    $refs!: {
        dateTime: HTMLElement
    }

    @Watch('value')
    syncValue() {
        if (!!this.fpInstance) {
            this.fpInstance.setDate(parseISO(this.value))
        }
        this.updateValue(true)
    }

    updateValue(quiet? : boolean) {
        // realStr gets passed up the chain via the value prop (v-model)
        // while valueStr is presented to the user
        let realStr : string = ''

        if (!!this.fpInstance) {
            if (this.fpInstance.selectedDates.length == 0) {
                this.valueStr = ''
                realStr = ''
            } else {
                this.valueStr = standardFormatTime(this.fpInstance.selectedDates[0])
                realStr = formatISO(this.fpInstance.selectedDates[0])
            }
        } else {
            this.valueStr = standardFormatTime(parseISO(this.value))
        }

        if (!quiet) {
            this.$emit('input', realStr)
        }
    }

    @Watch('readonly')
    updateFpInstance() {
        if (!this.readonly) {
            Vue.nextTick(() => {
                this.fpInstance = flatpickr(this.$refs.dateTime, {
                    dateFormat: this.enableTime ? 
                        'F J, Y, h:i K' :
                        'F J, Y',
                    wrap: true,
                    enableTime: this.enableTime,
                })
                // Don't pass this.updateValue directly otherwise the quiet variable will be initialized to something that's not false.
                this.fpInstance.config.onValueUpdate.push(() => { this.updateValue() })
                this.syncValue()
            })
        } else {
            if (!!this.fpInstance) {
                this.fpInstance.destroy()
                this.fpInstance = null
            }
            this.syncValue()
        }
    }

    mounted() {
        this.updateFpInstance()
    }

}

</script>
