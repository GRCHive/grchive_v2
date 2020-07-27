<template>
    <scoping-template
        page-name="Inventory"
    >
        <template v-slot:content>
            <v-tabs>
                <restrict-role-permission-tab
                    :permissions="serversPermissions"
                    :to="serversTo"
                    :org-id="currentOrg.Id"
                    :engagement-id="currentEngagement.Id"
                >
                    Servers
                </restrict-role-permission-tab>
                <restrict-role-permission-tab
                    :permissions="desktopsPermissions"
                    :to="desktopsTo"
                    :org-id="currentOrg.Id"
                    :engagement-id="currentEngagement.Id"
                >
                    Desktops
                </restrict-role-permission-tab>
                <restrict-role-permission-tab
                    :permissions="laptopsPermissions"
                    :to="laptopsTo"
                    :org-id="currentOrg.Id"
                    :engagement-id="currentEngagement.Id"
                >
                    Laptops
                </restrict-role-permission-tab>
                <restrict-role-permission-tab
                    :permissions="mobilePermissions"
                    :to="mobileTo"
                    :org-id="currentOrg.Id"
                    :engagement-id="currentEngagement.Id"
                >
                    Mobile
                </restrict-role-permission-tab>
                <restrict-role-permission-tab
                    :permissions="embeddedPermissions"
                    :to="embeddedTo"
                    :org-id="currentOrg.Id"
                    :engagement-id="currentEngagement.Id"
                >
                    Embedded
                </restrict-role-permission-tab>
            </v-tabs>
            <router-view></router-view>
        </template>
    </scoping-template>
</template>

<script lang="ts">

import Vue from 'vue'
import Component from 'vue-class-component'
import { Watch } from 'vue-property-decorator'
import { RawOrganization } from '@client/ts/types/orgs'
import { RawEngagement } from '@client/ts/types/engagements'
import { Permission } from '@client/ts/types/roles'
import { ErrorHandler } from '@client/ts/main'
import { createManualUnauthorizedError } from '@client/ts/error'
import ScopingTemplate from '@client/vue/orgs/engagements/scoping/ScopingTemplate.vue'
import RestrictRolePermissionTab from '@client/vue/loading/RestrictRolePermissionTab.vue'

@Component({
    components: {
        ScopingTemplate,
        RestrictRolePermissionTab
    }
})
export default class ScopingInventory extends Vue {
    get currentOrg() : RawOrganization | null {
        return this.$store.state.org.rawOrg
    }

    get currentEngagement() : RawEngagement | null {
        return this.$store.state.engagements.rawEngagement
    }

    get serversTo() : any {
        return {
            name: 'scopingServers',
            params: this.$route.params,
        }
    }

    get serversPermissions() : Permission[] {
        return [Permission.PServersList]
    }

    get desktopsTo() : any {
        return {
            name: 'scopingDesktops',
            params: this.$route.params,
        }
    }

    get desktopsPermissions() : Permission[] {
        return [Permission.PDesktopsList]
    }

    get laptopsTo() : any {
        return {
            name: 'scopingLaptops',
            params: this.$route.params,
        }
    }

    get laptopsPermissions() : Permission[] {
        return [Permission.PLaptopsList]
    }

    get mobileTo() : any {
        return {
            name: 'scopingMobile',
            params: this.$route.params,
        }
    }

    get mobilePermissions() : Permission[] {
        return [Permission.PMobileList]
    }

    get embeddedTo() : any {
        return {
            name: 'scopingEmbedded',
            params: this.$route.params,
        }
    }

    get embeddedPermissions() : Permission[] {
        return [Permission.PEmbeddedList]
    }

    get inventoryPermissions(): any {
        if (!this.currentOrg || !this.currentEngagement) {
            return null
        }

        return {
            server: this.$store.getters['permission/currentUserHasPermissions'](this.currentOrg!.Id, this.currentEngagement!.Id, this.serversPermissions),
            desktop: this.$store.getters['permission/currentUserHasPermissions'](this.currentOrg!.Id, this.currentEngagement!.Id, this.desktopsPermissions),
            laptop: this.$store.getters['permission/currentUserHasPermissions'](this.currentOrg!.Id, this.currentEngagement!.Id, this.laptopsPermissions),
            mobile: this.$store.getters['permission/currentUserHasPermissions'](this.currentOrg!.Id, this.currentEngagement!.Id, this.mobilePermissions),
            embedded: this.$store.getters['permission/currentUserHasPermissions'](this.currentOrg!.Id, this.currentEngagement!.Id, this.embeddedPermissions),
        }
    }

    // Need to redirect to the appropriate subpage based on permissions if we land on the 'scopingInventory' page.
    // If we don't have permissions then we can redirect to an error page.
    @Watch('inventoryPermissions')
    @Watch('$route')
    redirectToSubpage() {
        if (this.$route.name !== 'scopingInventory') {
            return
        }

        console.log(this.currentOrg, this.currentEngagement)
        if (!this.currentOrg || !this.currentEngagement) {
            return
        }

        console.log(this.inventoryPermissions)
        // Do nothing until we load all permissions so we can know for sure when we want to direct the user to an error page.
        if (this.inventoryPermissions.server === null ||
            this.inventoryPermissions.desktop === null ||
            this.inventoryPermissions.laptop === null ||
            this.inventoryPermissions.mobile === null ||
            this.inventoryPermissions.embedded == null) {
            return
        }

        if (this.inventoryPermissions.server === true) {
            this.$router.replace({
                name: 'scopingServers',
                params: this.$route.params
            })
        } else if (this.inventoryPermissions.desktop === true) {
            this.$router.replace({
                name: 'scopingDesktops',
                params: this.$route.params
            })
        } else if (this.inventoryPermissions.laptop === true) { 
            this.$router.replace({
                name: 'scopingLaptops',
                params: this.$route.params
            })
        } else if (this.inventoryPermissions.mobile === true) {
            this.$router.replace({
                name: 'scopingMobile',
                params: this.$route.params
            })
        } else if (this.inventoryPermissions.embedded === true) {
            this.$router.replace({
                name: 'scopingEmbedded',
                params: this.$route.params
            })
        } else {
            ErrorHandler.failurePageOnManualError(
                createManualUnauthorizedError('You do not have sufficient privileges to view any inventory items.',
                    [
                        this.serversPermissions,
                        this.desktopsPermissions,
                        this.laptopsPermissions,
                        this.mobilePermissions,
                        this.embeddedPermissions,
                    ].flat(),
                    true,
                )
            )
        }
    }

    mounted() {
        this.redirectToSubpage()
    }
}

</script>
