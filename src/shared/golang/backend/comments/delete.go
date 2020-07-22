package comments

import (
	"github.com/jmoiron/sqlx"
)

func (m *CommentManager) DeleteComment(tx *sqlx.Tx, commentId int64) error {
	_, err := tx.Exec(`
		DELETE FROM comments
		WHERE id = $1
	`, commentId)
	return err
}
