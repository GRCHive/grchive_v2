package users

import (
	"github.com/jmoiron/sqlx"
)

type UserManager struct {
	db *sqlx.DB
}

func CreateUserManager(db *sqlx.DB) *UserManager {
	return &UserManager{
		db: db,
	}
}

type User struct {
	Id               int64  `db:"id"`
	FullName         string `db:"full_name"`
	Email            string `db:"email"`
	FusionAuthUserId string `db:"fa_user_id"`
	EmailVerified    bool   `db:"email_verified"`
}
