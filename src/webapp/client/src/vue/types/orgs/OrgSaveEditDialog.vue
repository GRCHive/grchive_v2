<template>
    <generic-save-edit-dialog
        title="Organization"
        :op-pending="saveInProgress"
        :edit-mode="editMode"
        @save-edit="save"
        @cancel-edit="cancel"
        :edit-permissions="editPermissions"
    >
        <template v-slot:form="{ canEdit }">
            <org-form
                v-if="!!workingCopy"
                v-model="workingCopy"
                :readonly="!canEdit"
            >
            </org-form>
        </template>
    </generic-save-edit-dialog>
</template>

<script lang="ts">

import Vue from 'vue'
import Component from 'vue-class-component'
import { Watch, Prop } from 'vue-property-decorator'
import { RawOrganization, createEmptyOrg } from '@client/ts/types/orgs'
import { GrchiveApi, ErrorHandler } from '@client/ts/main'
import GenericSaveEditDialog from '@client/vue/types/GenericSaveEditDialog.vue'
import OrgForm from '@client/vue/types/orgs/OrgForm.vue'
import { Permission } from '@client/ts/types/roles'

@Component({
    components: {
        GenericSaveEditDialog,
        OrgForm,
    }
})
export default class OrgSaveEditDialog extends Vue {
    @Prop()
    value! : RawOrganization | null

    @Prop({ type: Boolean, default : false})
    editMode!: boolean

    @Prop({ default: -1 })
    parentOrgId! : number

    workingCopy : RawOrganization | null = null
    saveInProgress: boolean = false

    get editPermissions() : Permission[] {
        return [Permission.POrgProfileUpdate]
    }

    @Watch('value')
    syncWorkingCopy() {
        if (!this.value) {
            this.workingCopy = createEmptyOrg()
        } else {
            this.workingCopy = JSON.parse(JSON.stringify(this.value))
        }
    }

    mounted() {
        this.syncWorkingCopy()
    }

    cancel() {
        this.syncWorkingCopy()
        this.$emit('cancel-edit')
    }

    onSuccess(resp : RawOrganization) {
        this.$emit('input', resp)
        this.$emit('save-edit', resp)
        this.syncWorkingCopy()
    }

    onError(err : any) {
        ErrorHandler.failurePopupOnError(err, {
            context: 'Failed to save/edit the organization profile.'
        })
    }

    save() {
        if (!this.workingCopy) {
            return
        }

        this.saveInProgress = true

        if (this.editMode) {
            GrchiveApi.orgs.updateOrg(this.workingCopy!).then(this.onSuccess).catch(this.onError).finally(() => {
                this.saveInProgress = false
            })
        } else {
            if (this.parentOrgId == -1) {
                GrchiveApi.orgs.createOrg(this.workingCopy!).then(this.onSuccess).catch(this.onError).finally(() => {
                    this.saveInProgress = false
                })
            } else {
                GrchiveApi.orgs.createSuborg(this.parentOrgId, this.workingCopy!).then(this.onSuccess).catch(this.onError).finally(() => {
                    this.saveInProgress = false
                })
            }
        }
    }
}

</script>
