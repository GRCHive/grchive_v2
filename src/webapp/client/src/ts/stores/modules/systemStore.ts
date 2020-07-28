import { Module } from 'vuex'
import { RawSystem } from '@client/ts/types/systems'
import { RootState } from '@client/ts/stores/store'
import { GrchiveApi, ErrorHandler } from '@client/ts/main'

export interface SystemStoreState {
    rawSystem : RawSystem | null
}

export const SystemStoreModule : Module<SystemStoreState, RootState> = {
    namespaced: true,
    state: () => ({
        rawSystem: null,
    }),
    mutations: {
        setRawSystem(state : SystemStoreState, rawSystem : RawSystem | null) {
            state.rawSystem = rawSystem
        }
    },
    actions: {
        initializeSystemStore(context, { orgId , engId, systemId }) {
            GrchiveApi.systems.getSystem(orgId, engId, systemId).then((resp : RawSystem) => {
                context.commit('setRawSystem', resp)
            }).catch((err : any) => {
                ErrorHandler.failurePageOnError(err)
            })
        }
    },
}
