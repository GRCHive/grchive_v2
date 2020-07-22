package comments

import (
	"github.com/jmoiron/sqlx"
)

func (m *CommentManager) CreateComment(tx *sqlx.Tx, comment *Comment) error {
	rows, err := tx.NamedQuery(`
		INSERT INTO comments (thread_id, user_id, post_time, content)
		VALUES (:thread_id, :user_id, NOW(), :content)
		RETURNING id
	`, comment)

	if err != nil {
		return err
	}
	defer rows.Close()

	rows.Next()
	return rows.Scan(&comment.Id)
}
