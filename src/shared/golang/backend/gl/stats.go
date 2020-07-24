package gl

func (m *GLManager) GetNumGLAccountsForEngagement(engagementId int64, relevant bool) (int, error) {
	ret := 0
	err := m.db.Get(&ret, `
		SELECT COUNT(id)
		FROM gl_accounts
		WHERE engagement_id = $1
			AND financially_relevant = $2
	`, engagementId, relevant)
	return ret, err
}
