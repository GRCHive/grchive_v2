import { Module } from 'vuex'
import { 
    RawServer,
    RawDesktop,
    RawLaptop,
    RawMobile,
    RawEmbedded,
    InventoryType,
} from '@client/ts/types/inventory'
import { RootState } from '@client/ts/stores/store'
import { GrchiveApi, ErrorHandler } from '@client/ts/main'

export interface InventoryStoreState {
    rawServer: RawServer | null
    rawDesktop: RawDesktop | null
    rawLaptop: RawLaptop | null
    rawMobile: RawMobile | null
    rawEmbedded: RawEmbedded | null
}

export const InventoryStoreModule : Module<InventoryStoreState, RootState> = {
    namespaced: true,
    state: () => ({
        rawServer: null,
        rawDesktop: null,
        rawLaptop: null,
        rawMobile: null,
        rawEmbedded: null,
    }),
    mutations: {
        setRawServer(state : InventoryStoreState, rawServer : RawServer | null) {
            state.rawServer = rawServer
        },
        setRawDesktop(state : InventoryStoreState, rawDesktop : RawDesktop | null) {
            state.rawDesktop = rawDesktop
        },
        setRawLaptop(state : InventoryStoreState, rawLaptop : RawLaptop | null) {
            state.rawLaptop = rawLaptop
        },
        setRawMobile(state : InventoryStoreState, rawMobile : RawMobile | null) {
            state.rawMobile = rawMobile
        },
        setRawEmbedded(state : InventoryStoreState, rawEmbedded : RawEmbedded | null) {
            state.rawEmbedded = rawEmbedded
        },
    },
    actions: {
        initializeInventoryStore(context, { orgId , engId, inventoryId, inventoryType }) {
            switch (inventoryType) {
                case InventoryType.ITServer:
                    GrchiveApi.inventory.getInventory<RawServer>(inventoryType, orgId, engId, inventoryId).then((resp : RawServer) => {
                        context.commit('setRawServer', resp)
                    }).catch((err : any) => {
                        ErrorHandler.failurePageOnError(err)
                    })
                    break
                case InventoryType.ITDesktop:
                    GrchiveApi.inventory.getInventory<RawDesktop>(inventoryType, orgId, engId, inventoryId).then((resp : RawDesktop) => {
                        context.commit('setRawDesktop', resp)
                    }).catch((err : any) => {
                        ErrorHandler.failurePageOnError(err)
                    })
                    break
                case InventoryType.ITLaptop:
                    GrchiveApi.inventory.getInventory<RawLaptop>(inventoryType, orgId, engId, inventoryId).then((resp : RawLaptop) => {
                        context.commit('setRawLaptop', resp)
                    }).catch((err : any) => {
                        ErrorHandler.failurePageOnError(err)
                    })
                    break
                case InventoryType.ITMobile:
                    GrchiveApi.inventory.getInventory<RawMobile>(inventoryType, orgId, engId, inventoryId).then((resp : RawMobile) => {
                        context.commit('setRawMobile', resp)
                    }).catch((err : any) => {
                        ErrorHandler.failurePageOnError(err)
                    })
                    break
                case InventoryType.ITEmbedded:
                    GrchiveApi.inventory.getInventory<RawEmbedded>(inventoryType, orgId, engId, inventoryId).then((resp : RawEmbedded) => {
                        context.commit('setRawEmbedded', resp)
                    }).catch((err : any) => {
                        ErrorHandler.failurePageOnError(err)
                    })
                    break
                default:
                    break
            }
        }
    },
}
