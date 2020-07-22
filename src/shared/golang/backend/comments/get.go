package comments

func (m *CommentManager) GetCommentById(commentId int64, threadId int64) (*Comment, error) {
	comment := Comment{}
	err := m.db.Get(&comment, `
		SELECT *
		FROM comments
		WHERE id = $1 AND thread_id = $2
	`, commentId, threadId)
	return &comment, err
}
