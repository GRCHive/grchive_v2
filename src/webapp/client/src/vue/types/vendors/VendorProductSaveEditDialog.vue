<template>
    <generic-save-edit-dialog
        title="Vendor Product"
        :op-pending="saveInProgress"
        :edit-mode="editMode"
        :edit-permissions="editPermissions"
        @save-edit="save"
        @cancel-edit="cancel"
    >
        <template v-slot:form="{ canEdit }">
            <vendor-product-form
                v-if="!!workingCopy"
                v-model="workingCopy"
                :readonly="!canEdit"
            >
            </vendor-product-form>
        </template>
    </generic-save-edit-dialog>
</template>

<script lang="ts">

import Vue from 'vue'
import Component from 'vue-class-component'
import { Watch, Prop } from 'vue-property-decorator'
import { RawVendorProduct, createEmptyVendorProduct } from '@client/ts/types/vendors'
import { GrchiveApi, ErrorHandler } from '@client/ts/main'
import GenericSaveEditDialog from '@client/vue/types/GenericSaveEditDialog.vue'
import VendorProductForm from '@client/vue/types/vendors/VendorProductForm.vue'
import { Permission } from '@client/ts/types/roles'

@Component({
    components: {
        GenericSaveEditDialog,
        VendorProductForm,
    }
})
export default class VendorProductSaveEditDialog extends Vue {
    @Prop()
    value! : RawVendorProduct | null

    @Prop({ type: Boolean, default : false})
    editMode!: boolean

    @Prop({ required: true })
    parentEngagementId! : number

    @Prop({ required: true })
    parentOrgId! : number

    @Prop({ required: true })
    parentVendorId! : number

    workingCopy : RawVendorProduct | null = null
    saveInProgress: boolean = false

    get editPermissions() : Permission[] {
        return [Permission.PVendorProductsUpdate]
    }

    @Watch('value')
    syncWorkingCopy() {
        if (!this.value) {
            this.workingCopy = createEmptyVendorProduct()
        } else {
            this.workingCopy = JSON.parse(JSON.stringify(this.value))
        }

        if (!!this.workingCopy) {
            this.workingCopy.VendorId = this.parentVendorId
        }
    }

    mounted() {
        this.syncWorkingCopy()
    }

    cancel() {
        this.syncWorkingCopy()
        this.$emit('cancel-edit')
    }

    onSuccess(resp : RawVendorProduct) {
        this.$emit('input', resp)
        this.$emit('save-edit', resp)
        this.syncWorkingCopy()
    }

    onError(err : any) {
        ErrorHandler.failurePopupOnError(err, {
            context: 'Failed to create/edit vendor product.'
        })
    }

    save() {
        if (!this.workingCopy) {
            return
        }

        this.saveInProgress = true
        if (!this.editMode) {
            GrchiveApi.vendors.createVendorProduct(this.parentOrgId, this.parentEngagementId, this.workingCopy).then(this.onSuccess).catch(this.onError).finally(() => {
                this.saveInProgress = false
            })
        } else {
            GrchiveApi.vendors.updateVendorProduct(this.parentOrgId, this.parentEngagementId, this.workingCopy).then(this.onSuccess).catch(this.onError).finally(() => {
                this.saveInProgress = false
            })
        }
    }
}

</script>
