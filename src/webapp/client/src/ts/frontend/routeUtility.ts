import { ResourceId } from '@client/ts/types/resource'

export function createResourceIdFromRouteParams(params : any) : ResourceId {
    let id : ResourceId = {
        orgId: Number(params.orgId),
        engagementId: Number(params.engId),
    }

    if (!!params.riskId) {
        id.riskId = Number(params.riskId)
    } else if (!!params.controlId) {
        id.controlId = Number(params.controlId)
    }

    return id
}
