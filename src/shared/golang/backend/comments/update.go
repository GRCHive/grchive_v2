package comments

import (
	"github.com/jmoiron/sqlx"
)

func (m *CommentManager) UpdateComment(tx *sqlx.Tx, comment *Comment) error {
	_, err := tx.NamedExec(`
		UPDATE comments
		SET content = :content
		WHERE id = :id
	`, comment)
	return err
}
