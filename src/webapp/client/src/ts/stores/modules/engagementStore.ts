import { Module } from 'vuex'
import { RawEngagement } from '@client/ts/types/engagements'
import { RootState } from '@client/ts/stores/store'
import { GrchiveApi, ErrorHandler } from '@client/ts/main'

export interface EngagementStoreState {
    rawEngagement : RawEngagement | null
}

export const EngagementStoreModule : Module<EngagementStoreState, RootState> = {
    namespaced: true,
    state: () => ({
        rawEngagement: null,
    }),
    mutations: {
        setRawEngagement(state : EngagementStoreState, rawEngagement : RawEngagement | null) {
            state.rawEngagement = rawEngagement
        }
    },
    actions: {
        initializeEngagementStore(context, { orgId , engId }) {
            GrchiveApi.engagements.getEngagement(orgId, engId).then((resp : RawEngagement) => {
                context.commit('setRawEngagement', resp)
            }).catch((err : any) => {
                ErrorHandler.failurePageOnError(err)
            })
        }
    },
}

