<template>
    <generic-save-edit-dialog
        title="Risk"
        :op-pending="saveInProgress"
        :edit-mode="editMode"
        :edit-permissions="editPermissions"
        @save-edit="save"
        @cancel-edit="cancel"
    >
        <template v-slot:form="{ canEdit }">
            <risk-form
                v-if="!!workingCopy"
                v-model="workingCopy"
                :readonly="!canEdit"
            >
            </risk-form>
        </template>
    </generic-save-edit-dialog>
</template>

<script lang="ts">

import Vue from 'vue'
import Component from 'vue-class-component'
import { Watch, Prop } from 'vue-property-decorator'
import { RawRisk, createEmptyRisk } from '@client/ts/types/risks'
import { GrchiveApi } from '@client/ts/main'
import GenericSaveEditDialog from '@client/vue/types/GenericSaveEditDialog.vue'
import RiskForm from '@client/vue/types/risks/RiskForm.vue'
import { Permission } from '@client/ts/types/roles'

@Component({
    components: {
        GenericSaveEditDialog,
        RiskForm,
    }
})
export default class RiskSaveEditDialog extends Vue {
    @Prop()
    value! : RawRisk | null

    @Prop({ type: Boolean, default : false})
    editMode!: boolean

    @Prop({ required: true })
    parentEngagementId! : number

    @Prop({ required: true })
    parentOrgId! : number

    workingCopy : RawRisk | null = null
    saveInProgress: boolean = false

    get editPermissions() : Permission[] {
        return [Permission.PRisksUpdate]
    }

    @Watch('value')
    syncWorkingCopy() {
        if (!this.value) {
            this.workingCopy = createEmptyRisk()
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

    onSuccess(resp : RawRisk) {
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
            GrchiveApi.risks.createRisk(this.parentOrgId, this.workingCopy!).then(this.onSuccess).finally(() => {
                this.saveInProgress = false
            })
        } else {
            GrchiveApi.risks.updateRisk(this.parentOrgId, this.workingCopy!).then(this.onSuccess).finally(() => {
                this.saveInProgress = false
            })
        }
    }
}

</script>
