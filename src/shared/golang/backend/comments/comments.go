package comments

import (
	"github.com/jmoiron/sqlx"
	"time"
)

type CommentManager struct {
	db *sqlx.DB
}

func CreateCommentManager(db *sqlx.DB) *CommentManager {
	return &CommentManager{
		db: db,
	}
}

type Comment struct {
	Id       int64     `db:"id"`
	ThreadId int64     `db:"thread_id"`
	UserId   int64     `db:"user_id"`
	PostTime time.Time `db:"post_time"`
	Content  string    `db:"content"`
}
