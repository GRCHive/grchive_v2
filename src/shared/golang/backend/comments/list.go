package comments

func (m *CommentManager) GetCommentsForThread(threadId int64) ([]*Comment, error) {
	comments := []*Comment{}
	err := m.db.Select(&comments, `
		SELECT *
		FROM comments
		WHERE thread_id = $1
		ORDER BY id ASC
	`, threadId)
	return comments, err
}
