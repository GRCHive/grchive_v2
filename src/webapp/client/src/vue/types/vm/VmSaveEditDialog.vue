<template>
    <generic-save-edit-dialog
        title="Vm"
        :op-pending="saveInProgress"
        :edit-mode="editMode"
        :edit-permissions="editPermissions"
        @save-edit="save"
        @cancel-edit="cancel"
    >
        <template v-slot:form="{ canEdit }">
            <vm-form
                v-if="!!workingCopy"
                v-model="workingCopy"
                :readonly="!canEdit"
            >
            </vm-form>
        </template>
    </generic-save-edit-dialog>
</template>

<script lang="ts">

import Vue from 'vue'
import Component from 'vue-class-component'
import { Watch, Prop } from 'vue-property-decorator'
import { VirtualMachine, createEmptyVm } from '@client/ts/types/vm'
import { GrchiveApi, ErrorHandler } from '@client/ts/main'
import GenericSaveEditDialog from '@client/vue/types/GenericSaveEditDialog.vue'
import VmForm from '@client/vue/types/vm/VmForm.vue'
import { Permission } from '@client/ts/types/roles'

@Component({
    components: {
        GenericSaveEditDialog,
        VmForm,
    }
})
export default class VmSaveEditDialog extends Vue {
    @Prop()
    value! : VirtualMachine | null

    @Prop({ type: Boolean, default : false})
    editMode!: boolean

    @Prop({ required: true })
    parentEngagementId! : number

    @Prop({ required: true })
    parentOrgId! : number

    workingCopy : VirtualMachine | null = null
    saveInProgress: boolean = false

    get editPermissions() : Permission[] {
        return [Permission.PVMUpdate]
    }

    @Watch('value')
    syncWorkingCopy() {
        if (!this.value) {
            this.workingCopy = createEmptyVirtualMachine()
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

    onSuccess(resp : VirtualMachine) {
        this.$emit('input', resp)
        this.$emit('save-edit', resp)
        this.syncWorkingCopy()
    }

    onError(err : any) {
        ErrorHandler.failurePopupOnError(err, {
            context: 'Failed to create/edit vm.'
        })
    }

    save() {
        if (!this.workingCopy) {
            return
        }

        this.saveInProgress = true
        if (!this.editMode) {
            GrchiveApi.vm.createVm(this.parentOrgId, this.workingCopy!).then(this.onSuccess).catch(this.onError).finally(() => {
                this.saveInProgress = false
            })
        } else {
            GrchiveApi.vm.updateVm(this.parentOrgId, this.workingCopy!).then(this.onSuccess).catch(this.onError).finally(() => {
                this.saveInProgress = false
            })
        }
    }
}

</script>
