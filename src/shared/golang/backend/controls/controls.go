package controls

import (
	"github.com/jmoiron/sqlx"
)

type ControlFrequencyType int32
type ControlType int32

const (
	CFTAdHoc    ControlFrequencyType = 1
	CFTInterval                      = 2
	CFTOther                         = 3
)

const (
	CTAccess           ControlType = 1
	CTAuthorization                = 2
	CTConfiguration                = 3
	CTGitc                         = 4
	CTInterface                    = 5
	CTManagementReview             = 6
	CTReconciliation               = 7
	CTReport                       = 8
)

type Control struct {
	Id                      int64                `db:"id"`
	EngagementId            int64                `db:"engagement_id"`
	Name                    string               `db:"name"`
	HumanId                 string               `db:"human_id"`
	Likelihood              int16                `db:"likelihood"`
	LikelihoodJustification string               `db:"likelihood_justification"`
	ControlType             ControlType          `db:"control_type"`
	OwnerId                 *int64               `db:"owner_id"`
	IsManual                bool                 `db:"is_manual"`
	FreqType                ControlFrequencyType `db:"freq_type"`
	FreqOther               string               `db:"freq_other"`
	FreqInterval            string               `db:"freq_interval"`
	Description             string               `db:"description"`
}

type ControlManager struct {
	db *sqlx.DB
}

func CreateControlManager(db *sqlx.DB) *ControlManager {
	return &ControlManager{
		db: db,
	}
}
