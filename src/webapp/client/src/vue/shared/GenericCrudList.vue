<template>
    <div>
        <v-list-item class="px-0">
            <v-list-item-content>
                <v-list-item-title class="text-h4">
                    {{ label }}
                </v-list-item-title>
            </v-list-item-content>

            <v-spacer></v-spacer>

            <v-list-item-action>
                <restrict-role-permission-button
                    color="primary"
                    :permissions="permissionsForCreate"
                    :org-id="currentOrg.Id"
                    :engagement-id="currentEngagement.Id"
                    @click="showHideNew = true"
                >
                    New
                </restrict-role-permission-button>

                <v-dialog
                    v-model="showHideNew"
                    persistent
                    max-width="40%"
                >
                    <slot
                        name="create"
                        v-bind:onCancel="onCancel"
                        v-bind:onSave="onSave"
                        v-bind:orgId="currentOrg.Id"
                        v-bind:engagementId="currentEngagement.Id"
                    >
                    </slot>
                </v-dialog>
            </v-list-item-action>
        </v-list-item>
        <v-divider class="mb-4"></v-divider>

        <loading-container
            :loading="!allData"
        >
            <template v-slot:default="{show}">
                <full-height-base>
                    <slot
                        name="list"
                        v-bind:show="show"
                    >
                    </slot>
                </full-height-base>
            </template>
        </loading-container>
    </div>
</template>

<script lang="ts">

import Vue from 'vue'
import Component from 'vue-class-component'
import { Prop, Watch } from 'vue-property-decorator'
import { RawOrganization } from '@client/ts/types/orgs'
import { RawEngagement } from '@client/ts/types/engagements'
import { Permission } from '@client/ts/types/roles'
import RestrictRolePermissionButton from '@client/vue/loading/RestrictRolePermissionButton.vue'
import FullHeightBase from '@client/vue/shared/FullHeightBase.vue'
import LoadingContainer from '@client/vue/loading/LoadingContainer.vue'

@Component ({
    components: {
        FullHeightBase,
        LoadingContainer,
        RestrictRolePermissionButton
    }
})
export default class GenericCrudList extends Vue {
    @Prop()
    label!: string 

    @Prop({ type: Array })
    permissionsForCreate!: Permission[]

    @Prop({ type: Array })
    allData!: any[] | null
    
    showHideNew: boolean = false
    get currentOrg() : RawOrganization | null {
        return this.$store.state.org.rawOrg
    }

    get currentEngagement() : RawEngagement | null {
        return this.$store.state.engagements.rawEngagement
    }


    onCancel() {
        this.showHideNew = false
    }

    onSave(v : any) {
        this.showHideNew = false
        this.$emit('new-item', v)
    }

    mounted() {
        this.eventRefreshList()
    }

    @Watch('currentOrg')
    @Watch('currentEngagement')
    eventRefreshList() {
        this.$emit(
            'refresh-list',
            this.currentOrg?.Id,
            this.currentEngagement?.Id
        )
    }
}

</script>
