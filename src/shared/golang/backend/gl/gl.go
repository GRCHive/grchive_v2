package gl

import (
	"github.com/jmoiron/sqlx"
)

type GLAccountType int32

const (
	GLATAsset     GLAccountType = 1
	GLATLiability               = 2
	GLATEquity                  = 3
	GLATRevenue                 = 4
	GLATExpense                 = 5
	GLATContra                  = 6
)

type GLAccount struct {
	Id                  int64         `db:"id"`
	EngagementId        int64         `db:"engagement_id"`
	Name                string        `db:"name"`
	Description         string        `db:"description"`
	AccountId           string        `db:"account_id"`
	AccountType         GLAccountType `db:"account_type"`
	ParentAccountId     *int64        `db:"parent_account_id"`
	FinanciallyRelevant bool          `db:"financially_relevant"`
}

type GLManager struct {
	db *sqlx.DB
}

func CreateGLManager(db *sqlx.DB) *GLManager {
	return &GLManager{
		db: db,
	}
}
