package comments

func (m *CommentManager) GetCommentsForThread(threadId int64) ([]*Comment, error) {
	comments := []*Comment{}
	err := m.db.Select(&comments, `
		SELECT *
		FROM comments
		WHERE thread_id = $1
		ORDER BY id DESC
	`, threadId)
	return comments, err
}
