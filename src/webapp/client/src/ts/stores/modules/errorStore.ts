import { Module } from 'vuex'
import { RootState } from '@client/ts/stores/store'
import { GrchiveError, GrchiveApiError } from '@client/ts/types/errors'

interface StoreState {
    relevantErrors : GrchiveError[]
}

export const ErrorStoreModule : Module<StoreState, RootState> = {
    namespaced: true,
    state: () => ({
        relevantErrors: []
    }),
    mutations: {
        addApiError(state : StoreState, err : any) {
            state.relevantErrors.unshift(new GrchiveApiError(err))
        }
    }
}
