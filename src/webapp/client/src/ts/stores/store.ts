import { StoreOptions } from 'vuex'

import { UserStoreModule, UserStoreState } from '@client/ts/stores/modules/userStore'
import { PermissionStoreModule, PermissionStoreState } from '@client/ts/stores/modules/permissionStore'
import { AppLayoutStoreModule, AppLayoutStoreState } from '@client/ts/stores/modules/appLayoutStore'
import { OrgStoreModule, OrgStoreState } from '@client/ts/stores/modules/orgStore'
import { EngagementStoreModule, EngagementStoreState } from '@client/ts/stores/modules/engagementStore'
import { RiskStoreModule, RiskStoreState } from '@client/ts/stores/modules/riskStore'
import { GeneralLedgerStoreModule, GeneralLedgerStoreState } from '@client/ts/stores/modules/glStore'
import { ControlStoreModule, ControlStoreState } from '@client/ts/stores/modules/controlStore'
import { ErrorStoreModule, ErrorStoreState } from '@client/ts/stores/modules/errorStore'
import { VendorStoreModule, VendorStoreState } from '@client/ts/stores/modules/vendorStore'
import { InventoryStoreModule, InventoryStoreState } from '@client/ts/stores/modules/inventoryStore'
import { DatabaseStoreModule, DatabaseStoreState } from '@client/ts/stores/modules/databaseStore'
import { SystemStoreModule, SystemStoreState } from '@client/ts/stores/modules/systemStore'

import { InventoryType } from '@client/ts/types/inventory'

export interface RootState {
    user: UserStoreState
    appLayout: AppLayoutStoreState
    org: OrgStoreState
    permission: PermissionStoreState
    engagements: EngagementStoreState
    risks: RiskStoreState
    controls: ControlStoreState
    gl : GeneralLedgerStoreState
    errors : ErrorStoreState
    vendors: VendorStoreState
    inventory: InventoryStoreState
    databases: DatabaseStoreState
    systems : SystemStoreState
}

export interface CurrentResourceInitialization {
    orgId? : number
    engagementId?: number
    riskId?: number
    controlId?: number
    glAccountId? : number
    vendorId? : number
    vendorProductId? : number
    serverId? : number
    desktopId? : number
    laptopId? : number
    mobileId? : number
    embeddedId? : number
    databaseId? : number
    systemId? : number
}

export const RootStoreOptions : StoreOptions<RootState> = {
    strict: true,
    modules: {
        user: UserStoreModule,
        appLayout: AppLayoutStoreModule,
        org : OrgStoreModule,
        permission: PermissionStoreModule,
        engagements: EngagementStoreModule,
        risks: RiskStoreModule,
        controls: ControlStoreModule,
        gl: GeneralLedgerStoreModule,
        errors: ErrorStoreModule,
        vendors: VendorStoreModule,
        inventory: InventoryStoreModule,
        databases: DatabaseStoreModule,
        systems : SystemStoreModule,
    },
    actions: {
        initialize(context) {
            context.dispatch('user/initializeUserStore')
            context.commit('appLayout/initializeAppLayoutStore')
        },
        initializeCurrentResource(context, params : CurrentResourceInitialization) {
            if (!!params.orgId) {
                if (!context.state.org.rawOrg || context.state.org.rawOrg.Id != params.orgId) {
                    context.dispatch('org/initializeOrgStore', params.orgId)
                }
            } else {
                context.commit('org/setRawOrg', null)
            }

            if (!!params.engagementId) {
                if (!context.state.engagements.rawEngagement || context.state.engagements.rawEngagement.Id != params.engagementId) {
                    context.dispatch('engagements/initializeEngagementStore', {
                        orgId: params.orgId,
                        engId: params.engagementId,
                    })
                }
            } else {
                context.commit('engagements/setRawEngagement', null)
            }

            if (!!params.riskId) {
                if (!context.state.risks.rawRisk || context.state.risks.rawRisk.Id != params.riskId) {
                    context.dispatch('risks/initializeRiskStore', {
                        orgId: params.orgId,
                        engId: params.engagementId,
                        riskId: params.riskId,
                    })
                }
            } else {
                context.commit('risks/setRawRisk', null)
            }

            if (!!params.controlId) {
                if (!context.state.controls.rawControl || context.state.controls.rawControl.Id != params.controlId) {
                    context.dispatch('controls/initializeControlStore', {
                        orgId: params.orgId,
                        engId: params.engagementId,
                        controlId: params.controlId,
                    })
                }
            } else {
                context.commit('controls/setRawControl', null)
            }

            if (!!params.glAccountId) {
                if (!context.state.gl.rawGLAccount || context.state.gl.rawGLAccount.Id != params.glAccountId) {
                    context.dispatch('gl/initializeGeneralLedgerStore', {
                        orgId: params.orgId,
                        engId: params.engagementId,
                        glAccountId: params.glAccountId,
                    })
                }
            } else {
                context.commit('gl/setRawGeneralLedger', null)
            }

            if (!!params.vendorId) {
                if (!context.state.vendors.rawVendor || context.state.vendors.rawVendor.Id != params.vendorId) {
                    context.dispatch('vendors/initializeVendorStore', {
                        orgId: params.orgId,
                        engId: params.engagementId,
                        vendorId: params.vendorId,
                    })
                }
            } else {
                context.commit('vendors/setRawVendor', null)
            }

            if (!!params.vendorProductId) {
                if (!context.state.vendors.rawVendorProduct || context.state.vendors.rawVendorProduct.Id != params.vendorProductId) {
                    context.dispatch('vendors/initializeVendorProductStore', {
                        orgId: params.orgId,
                        engId: params.engagementId,
                        vendorId: params.vendorId,
                        vendorProductId: params.vendorProductId,
                    })
                }
            } else {
                context.commit('vendors/setRawVendorProduct', null)
            }

            if (!!params.serverId) {
                if (!context.state.inventory.rawServer || context.state.inventory.rawServer.Id != params.serverId) {
                    context.dispatch('inventory/initializeInventoryStore', {
                        orgId: params.orgId,
                        engId: params.engagementId,
                        inventoryId: params.serverId,
                        inventoryType: InventoryType.ITServer,
                    })
                }
            } else {
                context.commit('inventory/setRawServer', null)
            }

            if (!!params.desktopId) {
                if (!context.state.inventory.rawDesktop || context.state.inventory.rawDesktop.Id != params.desktopId) {
                    context.dispatch('inventory/initializeInventoryStore', {
                        orgId: params.orgId,
                        engId: params.engagementId,
                        inventoryId: params.desktopId,
                        inventoryType: InventoryType.ITDesktop,
                    })
                }
            } else {
                context.commit('inventory/setRawDesktop', null)
            }

            if (!!params.laptopId) {
                if (!context.state.inventory.rawLaptop || context.state.inventory.rawLaptop.Id != params.laptopId) {
                    context.dispatch('inventory/initializeInventoryStore', {
                        orgId: params.orgId,
                        engId: params.engagementId,
                        inventoryId: params.laptopId,
                        inventoryType: InventoryType.ITLaptop,
                    })
                }
            } else {
                context.commit('inventory/setRawLaptop', null)
            }

            if (!!params.mobileId) {
                if (!context.state.inventory.rawMobile || context.state.inventory.rawMobile.Id != params.mobileId) {
                    context.dispatch('inventory/initializeInventoryStore', {
                        orgId: params.orgId,
                        engId: params.engagementId,
                        inventoryId: params.mobileId,
                        inventoryType: InventoryType.ITMobile,
                    })
                }
            } else {
                context.commit('inventory/setRawMobile', null)
            }

            if (!!params.embeddedId) {
                if (!context.state.inventory.rawEmbedded || context.state.inventory.rawEmbedded.Id != params.embeddedId) {
                    context.dispatch('inventory/initializeInventoryStore', {
                        orgId: params.orgId,
                        engId: params.engagementId,
                        inventoryId: params.embeddedId,
                        inventoryType: InventoryType.ITEmbedded,
                    })
                }
            } else {
                context.commit('inventory/setRawEmbedded', null)
            }

            if (!!params.databaseId) {
                if (!context.state.databases.rawDatabase || context.state.databases.rawDatabase.Id != params.databaseId) {
                    context.dispatch('databases/initializeDatabaseStore', {
                        orgId: params.orgId,
                        engId: params.engagementId,
                        databaseId: params.databaseId,
                    })
                }
            } else {
                context.commit('databases/setRawDatabase', null)
            }

            if (!!params.systemId) {
                if (!context.state.systems.rawSystem || context.state.systems.rawSystem.Id != params.systemId) {
                    context.dispatch('systems/initializeSystemStore', {
                        orgId: params.orgId,
                        engId: params.engagementId,
                        systemId: params.systemId,
                    })
                }
            } else {
                context.commit('systems/setRawSystem', null)
            }
        }
    },
    getters: {
        isFinishedLoading(_, getters) : boolean {
            return getters['user/isFinishedLoading']
        }
    }
}
