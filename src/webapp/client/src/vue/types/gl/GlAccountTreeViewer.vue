<template>
    <v-treeview
        activatable
        open-all
        :items="accountItems"
        transition
        @update:active="goToAccount"
        ref="treeview"
    >
        <template v-slot:prepend="{ item, open }">
            <template v-if="item.children.length > 0">
                <v-icon>
                    {{ open ? "mdi-folder-open" : "mdi-folder" }}
                </v-icon>
            </template>
        </template>

        <template v-slot:append="{ item }">
            <span>{{ accountTypeToString(item.value.AccountType) }}</span>

            <v-icon
                v-if="item.value.FinanciallyRelevant"
                color="success"
            >
                mdi-currency-usd
            </v-icon>

            <v-icon
                v-else
                color="error"
            >
                mdi-currency-usd-off
            </v-icon>
        </template>

    </v-treeview>
</template>

<script lang="ts">

import Vue from 'vue'
import Component from 'vue-class-component'
import { Prop } from 'vue-property-decorator'
import { RawGLAccount, GLAccountTree, accountTypeToString } from '@client/ts/types/gl'

@Component
export default class GlAccountTreeViewer extends Vue {
    @Prop({required: true})
    accounts! : RawGLAccount[]

    @Prop({ default: null})
    root! : RawGLAccount | null

    accountTypeToString = accountTypeToString

    get accountTree() : GLAccountTree {
        return new GLAccountTree(this.accounts, this.root)
    }

    get accountItems() : any[] {
        // id: unique key
        // name: Text to display
        // children: children in tree.
        let createTree = (ele : GLAccountTree) : any => {
            if (!ele.nodeAcc) {
                return ele.childNodes.map(createTree)
            } else {
                return {
                    id: ele.nodeAcc!.Id,
                    name: `${ele.nodeAcc!.AccountId}: ${ele.nodeAcc!.Name}`,
                    children: ele.childNodes.map(createTree),
                    value: ele.nodeAcc,
                }
            }
        }

        return [createTree(this.accountTree)].flat()
    }

    goToAccount(inp : Array<number>) {
        if (!inp.length) {
            return
        }

        if (!!this.root) {
            if (inp[0] == this.root.Id) {
                return
            }
        }

        this.$router.push({
            name: 'glAccHome',
            params: {
                ...this.$route.params,
                accId: `${inp[0]}`,
            },
        })
    }
}

</script>

<style scoped>

>>>.v-treeview-node {
    cursor: pointer;
}

</style>
