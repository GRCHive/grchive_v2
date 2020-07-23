package gl

func (m *GLManager) GetGeneralLedgerForEngagement(engagementId int64) (*GeneralLedger, error) {
	ledger := GeneralLedger{}
	err := m.db.Get(&ledger, `
		SELECT *
		FROM general_ledger
		WHERE engagement_id = $1
	`, engagementId)
	return &ledger, err
}

func (m *GLManager) GetGLAccountFromId(id int64) (*GLAccount, error) {
	acc := GLAccount{}
	err := m.db.Get(&acc, `
		SELECT *
		FROM gl_accounts
		WHERE id = $1
	`, id)
	return &acc, err
}
