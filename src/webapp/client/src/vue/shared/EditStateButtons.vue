<template>
    <v-list-item class="px-0">
        <v-list-item-action v-if="canEdit">
            <v-btn
                color="error"
                @click="cancel"
                :loading="editPending"
            >
                Cancel
            </v-btn>
        </v-list-item-action>

        <v-spacer></v-spacer>

        <v-list-item-action v-if="!canEdit">
            <v-btn
                color="primary"
                @click="edit"
                :loading="editPending"
            >
                Edit
            </v-btn>
        </v-list-item-action>

        <v-list-item-action v-if="canEdit">
            <v-btn
                color="success"
                @click="save"
                :loading="editPending"
                :disabled="disabled"
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
    canEdit: boolean = false

    @Prop({ required: true })
    value!: boolean

    @Prop({ default: false })
    editPending! : boolean

    @Prop({ default: false })
    disabled! : boolean

    @Prop({ default: false })
    editMode! : boolean

    changeCanEdit(v : boolean) {
        if (!v && !this.editMode) {
            return
        }
        this.canEdit = v
        this.$emit('input', this.canEdit)
    }

    edit() {
        this.changeCanEdit(true)
    }

    cancel() {
        this.changeCanEdit(false)
        this.$emit('cancel-edit')
    }

    save() {
        this.changeCanEdit(false)
        this.$emit('save-edit')
    }

    mounted() {
        if (!this.editMode) {
            this.changeCanEdit(true)
        }
    }
}

</script>
