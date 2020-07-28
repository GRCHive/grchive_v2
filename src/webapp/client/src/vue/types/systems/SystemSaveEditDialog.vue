<template>
    <generic-save-edit-dialog
        title="System"
        :op-pending="saveInProgress"
        :edit-mode="editMode"
        :edit-permissions="editPermissions"
        @save-edit="save"
        @cancel-edit="cancel"
    >
        <template v-slot:form="{ canEdit }">
            <system-form
                v-if="!!workingCopy"
                v-model="workingCopy"
                :readonly="!canEdit"
            >
            </system-form>
        </template>
    </generic-save-edit-dialog>
</template>

<script lang="ts">

import Vue from 'vue'
import Component from 'vue-class-component'
import { Watch, Prop } from 'vue-property-decorator'
import { RawSystem, createEmptySystem } from '@client/ts/types/systems'
import { GrchiveApi, ErrorHandler } from '@client/ts/main'
import GenericSaveEditDialog from '@client/vue/types/GenericSaveEditDialog.vue'
import SystemForm from '@client/vue/types/systems/SystemForm.vue'
import { Permission } from '@client/ts/types/roles'

@Component({
    components: {
        GenericSaveEditDialog,
        SystemForm,
    }
})
export default class SystemSaveEditDialog extends Vue {
    @Prop()
    value! : RawSystem | null

    @Prop({ type: Boolean, default : false})
    editMode!: boolean

    @Prop({ required: true })
    parentEngagementId! : number

    @Prop({ required: true })
    parentOrgId! : number

    workingCopy : RawSystem | null = null
    saveInProgress: boolean = false

    get editPermissions() : Permission[] {
        return [Permission.PSystemsUpdate]
    }

    @Watch('value')
    syncWorkingCopy() {
        if (!this.value) {
            this.workingCopy = createEmptySystem()
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

    onSuccess(resp : RawSystem) {
        this.$emit('input', resp)
        this.$emit('save-edit', resp)
        this.syncWorkingCopy()
    }

    onError(err : any) {
        ErrorHandler.failurePopupOnError(err, {
            context: 'Failed to create/edit system.'
        })
    }

    save() {
        if (!this.workingCopy) {
            return
        }

        this.saveInProgress = true
        if (!this.editMode) {
            GrchiveApi.systems.createSystem(this.parentOrgId, this.workingCopy!).then(this.onSuccess).catch(this.onError).finally(() => {
                this.saveInProgress = false
            })
        } else {
            GrchiveApi.systems.updateSystem(this.parentOrgId, this.workingCopy!).then(this.onSuccess).catch(this.onError).finally(() => {
                this.saveInProgress = false
            })
        }
    }
}

</script>
