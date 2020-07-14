import { Module } from 'vuex'
import { RootState } from '@client/ts/stores/store'
import { GrchiveError, GrchiveApiError } from '@client/ts/types/errors'

export interface ErrorStoreState {
    relevantErrors : GrchiveError[]
}

export const ErrorStoreModule : Module<ErrorStoreState, RootState> = {
    namespaced: true,
    state: () => ({
        relevantErrors: []
    }),
    mutations: {
        addApiError(state : ErrorStoreState, err : any) {
            state.relevantErrors.unshift(new GrchiveApiError(err))
        }
    }
}
