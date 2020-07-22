package comments

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

func (m *CommentManager) CreateComment(tx *sqlx.Tx, comment *Comment) error {
	fmt.Printf("Create comment: %d %d %s\n", comment.ThreadId, comment.UserId, comment.Content)
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
