import { ApiHttpHandler } from '@client/ts/api/handler'
import { VirtualMachine } from '@client/ts/types/vm'

function createBaseApiUrl(orgId : number, engagementId : number) : string {
    return `/orgs/${orgId}/engagements/${engagementId}/vm`
}

function createSingleBaseApiUrl(orgId : number, engagementId : number, vmId : number) : string {
    return `${createBaseApiUrl(orgId, engagementId)}/${vmId}`
}

export class VMApiClient {
    handler : ApiHttpHandler
    constructor(handler : ApiHttpHandler) {
        this.handler = handler
    }

    listVM(orgId : number, engagementId : number): Promise<VirtualMachine[]> {
        return this.handler.get<VirtualMachine[]>(createBaseApiUrl(orgId, engagementId), {})
    }

    createVM(orgId: number, vm : VirtualMachine) : Promise<VirtualMachine> {
        return this.handler.post<VirtualMachine>(createBaseApiUrl(orgId, vm.EngagementId), {json: vm})
    }

    getVM(orgId : number, engagementId : number, vmId : number) : Promise<VirtualMachine> {
        return this.handler.get<VirtualMachine>(createSingleBaseApiUrl(orgId, engagementId, vmId), {})
    }

    deleteVM(orgId : number, engagementId : number, vmId : number) : Promise<void> {
        return this.handler.delete<void>(createSingleBaseApiUrl(orgId, engagementId, vmId), {})
    }

    updateVM(orgId : number, vm : VirtualMachine) : Promise<VirtualMachine> {
        return this.handler.put<VirtualMachine>(createSingleBaseApiUrl(orgId, vm.EngagementId, vm.Id), {json : vm})
    }
}
