import { Module } from 'vuex'
import { RawUser } from '@client/ts/types/users'
import { RootState } from '@client/ts/stores/store'
import { GrchiveApi } from '@client/ts/main'

export interface UserStoreState {
    rawUser : RawUser | null
}

export const UserStoreModule : Module<UserStoreState, RootState> = {
    namespaced: true,
    state: () => ({
        rawUser : null,
    }),
    mutations: {
        setRawUser(state : UserStoreState, rawUser : RawUser) {
            state.rawUser = rawUser
        }
    },
    actions: {
        initializeUserStore(context) {
            GrchiveApi.user.getCurrentUser().then((resp : RawUser | null) => {
                context.commit('setRawUser', resp)
            })
        }
    },
    getters: {
        isFinishedLoading(state) : boolean {
            return !!state.rawUser
        }
    }
}
