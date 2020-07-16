import { StoreOptions } from 'vuex'

import { UserStoreModule, UserStoreState } from '@client/ts/stores/modules/userStore'
import { ErrorStoreModule, ErrorStoreState } from '@client/ts/stores/modules/errorStore'
import { AppLayoutStoreModule, AppLayoutStoreState } from '@client/ts/stores/modules/appLayoutStore'
import { OrgStoreModule, OrgStoreState } from '@client/ts/stores/modules/orgStore'

export interface RootState {
    user: UserStoreState
    errors: ErrorStoreState
    appLayout: AppLayoutStoreState
    org: OrgStoreState
}

export const RootStoreOptions : StoreOptions<RootState> = {
    strict: true,
    modules: {
        user: UserStoreModule,
        errors: ErrorStoreModule,
        appLayout: AppLayoutStoreModule,
        org : OrgStoreModule
    },
    actions: {
        initialize(context) {
            context.dispatch('user/initializeUserStore')
            context.commit('appLayout/initializeAppLayoutStore')
        }
    },
    getters: {
        isFinishedLoading(_, getters) : boolean {
            return getters['user/isFinishedLoading']
        }
    }
}
