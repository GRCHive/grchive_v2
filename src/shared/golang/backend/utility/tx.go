package utility

import (
	"github.com/jmoiron/sqlx"
)

type TxHandler func(tx *sqlx.Tx) error

func WrapDatabaseTx(db *sqlx.DB, fns ...TxHandler) error {
	tx, err := db.Beginx()
	if err != nil {
		return err
	}

	for _, fn := range fns {
		err := fn(tx)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}
