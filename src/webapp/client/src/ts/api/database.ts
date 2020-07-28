import { ApiHttpHandler } from '@client/ts/api/handler'
import { RawDatabase } from '@client/ts/types/databases'

function createBaseApiUrl(orgId : number, engagementId : number) : string {
    return `/orgs/${orgId}/engagements/${engagementId}/databases`
}

function createSingleBaseApiUrl(orgId : number, engagementId : number, dbId : number) : string {
    return `${createBaseApiUrl(orgId, engagementId)}/${dbId}`
}

export class DatabaseApiClient {
    handler : ApiHttpHandler
    constructor(handler : ApiHttpHandler) {
        this.handler = handler
    }

    listDatabases(orgId : number, engagementId : number): Promise<RawDatabase[]> {
        return this.handler.get<RawDatabase[]>(createBaseApiUrl(orgId, engagementId), {})
    }

    createDatabase(orgId: number, db : RawDatabase) : Promise<RawDatabase> {
        return this.handler.post<RawDatabase>(createBaseApiUrl(orgId, db.EngagementId), {json: db})
    }

    getDatabase(orgId : number, engagementId : number, dbId : number) : Promise<RawDatabase> {
        return this.handler.get<RawDatabase>(createSingleBaseApiUrl(orgId, engagementId, dbId), {})
    }

    deleteDatabase(orgId : number, engagementId : number, dbId : number) : Promise<void> {
        return this.handler.delete<void>(createSingleBaseApiUrl(orgId, engagementId, dbId), {})
    }

    updateDatabase(orgId : number, db : RawDatabase) : Promise<RawDatabase> {
        return this.handler.put<RawDatabase>(createSingleBaseApiUrl(orgId, db.EngagementId, db.Id), {json : db})
    }
}
