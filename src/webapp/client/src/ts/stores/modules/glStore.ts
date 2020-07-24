import { Module } from 'vuex'
import { RawGLAccount } from '@client/ts/types/gl'
import { RootState } from '@client/ts/stores/store'
import { GrchiveApi, ErrorHandler } from '@client/ts/main'

export interface GeneralLedgerStoreState {
    rawGLAccount : RawGLAccount | null
}

export const GeneralLedgerStoreModule : Module<GeneralLedgerStoreState, RootState> = {
    namespaced: true,
    state: () => ({
        rawGLAccount: null,
    }),
    mutations: {
        setRawGeneralLedger(state : GeneralLedgerStoreState, acc : RawGLAccount | null) {
            state.rawGLAccount = acc
        }
    },
    actions: {
        initializeGeneralLedgerStore(context, { orgId , engId, glAccountId }) {
            GrchiveApi.gl.getAccount(orgId, engId, glAccountId).then((resp : RawGLAccount) => {
                context.commit('setRawGeneralLedger', resp)
            }).catch((err : any) => {
                ErrorHandler.failurePageOnError(err)
            })
        }
    },
}
