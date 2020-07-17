<template>
    <v-card>
        <v-card-title>
            {{ editMode ? "Edit" : "New" }} {{ title }}
        </v-card-title>
        <v-divider class="mb-4"></v-divider>

        <v-form onSubmit="return false;" v-model="formValid" class="mx-4">
            <slot name="form" v-bind:canEdit="canEdit">
            </slot>
        </v-form>

        <v-card-actions>
            <edit-state-buttons
                v-model="canEdit"
                :edit-pending="opPending"
                :disabled="!formValid"
                :edit-mode="editMode"
                @save-edit="save"
                @cancel-edit="cancel"
            >
            </edit-state-buttons>
        </v-card-actions>
    </v-card>
</template>

<script lang="ts">

import Vue from 'vue'
import Component from 'vue-class-component'
import { Prop } from 'vue-property-decorator'
import EditStateButtons from '@client/vue/shared/EditStateButtons.vue'

@Component({
    components: {
        EditStateButtons,
    }
})
export default class GenericSaveEditDialog extends Vue {
    @Prop({ required : true })
    title! : string

    @Prop({ required : true })
    opPending!: boolean

    @Prop({ type: Boolean, default: false })
    editMode!: boolean

    canEdit : boolean = false
    formValid: boolean = false

    cancel() {
        this.$emit('cancel-edit')
    }

    save() {
        this.$emit('save-edit')
    }
}

</script>
