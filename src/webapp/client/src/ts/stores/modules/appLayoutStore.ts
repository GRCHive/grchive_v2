import { Module } from 'vuex'
import { RootState } from '@client/ts/stores/store'

export interface AppLayoutStoreState {
    miniNav: boolean
    userHideVerificationBanner: boolean
}

const MiniNavKey : string = "miniNavBar"

export const AppLayoutStoreModule : Module<AppLayoutStoreState, RootState> = {
    namespaced: true,
    state: () => ({
        miniNav: false,
        userHideVerificationBanner: false,
    }),
    mutations: {
        toggleMini(state : AppLayoutStoreState) {
            state.miniNav = !state.miniNav
            window.localStorage.setItem(MiniNavKey, state.miniNav ? "true" : "false")
        },
        initializeAppLayoutStore(state : AppLayoutStoreState) {
            {
                let val = window.localStorage.getItem(MiniNavKey)
                if (!!val) {
                    state.miniNav = (val == "true")
                }
            }
        },
        closeVerificationBanner(state : AppLayoutStoreState) {
            state.userHideVerificationBanner = true
        }
    },
    getters: {
        showEmailVerificationBanner(state, _2, rootState) : boolean {
            if (!rootState.user.rawUser) {
                return false
            }

            return !rootState.user.rawUser.EmailVerified && !state.userHideVerificationBanner
        }
    },
}
