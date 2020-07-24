import { Module } from 'vuex'
import { RawOrganization } from '@client/ts/types/orgs'
import { RootState } from '@client/ts/stores/store'
import { GrchiveApi, ErrorHandler } from '@client/ts/main'

export interface OrgStoreState {
    rawOrg : RawOrganization | null
}

export const OrgStoreModule : Module<OrgStoreState, RootState> = {
    namespaced: true,
    state: () => ({
        rawOrg : null,
    }),
    mutations: {
        setRawOrg(state : OrgStoreState, rawOrg : RawOrganization | null) {
            state.rawOrg = rawOrg
        }
    },
    actions: {
        initializeOrgStore(context, orgId : number) {
            GrchiveApi.orgs.getOrg(orgId).then((resp : RawOrganization) => {
                context.commit('setRawOrg', resp)
            }).catch((err: any) => {
                ErrorHandler.failurePageOnError(err)
            })
        }
    },
    getters: {
        isFinishedLoading(state) : boolean {
            return !!state.rawOrg
        }
    }
}
