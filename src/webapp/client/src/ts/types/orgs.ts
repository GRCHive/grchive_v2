export interface RawOrganization {
    Id          : number
    Name        : string
    Description : string
    ParentOrgId : number | null
    OwnerUserId : number
}

export function createEmptyOrg() : RawOrganization {
    return {
        Id: -1,
        Name: '',
        Description: '',
        ParentOrgId: null,
        OwnerUserId: -1
    }
}

// This is structure is meant to represent a tree of orgs.
// We assume that there is always a singular root node. This root node
// has no organization associated to it but is just meant to allow us to
// only keep track of 1 tree instead of multiple trees (i.e. one for each org
// without a parent org).
export class OrgTree {
    orgList : RawOrganization[]

    nodeOrg : RawOrganization | null = null
    childNodes : OrgTree[] = []

    constructor(orgs : RawOrganization[], node? : RawOrganization) {
        this.nodeOrg = (!!node && node) || null
        this.orgList = orgs
        this.buildTree()
    }

    buildTree() {
        // The nodes we want to process here and make this node's children
        // are all nodes that have ParentOrgId set to this node's org's ID.
        // In the case where this node is the root node, ParentOrgId should be null.
        let filterCurrentLevel = (org : RawOrganization) =>
            (!this.nodeOrg && !org.ParentOrgId) || (!!this.nodeOrg && org.ParentOrgId === this.nodeOrg.Id)

        let currentLevelNodes = this.orgList.filter(filterCurrentLevel)

        // We need to pass every single other node to each children since we don't have a good way
        // to determine whether a node is a child of another node at this point
        let restOfNodes = this.orgList.filter((org : RawOrganization) => !filterCurrentLevel(org))

        for (let node of currentLevelNodes) {
            let subTree = new OrgTree(restOfNodes, node)
            this.childNodes.push(subTree)
        }
    }
}
