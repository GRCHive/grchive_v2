<template>
    <generic-save-edit-dialog
        title="Control"
        :op-pending="saveInProgress"
        :edit-mode="editMode"
        :edit-permissions="editPermissions"
        @save-edit="save"
        @cancel-edit="cancel"
    >
        <template v-slot:form="{ canEdit }">
            <control-form
                v-if="!!workingCopy"
                v-model="workingCopy"
                :org-id="parentOrgId"
                :readonly="!canEdit"
            >
            </control-form>
        </template>
    </generic-save-edit-dialog>
</template>

<script lang="ts">

import Vue from 'vue'
import Component from 'vue-class-component'
import { Watch, Prop } from 'vue-property-decorator'
import { RawControl, createEmptyControl } from '@client/ts/types/controls'
import { GrchiveApi } from '@client/ts/main'
import GenericSaveEditDialog from '@client/vue/types/GenericSaveEditDialog.vue'
import ControlForm from '@client/vue/types/controls/ControlForm.vue'
import { Permission } from '@client/ts/types/roles'

@Component({
    components: {
        GenericSaveEditDialog,
        ControlForm,
    }
})
export default class ControlSaveEditDialog extends Vue {
    @Prop()
    value! : RawControl | null

    @Prop({ type: Boolean, default : false})
    editMode!: boolean

    @Prop({ required: true })
    parentEngagementId! : number

    @Prop({ required: true })
    parentOrgId! : number

    workingCopy : RawControl | null = null
    saveInProgress: boolean = false

    get editPermissions() : Permission[] {
        return [Permission.PControlsUpdate]
    }

    @Watch('value')
    syncWorkingCopy() {
        if (!this.value) {
            this.workingCopy = createEmptyControl()
        } else {
            this.workingCopy = JSON.parse(JSON.stringify(this.value))
        }

        if (!!this.workingCopy) {
            this.workingCopy.EngagementId = this.parentEngagementId
        }
    }

    mounted() {
        this.syncWorkingCopy()
    }

    cancel() {
        this.syncWorkingCopy()
        this.$emit('cancel-edit')
    }

    onSuccess(resp : RawControl | null) {
        if (!resp) {
            return
        }

        this.$emit('input', resp)
        this.$emit('save-edit', resp)
        this.syncWorkingCopy()
    }

    save() {
        if (!this.workingCopy) {
            return
        }

        this.saveInProgress = true
        if (!this.editMode) {
            GrchiveApi.controls.createControl(this.parentOrgId, this.workingCopy).then(this.onSuccess).finally(() => {
                this.saveInProgress = false
            })
        } else {
            GrchiveApi.controls.updateControl(this.parentOrgId, this.workingCopy).then(this.onSuccess).finally(() => {
                this.saveInProgress = false
            })
        }
    }
}

</script>
