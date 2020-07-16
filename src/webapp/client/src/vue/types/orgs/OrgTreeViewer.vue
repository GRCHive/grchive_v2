<template>
    <v-treeview
        activatable
        open-all
        :items="orgItems"
        transition
        @update:active="goToOrg"
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
    </v-treeview>
</template>

<script lang="ts">

import Vue from 'vue'
import Component from 'vue-class-component'
import { Prop } from 'vue-property-decorator'
import { RawOrganization, OrgTree } from '@client/ts/types/orgs'

@Component
export default class OrgTreeViewer extends Vue {
    @Prop({ required: true })
    orgs! : RawOrganization[]

    @Prop({ default: -1 })
    currentOrgId!: number

    get orgTree() : OrgTree {
        return new OrgTree(this.orgs)
    }

    get orgItems() : any[] {
        // id: unique key
        // name: Text to display
        // children: children in tree.
        let createTree = (ele : OrgTree) : any => {
            return {
                id: ele.nodeOrg!.Id,
                name: ele.nodeOrg!.Name,
                children: ele.childNodes.map(createTree),
                value: ele,
            }
        }

        return this.orgTree.childNodes.map(createTree)
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
}

</script>

<style scoped>

>>>.v-treeview-node {
    cursor: pointer;
}

</style>
