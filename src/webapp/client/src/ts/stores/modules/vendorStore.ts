import { Module } from 'vuex'
import { RawVendor, RawVendorProduct } from '@client/ts/types/vendors'
import { RootState } from '@client/ts/stores/store'
import { GrchiveApi, ErrorHandler } from '@client/ts/main'

export interface VendorStoreState {
    rawVendor : RawVendor | null
    rawVendorProduct: RawVendorProduct | null
}

export const VendorStoreModule : Module<VendorStoreState, RootState> = {
    namespaced: true,
    state: () => ({
        rawVendor: null,
        rawVendorProduct: null,
    }),
    mutations: {
        setRawVendor(state : VendorStoreState, rawVendor : RawVendor | null) {
            state.rawVendor = rawVendor
        },
        setRawVendorProduct(state : VendorStoreState, rawVendorProduct : RawVendorProduct | null) {
            state.rawVendorProduct = rawVendorProduct
        }
    },
    actions: {
        initializeVendorStore(context, { orgId , engId, vendorId }) {
            GrchiveApi.vendors.getVendor(orgId, engId, vendorId).then((resp : RawVendor) => {
                context.commit('setRawVendor', resp)
            }).catch((err : any) => {
                ErrorHandler.failurePageOnError(err)
            })
        },
        initializeVendorProductStore(context, { orgId , engId, vendorId, vendorProductId }) {
            GrchiveApi.vendors.getVendorProduct(orgId, engId, vendorId, vendorProductId).then((resp : RawVendorProduct) => {
                context.commit('setRawVendorProduct', resp)
            }).catch((err : any) => {
                ErrorHandler.failurePageOnError(err)
            })
        }
    },
}
