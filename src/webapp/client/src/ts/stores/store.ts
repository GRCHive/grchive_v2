import { StoreOptions } from 'vuex'

import { UserStoreModule } from '@client/ts/stores/modules/userStore'
import { ErrorStoreModule } from '@client/ts/stores/modules/errorStore'
import { AppLayoutStoreModule } from '@client/ts/stores/modules/appLayoutStore'

export interface RootState {
}

export const RootStoreOptions : StoreOptions<RootState> = {
    strict: true,
    modules: {
        user: UserStoreModule,
        errors: ErrorStoreModule,
        appLayout: AppLayoutStoreModule,
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
