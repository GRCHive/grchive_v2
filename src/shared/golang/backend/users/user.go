package users

import (
	"github.com/jmoiron/sqlx"
	"gitlab.com/grchive/grchive-v2/shared/backend/utility"
)

type UserManager struct {
	db *sqlx.DB
}

func CreateUserManager(db *sqlx.DB) *UserManager {
	return &UserManager{
		db: db,
	}
}

func (m *UserManager) WrapDatabaseTx(fns ...utility.TxHandler) error {
	return utility.WrapDatabaseTx(m.db, fns...)
}

type User struct {
	Id               int64  `db:"id"`
	FullName         string `db:"full_name"`
	Email            string `db:"email"`
	FusionAuthUserId string `db:"fa_user_id"`
}
