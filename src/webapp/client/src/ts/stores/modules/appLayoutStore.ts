import { Module } from 'vuex'
import { RootState } from '@client/ts/stores/store'

interface StoreState {
    miniNav: boolean
}

const MiniNavKey : string = "miniNavBar"

export const AppLayoutStoreModule : Module<StoreState, RootState> = {
    namespaced: true,
    state: () => ({
        miniNav: false,
    }),
    mutations: {
        toggleMini(state : StoreState) {
            state.miniNav = !state.miniNav
            window.localStorage.setItem(MiniNavKey, state.miniNav ? "true" : "false")
        },
        initializeAppLayoutStore(state : StoreState) {
            {
                let val = window.localStorage.getItem(MiniNavKey)
                if (!!val) {
                    state.miniNav = (val == "true")
                }
            }
        }
    },
}
