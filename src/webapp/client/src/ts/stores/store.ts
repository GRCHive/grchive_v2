import { StoreOptions } from 'vuex'

import { UserStoreModule, UserStoreState } from '@client/ts/stores/modules/userStore'
import { PermissionStoreModule, PermissionStoreState } from '@client/ts/stores/modules/permissionStore'
import { AppLayoutStoreModule, AppLayoutStoreState } from '@client/ts/stores/modules/appLayoutStore'
import { OrgStoreModule, OrgStoreState } from '@client/ts/stores/modules/orgStore'
import { EngagementStoreModule, EngagementStoreState } from '@client/ts/stores/modules/engagementStore'
import { RiskStoreModule, RiskStoreState } from '@client/ts/stores/modules/riskStore'
import { GeneralLedgerStoreModule, GeneralLedgerStoreState } from '@client/ts/stores/modules/glStore'
import { ControlStoreModule, ControlStoreState } from '@client/ts/stores/modules/controlStore'
import { ErrorStoreModule, ErrorStoreState } from '@client/ts/stores/modules/errorStore'

export interface RootState {
    user: UserStoreState
    appLayout: AppLayoutStoreState
    org: OrgStoreState
    permission: PermissionStoreState
    engagements: EngagementStoreState
    risks: RiskStoreState
    controls: ControlStoreState
    gl : GeneralLedgerStoreState
    errors : ErrorStoreState
}

export interface CurrentResourceInitialization {
    orgId? : number
    engagementId?: number
    riskId?: number
    controlId?: number
    glAccountId? : number
}

export const RootStoreOptions : StoreOptions<RootState> = {
    strict: true,
    modules: {
        user: UserStoreModule,
        appLayout: AppLayoutStoreModule,
        org : OrgStoreModule,
        permission: PermissionStoreModule,
        engagements: EngagementStoreModule,
        risks: RiskStoreModule,
        controls: ControlStoreModule,
        gl: GeneralLedgerStoreModule,
        errors: ErrorStoreModule
    },
    actions: {
        initialize(context) {
            context.dispatch('user/initializeUserStore')
            context.commit('appLayout/initializeAppLayoutStore')
        },
        initializeCurrentResource(context, params : CurrentResourceInitialization) {
            if (!!params.orgId) {
                if (!context.state.org.rawOrg || context.state.org.rawOrg.Id != params.orgId) {
                    context.dispatch('org/initializeOrgStore', params.orgId)
                }
            } else {
                context.commit('org/setRawOrg', null)
            }

            if (!!params.engagementId) {
                if (!context.state.engagements.rawEngagement || context.state.engagements.rawEngagement.Id != params.engagementId) {
                    context.dispatch('engagements/initializeEngagementStore', {
                        orgId: params.orgId,
                        engId: params.engagementId,
                    })
                }
            } else {
                context.commit('engagements/setRawEngagement', null)
            }

            if (!!params.riskId) {
                if (!context.state.risks.rawRisk || context.state.risks.rawRisk.Id != params.riskId) {
                    context.dispatch('risks/initializeRiskStore', {
                        orgId: params.orgId,
                        engId: params.engagementId,
                        riskId: params.riskId,
                    })
                }
            } else {
                context.commit('risks/setRawRisk', null)
            }

            if (!!params.controlId) {
                if (!context.state.controls.rawControl || context.state.controls.rawControl.Id != params.controlId) {
                    context.dispatch('controls/initializeControlStore', {
                        orgId: params.orgId,
                        engId: params.engagementId,
                        controlId: params.controlId,
                    })
                }
            } else {
                context.commit('controls/setRawControl', null)
            }

            if (!!params.glAccountId) {
                if (!context.state.gl.rawGLAccount || context.state.gl.rawGLAccount.Id != params.glAccountId) {
                    context.dispatch('gl/initializeGeneralLedgerStore', {
                        orgId: params.orgId,
                        engId: params.engagementId,
                        glAccountId: params.glAccountId,
                    })
                }
            } else {
                context.commit('gl/setRawGeneralLedger', null)
            }

        }
    },
    getters: {
        isFinishedLoading(_, getters) : boolean {
            return getters['user/isFinishedLoading']
        }
    }
}
