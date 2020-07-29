import { ResourceId } from '@client/ts/types/resource'
import { createBaseApiUrl } from '@client/ts/api/inventory'
import {
    InventoryType
} from '@client/ts/types/inventory'

export function resourceIdToApiUrl(id : ResourceId) : string {
    const baseUrl : string = `/orgs/${id.orgId}/engagements/${id.engagementId}`

    let apiUrl : string =''
    if (!!id.riskId) {
        apiUrl = `${baseUrl}/risks/${id.riskId}`
    } else if (!!id.controlId) {
        apiUrl = `${baseUrl}/controls/${id.controlId}`
    } else if (!!id.glAccountId) {
        if (id.glAccountId === -1) {
            apiUrl = `${baseUrl}/gl`
        } else {
            apiUrl = `${baseUrl}/gl/accs/${id.glAccountId}`
        }
    } else if (!!id.vendorId) {
        apiUrl = `${baseUrl}/vendors/${id.vendorId}`
    } else if (!!id.serverId) {
        apiUrl = `${createBaseApiUrl(InventoryType.ITServer, id.orgId, id.engagementId)}/${id.serverId}`
    } else if (!!id.desktopId) {
        apiUrl = `${createBaseApiUrl(InventoryType.ITDesktop, id.orgId, id.engagementId)}/${id.desktopId}`
    } else if (!!id.laptopId) {
        apiUrl = `${createBaseApiUrl(InventoryType.ITLaptop, id.orgId, id.engagementId)}/${id.laptopId}`
    } else if (!!id.mobileId) {
        apiUrl = `${createBaseApiUrl(InventoryType.ITMobile, id.orgId, id.engagementId)}/${id.mobileId}`
    } else if (!!id.embeddedId) {
        apiUrl = `${createBaseApiUrl(InventoryType.ITEmbedded, id.orgId, id.engagementId)}/${id.embeddedId}`
    } else if (!!id.databaseId) {
        apiUrl = `${baseUrl}/databases/${id.databaseId}`
    } else if (!!id.systemId) {
        apiUrl = `${baseUrl}/systems/${id.systemId}`
    }

    return apiUrl
}
