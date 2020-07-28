import { Module } from 'vuex'
import { RawDatabase } from '@client/ts/types/databases'
import { RootState } from '@client/ts/stores/store'
import { GrchiveApi, ErrorHandler } from '@client/ts/main'

export interface DatabaseStoreState {
    rawDatabase : RawDatabase | null
}

export const DatabaseStoreModule : Module<DatabaseStoreState, RootState> = {
    namespaced: true,
    state: () => ({
        rawDatabase: null,
    }),
    mutations: {
        setRawDatabase(state : DatabaseStoreState, rawDatabase : RawDatabase | null) {
            state.rawDatabase = rawDatabase
        }
    },
    actions: {
        initializeDatabaseStore(context, { orgId , engId, databaseId }) {
            GrchiveApi.databases.getDatabase(orgId, engId, databaseId).then((resp : RawDatabase) => {
                context.commit('setRawDatabase', resp)
            }).catch((err : any) => {
                ErrorHandler.failurePageOnError(err)
            })
        }
    },
}
