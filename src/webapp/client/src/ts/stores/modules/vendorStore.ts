import { Module } from 'vuex'
import { RawVendor } from '@client/ts/types/vendors'
import { RootState } from '@client/ts/stores/store'
import { GrchiveApi, ErrorHandler } from '@client/ts/main'

export interface VendorStoreState {
    rawVendor : RawVendor | null
}

export const VendorStoreModule : Module<VendorStoreState, RootState> = {
    namespaced: true,
    state: () => ({
        rawVendor: null,
    }),
    mutations: {
        setRawVendor(state : VendorStoreState, rawVendor : RawVendor | null) {
            state.rawVendor = rawVendor
        }
    },
    actions: {
        initializeVendorStore(context, { orgId , engId, vendorId }) {
            GrchiveApi.vendors.getVendor(orgId, engId, vendorId).then((resp : RawVendor) => {
                context.commit('setRawVendor', resp)
            }).catch((err : any) => {
                ErrorHandler.failurePageOnError(err)
            })
        }
    },
}
