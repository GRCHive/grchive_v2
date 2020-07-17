<template>
    <v-treeview
        activatable
        open-all
        :items="orgItems"
        transition
        @update:active="goToOrg"
        ref="treeview"
    >
        <template v-slot:prepend="{ item, open }">
            <v-icon
                v-if="!item.value.ParentOrgId"
            >
                {{ open ? "mdi-account-group" : "mdi-account-group-outline" }}
            </v-icon>

            <v-icon
                v-else-if="item.children.length > 0"
            >
                {{ open ? "mdi-account-multiple" : "mdi-account-multiple-outline" }}
            </v-icon>

            <v-icon
                v-else
            >
                {{ open ? "mdi-account" : "mdi-account-outline" }}
            </v-icon>
        </template>

        <template v-slot:append="{ item }">
            <v-dialog
                persistent
                max-width="40%"
                v-model="showHideNew"
            >
                <org-save-edit-dialog
                    @cancel-edit="showHideNew = false"
                    @save-edit="onSaveSuborg"
                    :parent-org-id="item.id"
                >
                </org-save-edit-dialog>
            </v-dialog>

            <restrict-role-permission-button
                v-if="allowAddSuborgs"
                icon
                color="primary"
                :permissions="permissionsForCreate"
                :org-id="item.id"
                @click="showHideNew = true"
            >
                <v-icon>
                    mdi-plus
                </v-icon>
            </restrict-role-permission-button>
        </template>
    </v-treeview>
</template>

<script lang="ts">

import Vue from 'vue'
import Component from 'vue-class-component'
import { VTreeview } from 'vuetify/lib'
import { Watch, Prop } from 'vue-property-decorator'
import { Permission } from '@client/ts/types/roles'
import { RawOrganization, OrgTree } from '@client/ts/types/orgs'
import RestrictRolePermissionButton from '@client/vue/loading/RestrictRolePermissionButton.vue'
import OrgSaveEditDialog from '@client/vue/types/orgs/OrgSaveEditDialog.vue'

@Component({
    components: {
        RestrictRolePermissionButton,
        OrgSaveEditDialog,
    }
})
export default class OrgTreeViewer extends Vue {
    @Prop({ required: true })
    orgs! : RawOrganization[] | null

    @Prop()
    rootOrg! : RawOrganization | null

    workingOrgs : RawOrganization[] = []

    @Watch('orgs')
    syncOrgs() {
        if (!this.orgs) {
            this.workingOrgs = []
        } else {
            this.workingOrgs = this.orgs.slice()
        }
    }

    @Watch('workingOrgs')
    openAll() {
        Vue.nextTick(() => {
            //@ts-ignore
            this.$refs.treeview.updateAll(true)
        })
    }

    @Prop({ default: -1 })
    currentOrgId!: number

    @Prop({ type: Boolean, default: false })
    allowAddSuborgs! : boolean

    showHideNew: boolean = false

    get permissionsForCreate() : Permission[] {
        return [Permission.POrgProfileCreate]
    }

    get orgTree() : OrgTree {
        return new OrgTree(this.workingOrgs, this.rootOrg)
    }

    get orgItems() : any[] {
        // id: unique key
        // name: Text to display
        // children: children in tree.
        let createTree = (ele : OrgTree) : any => {
            if (!ele.nodeOrg) {
                return ele.childNodes.map(createTree)
            } else {
                return {
                    id: ele.nodeOrg!.Id,
                    name: ele.nodeOrg!.Name,
                    children: ele.childNodes.map(createTree),
                    value: ele.nodeOrg,
                }
            }
        }

        return [createTree(this.orgTree)].flat()
    }

    goToOrg(inp : Array<number>) {
        if (!inp.length) {
            return
        }

        if (inp[0] == this.currentOrgId) {
            return
        }

        this.$router.push({
            name: 'orgHome',
            params: { orgId: `${inp[0]}` },
        })
    }

    onSaveSuborg(org : RawOrganization) {
        this.workingOrgs.unshift(org)
        this.showHideNew = false
    }

    mounted() {
        this.syncOrgs()
    }
}

</script>

<style scoped>

>>>.v-treeview-node {
    cursor: pointer;
}

</style>
