import { ApiHttpHandler } from '@client/ts/api/handler'
import {
    RawControl,
} from '@client/ts/types/controls'

export class ControlApiClient {
    handler : ApiHttpHandler
    constructor(handler : ApiHttpHandler) {
        this.handler = handler
    }

    createControl(orgId : number, control : RawControl) : Promise<RawControl | null> {
        return this.handler.post<RawControl>(`/orgs/${orgId}/engagements/${control.EngagementId}/controls`, {json: control})
    }

    updateControl(orgId : number, control : RawControl) : Promise<RawControl | null> {
        return this.handler.put<RawControl>(`/orgs/${orgId}/engagements/${control.EngagementId}/controls/${control.Id}`, {json: control})
    }

    deleteControl(orgId : number, engagementId : number, controlId : number) : Promise<void | null> {
        return this.handler.delete<void>(`/orgs/${orgId}/engagements/${engagementId}/controls/${controlId}`, {})
    }

    getControl(orgId : number, engagementId: number, controlId : number) : Promise<RawControl | null> {
        return this.handler.get<RawControl>(`/orgs/${orgId}/engagements/${engagementId}/controls/${controlId}`, {})
    }

    listControls(orgId : number, engagementId : number): Promise<RawControl[] | null> {
        return this.handler.get<RawControl[]>(`/orgs/${orgId}/engagements/${engagementId}/controls`, {})
    }
}
