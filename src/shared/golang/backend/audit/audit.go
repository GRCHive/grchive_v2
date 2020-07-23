package audit

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"gitlab.com/grchive/grchive-v2/shared/backend/utility"
)

type AuditTrailId struct {
	OrgId        *int64
	EngagementId *int64
	ThreadId     *int64
	UserId       *int64
	IpAddress    *string
}

func WrapDatabaseAuditTx(db *sqlx.DB, id AuditTrailId, fns ...utility.TxHandler) error {
	// Do the SET LOCAL commands as the VERY last thing so that we can dynamically
	// get OrgId/EngagementId even for creating the org or engagement. The pointers
	// in that case will be set to the struct's value and the pointer passed here. So by the time
	// we get to the commands in this file, the pointer will have the correct ID value and we
	// can store the ID value into the audit log. The audit log triggers are deferred so will only
	// fire at the very end of the transaction.
	return utility.WrapDatabaseTx(db, append(fns, func(tx *sqlx.Tx) error {
		orgId := int64(-1)
		if id.OrgId != nil {
			orgId = *id.OrgId
		}
		_, err := tx.Exec(fmt.Sprintf("SET LOCAL grchive.current_org_id TO %d", orgId))
		return err
	}, func(tx *sqlx.Tx) error {
		engId := int64(-1)
		if id.EngagementId != nil {
			engId = *id.EngagementId
		}
		_, err := tx.Exec(fmt.Sprintf("SET LOCAL grchive.current_engagement_id TO %d", engId))
		return err
	}, func(tx *sqlx.Tx) error {
		threadId := int64(-1)
		if id.ThreadId != nil {
			threadId = *id.ThreadId
		}
		_, err := tx.Exec(fmt.Sprintf("SET LOCAL grchive.current_thread_id TO %d", threadId))
		return err
	}, func(tx *sqlx.Tx) error {
		userId := int64(-1)
		if id.UserId != nil {
			userId = *id.UserId
		}
		_, err := tx.Exec(fmt.Sprintf("SET LOCAL grchive.current_user_id TO %d", userId))
		return err
	}, func(tx *sqlx.Tx) error {
		if id.IpAddress != nil {
			_, err := tx.Exec(fmt.Sprintf("SET LOCAL grchive.ip_address TO '%s'", *id.IpAddress))
			return err
		} else {
			_, err := tx.Exec("SELECT set_config('grchive.ip_address', text(inet_client_addr()), true)")
			return err
		}
	})...)
}

func CreateNilSystemAuditId() AuditTrailId {
	return AuditTrailId{
		OrgId:        nil,
		EngagementId: nil,
		ThreadId:     nil,
		UserId:       nil,
		IpAddress:    nil,
	}
}
