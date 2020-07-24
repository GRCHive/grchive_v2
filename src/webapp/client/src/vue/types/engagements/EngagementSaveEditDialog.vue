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

            <single-engagement-finder
                v-model="baseEngagement"
                label="Roll-Forward From"
                v-if="!editMode"
                :clearable="canEdit"
                :org-id="parentOrgId"
            >
            </single-engagement-finder>
        </template>
    </generic-save-edit-dialog>
</template>

<script lang="ts">

import Vue from 'vue'
import Component from 'vue-class-component'
import { Watch, Prop } from 'vue-property-decorator'
import { RawEngagement, createEmptyEngagement, cleanRawEngagementFromJSON } from '@client/ts/types/engagements'
import { GrchiveApi, ErrorHandler } from '@client/ts/main'
import GenericSaveEditDialog from '@client/vue/types/GenericSaveEditDialog.vue'
import EngagementForm from '@client/vue/types/engagements/EngagementForm.vue'
import SingleEngagementFinder from '@client/vue/types/engagements/SingleEngagementFinder.vue'
import { Permission } from '@client/ts/types/roles'

@Component({
    components: {
        GenericSaveEditDialog,
        EngagementForm,
        SingleEngagementFinder,
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
    baseEngagement : RawEngagement | null = null
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

    onSuccess(resp : RawEngagement) {
        this.$emit('input', resp)
        this.$emit('save-edit', resp)
        this.syncWorkingCopy()
    }

    onError(err : any) {
        ErrorHandler.failurePopupOnError(err, {
            context: 'Failed to save/edit engagement.'
        })
    }

    save() {
        if (!this.workingCopy) {
            return
        }

        this.saveInProgress = true
        if (!this.editMode) {
            GrchiveApi.engagements.createEngagement(this.workingCopy, this.baseEngagement).then(this.onSuccess).catch(this.onError).finally(() => {
                this.saveInProgress = false
            })
        } else {
            GrchiveApi.engagements.updateEngagement(this.workingCopy).then(this.onSuccess).catch(this.onError).finally(() => {
                this.saveInProgress = false
            })
        }
    }
}

</script>
