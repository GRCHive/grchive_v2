import { ApiHttpHandler } from '@client/ts/api/handler'
import {
    RawVendor,
    RawVendorProduct,
} from '@client/ts/types/vendors'

export class VendorApiClient {
    handler : ApiHttpHandler
    constructor(handler : ApiHttpHandler) {
        this.handler = handler
    }

    listVendors(orgId : number, engagementId : number): Promise<RawVendor[]> {
        return this.handler.get<RawVendor[]>(`/orgs/${orgId}/engagements/${engagementId}/vendors`, {})
    }

    createVendor(orgId : number, vendor: RawVendor) : Promise<RawVendor> {
        return this.handler.post<RawVendor>(`/orgs/${orgId}/engagements/${vendor.EngagementId}/vendors`, {json : vendor})
    }

    getVendor(orgId : number, engagementId : number, vendorId : number) : Promise<RawVendor> {
        return this.handler.get<RawVendor>(`/orgs/${orgId}/engagements/${engagementId}/vendors/${vendorId}`, {})
    }

    updateVendor(orgId : number, vendor : RawVendor) : Promise<RawVendor> {
        return this.handler.put<RawVendor>(`/orgs/${orgId}/engagements/${vendor.EngagementId}/vendors/${vendor.Id}`, {json : vendor})
    }

    deleteVendor(orgId : number, engagementId : number, vendorId : number) : Promise<void> {
        return this.handler.delete<void>(`/orgs/${orgId}/engagements/${engagementId}/vendors/${vendorId}`, {})
    }

    listVendorProducts(orgId : number, engagementId : number, vendorId : number): Promise<RawVendorProduct[]> {
        return this.handler.get<RawVendorProduct[]>(`/orgs/${orgId}/engagements/${engagementId}/vendors/${vendorId}/products`, {})
    }

    createVendorProduct(orgId : number, engagementId : number, vendor: RawVendorProduct) : Promise<RawVendorProduct> {
        return this.handler.post<RawVendorProduct>(`/orgs/${orgId}/engagements/${engagementId}/vendors/${vendor.VendorId}/products`, {json : vendor})
    }

    updateVendorProduct(orgId : number, engagementId : number, vendor : RawVendorProduct) : Promise<RawVendorProduct> {
        return this.handler.put<RawVendorProduct>(`/orgs/${orgId}/engagements/${engagementId}/vendors/${vendor.VendorId}/products/${vendor.Id}`, {json : vendor})
    }

    getVendorProduct(orgId : number, engagementId : number, vendorId : number, pid : number) : Promise<RawVendorProduct> {
        return this.handler.get<RawVendorProduct>(`/orgs/${orgId}/engagements/${engagementId}/vendors/${vendorId}/products/${pid}`, {})
    }

    deleteVendorProduct(orgId : number, engagementId : number, vendorId : number, pid : number) : Promise<void> {
        return this.handler.delete<void>(`/orgs/${orgId}/engagements/${engagementId}/vendors/${vendorId}/products/${pid}`, {})
    }
}
