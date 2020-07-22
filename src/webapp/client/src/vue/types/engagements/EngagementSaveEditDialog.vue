<template>
    <generic-save-edit-dialog
        title="Engagement"
        :op-pending="saveInProgress"
        :edit-mode="editMode"
        :edit-permissions="editPermissions"
        @save-edit="save"
        @cancel-edit="cancel"
    >
        <template v-slot:form="{ canEdit }">
            <engagement-form
                v-if="!!workingCopy"
                v-model="workingCopy"
                :readonly="!canEdit"
                :disable-role-edit="editMode"
            >
            </engagement-form>
        </template>
    </generic-save-edit-dialog>
</template>

<script lang="ts">

import Vue from 'vue'
import Component from 'vue-class-component'
import { Watch, Prop } from 'vue-property-decorator'
import { RawEngagement, createEmptyEngagement, cleanRawEngagementFromJSON } from '@client/ts/types/engagements'
import { GrchiveApi } from '@client/ts/main'
import GenericSaveEditDialog from '@client/vue/types/GenericSaveEditDialog.vue'
import EngagementForm from '@client/vue/types/engagements/EngagementForm.vue'
import { Permission } from '@client/ts/types/roles'

@Component({
    components: {
        GenericSaveEditDialog,
        EngagementForm,
    }
})
export default class EngagementSaveEditDialog extends Vue {
    @Prop()
    value! : RawEngagement | null

    @Prop({ type: Boolean, default : false})
    editMode!: boolean

    @Prop({ default: -1 })
    parentOrgId! : number

    workingCopy : RawEngagement | null = null
    saveInProgress: boolean = false

    get editPermissions() : Permission[] {
        return [Permission.POrgEngagementUpdate]
    }

    @Watch('value')
    syncWorkingCopy() {
        if (!this.value) {
            this.workingCopy = createEmptyEngagement()
        } else {
            this.workingCopy = JSON.parse(JSON.stringify(this.value))
        }

        if (!!this.workingCopy) {
            cleanRawEngagementFromJSON(this.workingCopy)
            this.workingCopy.OrgId = this.parentOrgId
        }
    }

    mounted() {
        this.syncWorkingCopy()
    }

    cancel() {
        this.syncWorkingCopy()
        this.$emit('cancel-edit')
    }

    onSuccess(resp : RawEngagement | null) {
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
            GrchiveApi.engagements.createEngagement(this.workingCopy).then(this.onSuccess).finally(() => {
                this.saveInProgress = false
            })
        } else {
            GrchiveApi.engagements.updateEngagement(this.workingCopy).then(this.onSuccess).finally(() => {
                this.saveInProgress = false
            })
        }
    }
}

</script>
