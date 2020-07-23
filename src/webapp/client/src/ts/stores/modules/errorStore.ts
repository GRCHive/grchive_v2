import { Module } from 'vuex'
import { RootState } from '@client/ts/stores/store'
import { ErrorWrapper } from '@client/ts/error'

export interface ErrorStoreState {
    relevantErrors : ErrorWrapper[]
}

export const ErrorStoreModule : Module<ErrorStoreState, RootState> = {
    namespaced: true,
    state: () => ({
        relevantErrors: []
    }),
    mutations: {
        addError(state : ErrorStoreState, err : ErrorWrapper) {
            state.relevantErrors.push(err)
        },
        removeError(state : ErrorStoreState, errId : string) {
            state.relevantErrors = state.relevantErrors.filter((ele : ErrorWrapper) => ele.displayId != errId)
        }
    }
}
