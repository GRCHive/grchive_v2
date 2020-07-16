package audit

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"gitlab.com/grchive/grchive-v2/shared/backend/utility"
)

type AuditTrailId struct {
	UserId    int64
	IpAddress string
}

func WrapDatabaseAuditTx(db *sqlx.DB, id *AuditTrailId, fns ...utility.TxHandler) error {
	return utility.WrapDatabaseTx(db, append([]utility.TxHandler{
		func(tx *sqlx.Tx) error {
			var err error
			var userIdStmt *sqlx.Stmt
			var ipAddrStmt *sqlx.Stmt

			if id != nil {
				userIdStmt, err = tx.Preparex(fmt.Sprintf("SET LOCAL grchive.current_user_id TO %d", id.UserId))
				if err != nil {
					return err
				}

				ipAddrStmt, err = tx.Preparex(fmt.Sprintf("SET LOCAL grchive.ip_address TO %s", id.IpAddress))
				if err != nil {
					return err
				}
			} else {
				userIdStmt, err = tx.Preparex("SET LOCAL grchive.current_user_id TO -1")
				if err != nil {
					return err
				}

				ipAddrStmt, err = tx.Preparex(`
					SELECT set_config('grchive.ip_address', text(inet_client_addr()), true);
				`)
				if err != nil {
					return err
				}
			}

			_, err = userIdStmt.Exec()
			if err != nil {
				return err
			}

			_, err = ipAddrStmt.Exec()
			if err != nil {
				return err
			}

			return nil
		},
	}, fns...)...)
}
