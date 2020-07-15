package users

import (
	"github.com/jmoiron/sqlx"
)

func (m *UserManager) MarkUserVerified(tx *sqlx.Tx, userId int64) error {
	_, err := tx.Exec(`
		UPDATE users
		SET email_verified = true
		WHERE id = $1
	`, userId)
	return err
}

func (m *UserManager) UpdateUser(tx *sqlx.Tx, user *User) error {
	_, err := tx.NamedExec(`
		UPDATE users
		SET full_name = :full_name
		WHERE id = :id
	`, user)
	return err
}
