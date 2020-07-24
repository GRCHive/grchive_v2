import { Module } from 'vuex'
import { RawControl } from '@client/ts/types/controls'
import { RootState } from '@client/ts/stores/store'
import { GrchiveApi, ErrorHandler } from '@client/ts/main'

export interface ControlStoreState {
    rawControl : RawControl | null
}

export const ControlStoreModule : Module<ControlStoreState, RootState> = {
    namespaced: true,
    state: () => ({
        rawControl: null,
    }),
    mutations: {
        setRawControl(state : ControlStoreState, rawControl : RawControl | null) {
            state.rawControl = rawControl
        }
    },
    actions: {
        initializeControlStore(context, { orgId , engId, controlId }) {
            GrchiveApi.controls.getControl(orgId, engId, controlId).then((resp : RawControl) => {
                context.commit('setRawControl', resp)
            }).catch((err : any) => {
                ErrorHandler.failurePageOnError(err)
            })
        }
    },
}
