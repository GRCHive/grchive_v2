import { ApiHttpHandler } from '@client/ts/api/handler'
import {
    RawBaseInventory,
    InventoryType
} from '@client/ts/types/inventory'

export function createBaseApiUrl(typ : InventoryType, orgId : number, engagementId : number) : string {
    let suffix = ''
    switch (typ) {
        case InventoryType.ITServer:
            suffix = 'servers'
            break
        case InventoryType.ITDesktop:
            suffix = 'desktops'
            break
        case InventoryType.ITLaptop:
            suffix = 'laptops'
            break
        case InventoryType.ITMobileDevice:
            suffix = 'mobile'
            break
        case InventoryType.ITEmbeddedDevice:
            suffix = 'embedded'
            break
    }
    return `/orgs/${orgId}/engagements/${engagementId}/inventory/${suffix}`
}

export class InventoryApiClient {
    handler : ApiHttpHandler
    constructor(handler : ApiHttpHandler) {
        this.handler = handler
    }

    listInventory<T extends RawBaseInventory>(typ : InventoryType, orgId : number, engagementId : number): Promise<T[]> {
        return this.handler.get<T[]>(createBaseApiUrl(typ, orgId, engagementId), {})
    }

    createInventory<T extends RawBaseInventory>(typ : InventoryType, orgId : number, inv : T): Promise<T> {
        return this.handler.post<T>(createBaseApiUrl(typ, orgId, inv.Inventory.EngagementId), {json : inv})
    }

    updateInventory<T extends RawBaseInventory>(typ : InventoryType, orgId : number, inv : T): Promise<T> {
        return this.handler.put<T>(createBaseApiUrl(typ, orgId, inv.Inventory.EngagementId)+`/${inv.Id}`, {json : inv})
    }

    getInventory<T extends RawBaseInventory>(typ : InventoryType, orgId : number, engagementId: number, invId : number): Promise<T> {
        return this.handler.get<T>(createBaseApiUrl(typ, orgId, engagementId)+`/${invId}`, {})
    }

    deleteInventory(typ : InventoryType, orgId : number, engagementId : number, invId : number): Promise<void> {
        return this.handler.delete<void>(createBaseApiUrl(typ, orgId, engagementId)+`/${invId}`, {})
    }
}
