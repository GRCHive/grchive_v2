import { Module } from 'vuex'
import { RawOrganization } from '@client/ts/types/orgs'
import { RootState } from '@client/ts/stores/store'
import { GrchiveApi } from '@client/ts/main'

export interface OrgStoreState {
    rawOrg : RawOrganization | null
}

export const OrgStoreModule : Module<OrgStoreState, RootState> = {
    namespaced: true,
    state: () => ({
        rawOrg : null,
    }),
    mutations: {
        setRawOrg(state : OrgStoreState, rawOrg : RawOrganization) {
            state.rawOrg = rawOrg
        }
    },
    actions: {
        initializeOrgStore(context, orgId : number) {
            GrchiveApi.orgs.getOrg(orgId).then((resp : RawOrganization | null) => {
                context.commit('setRawOrg', resp)
            })
        }
    },
    getters: {
        isFinishedLoading(state) : boolean {
            return !!state.rawOrg
        }
    }
}
