<template>
    <v-list-item class="px-0">
        <v-list-item-action v-if="canEdit">
            <v-btn
                color="error"
                @click="cancel"
                :loading="editPending"
            >
                Cancel
            </v-btn>
        </v-list-item-action>

        <v-spacer></v-spacer>

        <v-list-item-action v-if="!canEdit">
            <restrict-role-permission-button
                color="primary"
                @click="edit"
                :loading="editPending"
                :permissions="editPermissions"
                :org-id="orgId"
                :engagement-id="engagementId"
            >
                Edit
            </restrict-role-permission-button>
        </v-list-item-action>

        <v-list-item-action v-if="canEdit">
            <v-btn
                color="success"
                @click="save"
                :loading="editPending"
                :disabled="disabled"
            >
                Save
            </v-btn>
        </v-list-item-action>
    </v-list-item>
</template>

<script lang="ts">

import Vue from 'vue'
import Component from 'vue-class-component'
import { Prop } from 'vue-property-decorator'
import { Permission } from '@client/ts/types/roles'
import RestrictRolePermissionButton from '@client/vue/loading/RestrictRolePermissionButton.vue'

@Component({
    components: {
        RestrictRolePermissionButton
    }
})
export default class EditStateButtons extends Vue {
    canEdit: boolean = false

    @Prop({ required: true })
    value!: boolean

    @Prop({ type: Boolean, default: false })
    editPending! : boolean

    @Prop({ type: Boolean, default: false })
    disabled! : boolean

    @Prop({ type: Boolean, default: false })
    editMode! : boolean

    @Prop({ type: Array, default: []})
    editPermissions!: Permission[]

    get orgId() : number {
        const org = this.$store.state.org.rawOrg
        if (!org) {
            return -1
        }
        return org.Id
    }

    get engagementId() : number {
        const eng = this.$store.state.engagements.rawEngagement
        if (!eng) {
            return -1
        }
        return eng.Id
    }

    changeCanEdit(v : boolean) {
        if (!v && !this.editMode) {
            return
        }
        this.canEdit = v
        this.$emit('input', this.canEdit)
    }

    edit() {
        this.changeCanEdit(true)
    }

    cancel() {
        this.changeCanEdit(false)
        this.$emit('cancel-edit')
    }

    save() {
        this.changeCanEdit(false)
        this.$emit('save-edit')
    }

    mounted() {
        if (!this.editMode) {
            this.changeCanEdit(true)
        }
    }
}

</script>
