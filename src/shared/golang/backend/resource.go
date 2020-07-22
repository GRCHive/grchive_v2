package backend

import (
	"errors"
	"fmt"
	"strconv"
)

type ResourceIdentifier int64

func (r ResourceIdentifier) Str() string {
	return fmt.Sprintf("resource-%d", r)
}

const (
	RIUser ResourceIdentifier = iota
	RIOrganization
	RIEngagement
	RIRisk
)

func (b *BackendInterface) GetResource(id ResourceIdentifier, key string) (interface{}, error) {
	switch id {
	case RIOrganization:
		orgId, err := strconv.ParseInt(key, 10, 64)
		if err != nil {
			return nil, err
		}
		return b.Orgs.GetOrgFromId(orgId)
	case RIEngagement:
		engId, err := strconv.ParseInt(key, 10, 64)
		if err != nil {
			return nil, err
		}
		return b.Engagements.GetEngagementFromId(engId)
	case RIRisk:
		riskId, err := strconv.ParseInt(key, 10, 64)
		if err != nil {
			return nil, err
		}
		return b.Risks.GetRiskFromId(riskId)
	case RIUser:
		userId, err := strconv.ParseInt(key, 10, 64)
		if err != nil {
			return nil, err
		}
		return b.Users.MustGetUserFromId(userId)
	default:
		return nil, errors.New("Unsupported resource identifier.")
	}
}
