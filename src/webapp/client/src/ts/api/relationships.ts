import { ApiHttpHandler } from '@client/ts/api/handler'
import { ResourceId, ResourceType } from '@client/ts/types/resource'
import { resourceIdToApiUrl } from '@client/ts/api/url'
import { RelationshipWrapper } from '@client/ts/types/relationships'

function createApiUrl(id : ResourceId, typ : ResourceType) : string {
    let baseUrl : string = resourceIdToApiUrl(id)
    baseUrl += '/relationships'
    if (typ == ResourceType.RTControl) {
        baseUrl += '/controls'
    } else if (typ == ResourceType.RTRisk) {
        baseUrl += '/risks'
    }
    return baseUrl
}

export class RelationshipApiClient {
    handler : ApiHttpHandler
    constructor(handler : ApiHttpHandler) {
        this.handler = handler
    }

    listRelationships<T>(id : ResourceId, typ : ResourceType) : Promise<RelationshipWrapper<T>[]> {
        return this.handler.get<RelationshipWrapper<T>[]>(createApiUrl(id, typ), {})
    }

    createRelationship<T>(from: ResourceId, typ : ResourceType,  to: ResourceId) : Promise<RelationshipWrapper<T>> {
        return this.handler.post<RelationshipWrapper<T>>(createApiUrl(from, typ), {json : to})
    }

    deleteRelationship(from : ResourceId, typ : ResourceType, toId : number) : Promise<void> {
        return this.handler.delete<void>(createApiUrl(from, typ) + `/${toId}`, {})
    }
}
