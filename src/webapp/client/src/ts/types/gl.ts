export enum GLAccountType {
	GLATAsset       = 1,
	GLATLiability   = 2,
	GLATEquity      = 3,
	GLATRevenue     = 4,
	GLATExpense     = 5,
	GLATContra      = 6,
}

export interface RawGLAccount {
    Id                    : number
    EngagementId          : number
    Name                  : string
    Description           : string
    AccountId             : string
    AccountType           : number
    ParentAccountId       : number | null
    FinanciallyRelevant   : boolean
    GlId                  : number
}

export function createEmptyGLAccount() : RawGLAccount {
    return {
        Id: -1,
        EngagementId: -1,
        Name: '',
        Description: '',
        AccountId: '',
        AccountType: GLAccountType.GLATAsset,
        ParentAccountId: null,
        FinanciallyRelevant: false,
        GlId: -1,
    }
}

export function accountTypeToString(t : GLAccountType) : string {
    switch (t) {
        case GLAccountType.GLATAsset:
            return 'Asset'
        case GLAccountType.GLATLiability:
            return 'Liability'
        case GLAccountType.GLATEquity:
            return 'Equity'
        case GLAccountType.GLATRevenue:
            return 'Revenue'
        case GLAccountType.GLATExpense:
            return 'Expense'
        case GLAccountType.GLATContra:
            return 'Contra'
    }
}

// This is structure is meant to represent a tree of GL accounts.
// We assume that there is always a singular root node. This root node
// has no account associated to it but is just meant to allow us to
// only keep track of 1 tree instead of multiple trees (i.e. one for each account
// without a parent account).
export class GLAccountTree {
    accList : RawGLAccount[]

    nodeAcc : RawGLAccount | null = null
    childNodes : GLAccountTree[] = []

    constructor(accs : RawGLAccount[], node : RawGLAccount | null) {
        this.nodeAcc = (!!node && node) || null
        this.accList = accs
        this.buildTree()
    }

    buildTree() {
        // The nodes we want to process here and make this node's children
        // are all nodes that have ParentAccountId set to this node's account's ID.
        // In the case where this node is the root node, ParentAccountId should be null.
        let filterCurrentLevel = (acc : RawGLAccount) =>
            (!this.nodeAcc && !acc.ParentAccountId) || (!!this.nodeAcc && acc.ParentAccountId === this.nodeAcc.Id)

        let currentLevelNodes = this.accList.filter(filterCurrentLevel)

        // We need to pass every single other node to each children since we don't have a good way
        // to determine whether a node is a child of another node at this point
        let restOfNodes = this.accList.filter((acc : RawGLAccount) => !filterCurrentLevel(acc))

        for (let node of currentLevelNodes) {
            let subTree = new GLAccountTree(restOfNodes, node)
            this.childNodes.push(subTree)
        }
    }
}

