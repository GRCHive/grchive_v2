import { ApiHttpHandler } from '@client/ts/api/handler'
import {
    RawGLAccount,
} from '@client/ts/types/gl'

export class GeneralLedgerApiClient {
    handler : ApiHttpHandler
    constructor(handler : ApiHttpHandler) {
        this.handler = handler
    }

    listAccounts(orgId : number, engagementId : number): Promise<RawGLAccount[]> {
        return this.handler.get<RawGLAccount[]>(`/orgs/${orgId}/engagements/${engagementId}/gl/accs`, {})
    }

    createAccount(orgId : number, acc : RawGLAccount) : Promise<RawGLAccount> {
        return this.handler.post<RawGLAccount>(`/orgs/${orgId}/engagements/${acc.EngagementId}/gl/accs`, {json: acc})
    }

    updateAccount(orgId : number, acc : RawGLAccount) : Promise<RawGLAccount> {
        return this.handler.put<RawGLAccount>(`/orgs/${orgId}/engagements/${acc.EngagementId}/gl/accs/${acc.Id}`, {json: acc})
    }

    getAccount(orgId : number, engagementId : number, accId : number) : Promise<RawGLAccount> {
        return this.handler.get<RawGLAccount>(`/orgs/${orgId}/engagements/${engagementId}/gl/accs/${accId}`, {})
    }

    deleteAccount(orgId : number, engagementId : number, accId : number) : Promise<void> {
        return this.handler.delete<void>(`/orgs/${orgId}/engagements/${engagementId}/gl/accs/${accId}`, {})
    }
}
