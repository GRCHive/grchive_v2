<template>
    <v-card>
        <v-card-title>
            Please Confirm
        </v-card-title>
        <v-divider></v-divider>

        <div class="mx-4">
            <slot></slot>

            <p>
                To confirm this action, please type <b>{{confirmation}}</b> (case sensitive) into the textbox.
            </p>

            <v-text-field
                v-model="checkConfirm"
                filled
                dense
            >
            </v-text-field>
        </div>

        <v-card-actions>
            <v-btn
                color="error"
                @click="cancel"
                :loading="inProgress"
            >
                Cancel
            </v-btn>
            
            <v-spacer></v-spacer>

            <v-btn
                color="success"
                @click="confirm"
                :loading="inProgress"
                :disabled="!allowConfirm"
            >
                Confirm
            </v-btn>
        </v-card-actions>
    </v-card>
</template>

<script lang="ts">

import Vue from 'vue'
import Component from 'vue-class-component'
import { Prop } from 'vue-property-decorator'

@Component
export default class ConfirmationDialog extends Vue {
    @Prop({required : true})
    confirmation! : string

    @Prop({ type: Boolean, default: false })
    inProgress! : boolean

    checkConfirm : string = ''

    get allowConfirm() : boolean {
        return this.checkConfirm.trim() == this.confirmation.trim()
    }

    cancel() {
        this.$emit('do-cancel')
    }

    confirm() {
        this.$emit('do-confirm')
    }
}

</script>
