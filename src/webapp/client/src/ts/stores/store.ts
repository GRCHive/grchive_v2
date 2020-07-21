import { StoreOptions } from 'vuex'

import { UserStoreModule, UserStoreState } from '@client/ts/stores/modules/userStore'
import { PermissionStoreModule, PermissionStoreState } from '@client/ts/stores/modules/permissionStore'
import { ErrorStoreModule, ErrorStoreState } from '@client/ts/stores/modules/errorStore'
import { AppLayoutStoreModule, AppLayoutStoreState } from '@client/ts/stores/modules/appLayoutStore'
import { OrgStoreModule, OrgStoreState } from '@client/ts/stores/modules/orgStore'
import { EngagementStoreModule, EngagementStoreState } from '@client/ts/stores/modules/engagementStore'
import { RiskStoreModule, RiskStoreState } from '@client/ts/stores/modules/riskStore'

export interface RootState {
    user: UserStoreState
    errors: ErrorStoreState
    appLayout: AppLayoutStoreState
    org: OrgStoreState
    permission: PermissionStoreState
    engagements: EngagementStoreState
    risks: RiskStoreState
}

export interface CurrentResourceInitialization {
    orgId? : number
    engagementId?: number
    riskId?: number
}

export const RootStoreOptions : StoreOptions<RootState> = {
    strict: true,
    modules: {
        user: UserStoreModule,
        errors: ErrorStoreModule,
        appLayout: AppLayoutStoreModule,
        org : OrgStoreModule,
        permission: PermissionStoreModule,
        engagements: EngagementStoreModule,
        risks: RiskStoreModule,
    },
    actions: {
        initialize(context) {
            context.dispatch('user/initializeUserStore')
            context.commit('appLayout/initializeAppLayoutStore')
        },
        initializeCurrentResource(context, params : CurrentResourceInitialization) {
            if (!!params.orgId) {
                context.dispatch('org/initializeOrgStore', params.orgId)
            } else {
                context.commit('org/setRawOrg', null)
            }

            if (!!params.engagementId) {
                context.dispatch('engagements/initializeEngagementStore', {
                    orgId: params.orgId,
                    engId: params.engagementId,
                })
            } else {
                context.commit('engagements/setRawEngagement', null)
            }

            if (!!params.riskId) {
                context.dispatch('risks/initializeRiskStore', {
                    orgId: params.orgId,
                    engId: params.engagementId,
                    riskId: params.riskId,
                })
            } else {
                context.commit('risks/setRawRisk', null)
            }
        }
    },
    getters: {
        isFinishedLoading(_, getters) : boolean {
            return getters['user/isFinishedLoading']
        }
    }
}
