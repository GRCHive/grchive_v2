import { Module } from 'vuex'
import { 
    RawServer,
    InventoryType,
} from '@client/ts/types/inventory'
import { RootState } from '@client/ts/stores/store'
import { GrchiveApi, ErrorHandler } from '@client/ts/main'

export interface InventoryStoreState {
    rawServer: RawServer | null
}

export const InventoryStoreModule : Module<InventoryStoreState, RootState> = {
    namespaced: true,
    state: () => ({
        rawServer: null,
    }),
    mutations: {
        setRawServer(state : InventoryStoreState, rawServer : RawServer | null) {
            state.rawServer = rawServer
        }
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
                default:
                    break
            }
        }
    },
}
