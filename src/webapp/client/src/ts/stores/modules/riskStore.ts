import { Module } from 'vuex'
import { RawRisk } from '@client/ts/types/risks'
import { RootState } from '@client/ts/stores/store'
import { GrchiveApi, ErrorHandler } from '@client/ts/main'

export interface RiskStoreState {
    rawRisk : RawRisk | null
}

export const RiskStoreModule : Module<RiskStoreState, RootState> = {
    namespaced: true,
    state: () => ({
        rawRisk: null,
    }),
    mutations: {
        setRawRisk(state : RiskStoreState, rawRisk : RawRisk | null) {
            state.rawRisk = rawRisk
        }
    },
    actions: {
        initializeRiskStore(context, { orgId , engId, riskId }) {
            GrchiveApi.risks.getRisk(orgId, engId, riskId).then((resp : RawRisk) => {
                context.commit('setRawRisk', resp)
            }).catch((err : any) => {
                ErrorHandler.failurePageOnError(err)
            })
        }
    },
}
