package gl

func (m *GLManager) ListGLAccountsForEngagement(engagementId int64) ([]*GLAccount, error) {
	accs := []*GLAccount{}
	err := m.db.Select(&accs, `
		SELECT *
		FROM gl_accounts
		WHERE engagement_id = $1
	`, engagementId)
	return accs, err
}
