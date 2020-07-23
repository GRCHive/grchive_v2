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
    }
}
