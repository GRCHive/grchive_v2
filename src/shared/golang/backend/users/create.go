package users

import (
	"github.com/jmoiron/sqlx"
)

func (m *UserManager) CreateUser(tx *sqlx.Tx, user *User) error {
	rows, err := tx.NamedQuery(`
		INSERT INTO users (full_name, email, fa_user_id)
		VALUES (:full_name, :email, :fa_user_id)
		RETURNING id
	`, user)

	if err != nil {
		return err
	}
	defer rows.Close()
	rows.Next()
	return rows.Scan(&user.Id)
}
