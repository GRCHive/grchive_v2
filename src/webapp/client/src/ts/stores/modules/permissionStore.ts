import Vue from 'vue'
import { Module } from 'vuex'
import { RootState } from '@client/ts/stores/store'
import { GrchiveApi } from '@client/ts/main'
import { Permission } from '@client/ts/types/roles'

export interface PermissionStoreState {
    orgPermissionMaps: Record<number, Record<Permission, boolean>>
    inProgressMaps: Record<number, Record<Permission, boolean>>
}

export const PermissionStoreModule : Module<PermissionStoreState, RootState> = {
    namespaced: true,
    state: () => ({
        orgPermissionMaps: Object(),
        inProgressMaps: Object()
    }),
    mutations: {
        setHasPermissions(state : PermissionStoreState, { orgId, permissions, val }) {
            if (!(orgId in state.orgPermissionMaps)) {
                Vue.set(state.orgPermissionMaps, orgId, new Object())
            }

            let orgMap : Record<Permission, boolean> = state.orgPermissionMaps[orgId]
            for (let p of permissions) {
                Vue.set(orgMap, <Permission>p, val)
            }
        },
        setPermissionLoading(state : PermissionStoreState, { orgId, permissions }) {
            if (!(orgId in state.orgPermissionMaps)) {
                Vue.set(state.inProgressMaps, orgId, new Object())
            }

            let orgMap : Record<Permission, boolean> = state.inProgressMaps[orgId]
            for (let p of permissions) {
                Vue.set(orgMap, <Permission>p, true)
            }
        }
    },
    actions: {
        initializeHasPermissions(context, { orgId, permissions }) {
            let toLoad = context.getters['getPermissionsToLoad'](orgId, permissions)
            context.commit('setPermissionLoading', {orgId, permissions: toLoad})

            if (toLoad.length == 0) {
                return
            }

            GrchiveApi.user.checkCurrentUserPermissions(orgId, toLoad).then((resp : boolean | null) => {
                context.commit('setHasPermissions', { orgId, permissions: toLoad, val: resp })
            })
        }
    },
    getters: {
        hasValue(state : PermissionStoreState) : (orgId : number, p : Permission) => boolean {
            return (orgId : number, p : Permission) : boolean => {
                if (!(orgId in state.orgPermissionMaps)) {
                    return false
                }

                let orgMap : Record<Permission, boolean> = state.orgPermissionMaps[orgId]
                return p in orgMap
            }
        },

        isLoading(state : PermissionStoreState) : (orgId : number, p : Permission) => boolean {
            return (orgId : number, p : Permission) : boolean => {
                if (!(orgId in state.inProgressMaps)) {
                    return false
                }

                let orgMap : Record<Permission, boolean> = state.inProgressMaps[orgId]
                return p in orgMap
            }
        },

        getPermissionsToLoad(_ : PermissionStoreState, getters) : (orgId : number, permissions : Permission[]) => Permission[]{
            return (orgId : number, permissions: Permission[]) : Permission[] => {
                return permissions.filter((p : Permission) => !getters.hasValue(orgId, p) && !getters.isLoading(orgId, p))
            }
        },
        currentUserHasPermissions(state : PermissionStoreState) : (orgId : number, permissions : Permission[]) => boolean | null {
            return (orgId : number, permissions: Permission[]) : boolean | null => {
                if (!(orgId in state.orgPermissionMaps)) {
                    return null
                }

                let orgMap : Record<Permission, boolean> = state.orgPermissionMaps[orgId]
                let hasPermission : boolean = false
                for (let p of permissions) {
                    if (!(p in orgMap)) {
                        return null
                    }
                    hasPermission = hasPermission || orgMap[p]
                }
                return hasPermission
            }
        }
    }
}
