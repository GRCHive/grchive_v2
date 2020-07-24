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

func (m *GLManager) GetGLSubaccountsFromId(id int64) ([]*GLAccount, error) {
	accs := []*GLAccount{}
	err := m.db.Select(&accs, `
		WITH RECURSIVE subaccs AS (
			SELECT *
			FROM gl_accounts
			WHERE id = $1
			UNION
				SELECT gl.*
				FROM gl_accounts AS gl
				INNER JOIN subaccs AS sub
					ON sub.id = gl.parent_account_id
		)
		SELECT *
		FROM subaccs
	`, id)
	return accs, err
}

func (m *GLManager) GetGLParentAccountsFromId(id int64) ([]*GLAccount, error) {
	accs := []*GLAccount{}
	err := m.db.Select(&accs, `
		WITH RECURSIVE subaccs AS (
			SELECT *
			FROM gl_accounts
			WHERE id = $1
			UNION
				SELECT gl.*
				FROM gl_accounts AS gl
				INNER JOIN subaccs AS sub
					ON sub.parent_account_id = gl.id
		)
		SELECT *
		FROM subaccs
	`, id)
	return accs, err
}
