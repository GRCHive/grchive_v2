import Vue from 'vue'
import { Module } from 'vuex'
import { RootState } from '@client/ts/stores/store'
import { GrchiveApi } from '@client/ts/main'
import { Permission } from '@client/ts/types/roles'

export interface PermissionStoreState {
    orgPermissionMaps: Record<number, Record<number, Record<Permission, boolean>>>
    inProgressMaps: Record<number, Record<number, Record<Permission, boolean>>>
}

export const PermissionStoreModule : Module<PermissionStoreState, RootState> = {
    namespaced: true,
    state: () => ({
        orgPermissionMaps: Object(),
        inProgressMaps: Object()
    }),
    mutations: {
        setHasPermissions(state : PermissionStoreState, { orgId, engagementId, permissions, val }) {
            if (!(orgId in state.orgPermissionMaps)) {
                Vue.set(state.orgPermissionMaps, orgId, new Object())
            }

            let orgMap = state.orgPermissionMaps[orgId]
            if (!(engagementId in orgMap)) {
                Vue.set(orgMap, engagementId, new Object())
            }

            let engMap = orgMap[engagementId]
            for (let p of permissions) {
                Vue.set(engMap, <Permission>p, val)
            }
        },
        setPermissionLoading(state : PermissionStoreState, { orgId, engagementId, permissions }) {
            if (!(orgId in state.orgPermissionMaps)) {
                Vue.set(state.inProgressMaps, orgId, new Object())
            }

            let orgMap = state.inProgressMaps[orgId]
            if (!(engagementId in orgMap)) {
                Vue.set(orgMap, engagementId, new Object())
            }

            let engMap = orgMap[engagementId]
            for (let p of permissions) {
                Vue.set(engMap, <Permission>p, true)
            }
        }
    },
    actions: {
        initializeHasPermissions(context, { orgId, engagementId, permissions }) {
            let toLoad = context.getters['getPermissionsToLoad'](orgId, engagementId, permissions)
            context.commit('setPermissionLoading', {orgId, engagementId, permissions: toLoad})

            if (toLoad.length == 0) {
                return
            }

            GrchiveApi.user.checkCurrentUserPermissions(orgId, engagementId, toLoad).then((resp : boolean | null) => {
                context.commit('setHasPermissions', { orgId, engagementId, permissions: toLoad, val: resp })
            })
        }
    },
    getters: {
        hasValue(state : PermissionStoreState) : (orgId : number, engagementId : number, p : Permission) => boolean {
            return (orgId : number, engagementId : number, p : Permission) : boolean => {
                if (!(orgId in state.orgPermissionMaps)) {
                    return false
                }

                let orgMap = state.orgPermissionMaps[orgId]
                if (!(engagementId in orgMap)) {
                    return false
                }

                let engMap = orgMap[engagementId]
                return p in engMap
            }
        },

        isLoading(state : PermissionStoreState) : (orgId : number, engagementId: number, p : Permission) => boolean {
            return (orgId : number, engagementId : number, p : Permission) : boolean => {
                if (!(orgId in state.inProgressMaps)) {
                    return false
                }

                let orgMap = state.inProgressMaps[orgId]
                if (!(engagementId in orgMap)) {
                    return false
                }
                let engMap = orgMap[engagementId]
                return p in engMap
            }
        },

        getPermissionsToLoad(_ : PermissionStoreState, getters) : (orgId : number, engagementId : number, permissions : Permission[]) => Permission[]{
            return (orgId : number, engagementId : number, permissions: Permission[]) : Permission[] => {
                return permissions.filter((p : Permission) => !getters.hasValue(orgId, engagementId, p) && !getters.isLoading(orgId, p))
            }
        },
        currentUserHasPermissions(state : PermissionStoreState) : (orgId : number, engagementId : number, permissions : Permission[]) => boolean | null {
            return (orgId : number, engagementId : number, permissions: Permission[]) : boolean | null => {
                if (!(orgId in state.orgPermissionMaps)) {
                    return null
                }

                let orgMap = state.orgPermissionMaps[orgId]
                if (!(engagementId in orgMap)) {
                    return null
                }

                let engMap = orgMap[engagementId]
                let hasPermission : boolean = true
                for (let p of permissions) {
                    if (!(p in engMap)) {
                        return null
                    }
                    hasPermission = hasPermission && engMap[p]
                }
                return hasPermission
            }
        }
    }
}
