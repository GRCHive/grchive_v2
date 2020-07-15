<template>
    <v-list-item class="px-0">
        <v-list-item-action v-if="value">
            <v-btn
                color="error"
                @click="cancel"
                :loading="editPending"
            >
                Cancel
            </v-btn>
        </v-list-item-action>

        <v-spacer></v-spacer>

        <v-list-item-action v-if="!value">
            <v-btn
                color="primary"
                @click="edit"
                :loading="editPending"
            >
                Edit
            </v-btn>
        </v-list-item-action>

        <v-list-item-action v-if="value">
            <v-btn
                color="success"
                @click="save"
                :loading="editPending"
            >
                Save
            </v-btn>
        </v-list-item-action>
    </v-list-item>
</template>

<script lang="ts">

import Vue from 'vue'
import Component from 'vue-class-component'
import { Prop } from 'vue-property-decorator'

@Component
export default class EditStateButtons extends Vue {
    @Prop({ required: true })
    value!: boolean

    @Prop({ default: false })
    editPending! : boolean

    edit() {
        this.value = true
        this.$emit('input', this.value)
    }

    cancel() {
        this.value = false
        this.$emit('input', this.value)
        this.$emit('cancel-edit')
    }

    save() {
        this.value = false
        this.$emit('input', this.value)
        this.$emit('save-edit')
    }
}

</script>
