import { Module } from 'vuex'
import { RawUser } from '@client/ts/types/users'
import { RootState } from '@client/ts/stores/store'

import { GrchiveApi } from '@client/ts/api/client'

interface StoreState {
    rawUser : RawUser | null
}

export const UserStoreModule : Module<StoreState, RootState> = {
    namespaced: true,
    state: () => ({
        rawUser : null,
    }),
    mutations: {
        setRawUser(state : StoreState, rawUser : RawUser) {
            state.rawUser = rawUser
        }
    },
    actions: {
        initializeUserStore(context) {
            GrchiveApi.user.getCurrentUser().then((resp : RawUser) => {
                context.commit('setRawUser', resp)
            }).catch((err : any) => {
                context.commit('errors/addApiError', err, { root: true })
            })
        }
    },
    getters: {
        isFinishedLoading(state) : boolean {
            return !!state.rawUser
        }
    }
}
