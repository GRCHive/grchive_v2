<template>
    <generic-save-edit-dialog
        title="General Ledger Account"
        :op-pending="saveInProgress"
        :edit-mode="editMode"
        :edit-permissions="editPermissions"
        @save-edit="save"
        @cancel-edit="cancel"
    >
        <template v-slot:form="{ canEdit }">
            <general-ledger-account-form
                v-if="!!workingCopy"
                v-model="workingCopy"
                :readonly="!canEdit"
                :org-id="parentOrgId"
                :engagement-id="parentEngagementId"
            >
            </general-ledger-account-form>
        </template>
    </generic-save-edit-dialog>
</template>

<script lang="ts">

import Vue from 'vue'
import Component from 'vue-class-component'
import { Watch, Prop } from 'vue-property-decorator'
import { RawGLAccount, createEmptyGLAccount } from '@client/ts/types/gl'
import { GrchiveApi } from '@client/ts/main'
import GenericSaveEditDialog from '@client/vue/types/GenericSaveEditDialog.vue'
import GeneralLedgerAccountForm from '@client/vue/types/gl/GeneralLedgerAccountForm.vue'
import { Permission } from '@client/ts/types/roles'

@Component({
    components: {
        GenericSaveEditDialog,
        GeneralLedgerAccountForm,
    }
})
export default class GeneralLedgerAccountSaveEditDialog extends Vue {
    @Prop()
    value! : RawGLAccount | null

    @Prop({ type: Boolean, default : false})
    editMode!: boolean

    @Prop({ required: true })
    parentEngagementId! : number

    @Prop({ required: true })
    parentOrgId! : number

    workingCopy : RawGLAccount | null = null
    saveInProgress: boolean = false

    get editPermissions() : Permission[] {
        return [Permission.PGLUpdate]
    }

    @Watch('value')
    syncWorkingCopy() {
        if (!this.value) {
            this.workingCopy = createEmptyGLAccount()
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

    onSuccess(resp : RawGLAccount | null) {
        console.log("on success: ", resp)
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
            GrchiveApi.gl.createAccount(this.parentOrgId, this.workingCopy).then(this.onSuccess).finally(() => {
                this.saveInProgress = false
            })
        } else {
            GrchiveApi.gl.updateAccount(this.parentOrgId, this.workingCopy).then(this.onSuccess).finally(() => {
                this.saveInProgress = false
            })
        }
    }
}

</script>
