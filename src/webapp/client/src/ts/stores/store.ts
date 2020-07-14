import { StoreOptions } from 'vuex'

import { UserStoreModule } from '@client/ts/stores/modules/userStore'
import { ErrorStoreModule } from '@client/ts/stores/modules/errorStore'

export interface RootState {
}

export const RootStoreOptions : StoreOptions<RootState> = {
    strict: true,
    modules: {
        user: UserStoreModule,
        errors: ErrorStoreModule
    },
    actions: {
        initialize(context) {
            context.dispatch('user/initializeUserStore')
        }
    },
    getters: {
        isFinishedLoading(_, getters) : boolean {
            return getters['user/isFinishedLoading']
        }
    }
}
