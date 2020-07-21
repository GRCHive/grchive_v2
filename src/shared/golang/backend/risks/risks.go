package risks

import (
	"github.com/jmoiron/sqlx"
)

type Risk struct {
	Id                    int64  `db:"id"`
	EngagementId          int64  `db:"engagement_id"`
	Name                  string `db:"name"`
	HumanId               string `db:"human_id"`
	Severity              int16  `db:"severity"`
	SeverityJustification string `db:"severity_justification"`
	Description           string `db:"description"`
}

type RiskManager struct {
	db *sqlx.DB
}

func CreateRiskManager(db *sqlx.DB) *RiskManager {
	return &RiskManager{
		db: db,
	}
}
