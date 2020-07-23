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
