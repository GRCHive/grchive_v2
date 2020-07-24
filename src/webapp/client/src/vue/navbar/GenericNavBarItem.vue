<template>
    <restrict-role-permission-container
        :permissions="accessPermissions"
        :org-id="orgId"
        :engagement-id="engagementId"
    >
        <template v-slot:default="{show}">
            <v-list-item
                :to="{name: `${item.path}`, params: `${item.params}` }"
                link
                :color="item.disabled ? `secondary` : `primary`"
                :disabled="item.disabled"
                :two-line="item.disabled"
                v-if="!item.hidden"
            >
                <v-list-item-icon v-if="item.icon != ''">
                    <v-icon>{{ item.icon }}</v-icon>
                </v-list-item-icon>

                <v-list-item-content>
                    <v-list-item-title>
                        {{ item.title }}
                    </v-list-item-title>
                    <v-list-item-subtitle v-if="item.disabled">
                        Coming Soon.
                    </v-list-item-subtitle>
                </v-list-item-content>
            </v-list-item>
        </template>
    </restrict-role-permission-container>
</template>

<script lang="ts">

import Vue from 'vue'
import Component from 'vue-class-component'
import { Prop } from 'vue-property-decorator'
import { NavLink } from '@client/ts/frontend/navLink'
import { Permission } from '@client/ts/types/roles'
import RestrictRolePermissionContainer from '@client/vue/loading/RestrictRolePermissionContainer.vue'

@Component({
    components: {
        RestrictRolePermissionContainer,
    }
})
export default class GenericNavBarItem extends Vue {
    @Prop()
    item! : NavLink

    get accessPermissions() : Permission[] {
        if (!this.item.permissions) {
            return []
        }
        return this.item.permissions
    }

    get orgId() : number {
        let orgId = this.$route.params.orgId
        if (!!orgId) {
            return Number(orgId)
        }
        return -1
    }

    get engagementId() : number {
        let engId = this.$route.params.engId
        if (!!engId) {
            return Number(engId)
        }
        return -1
    }
}

</script>
