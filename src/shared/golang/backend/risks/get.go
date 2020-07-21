package risks

func (m *RiskManager) GetRiskFromId(riskId int64) (*Risk, error) {
	ret := Risk{}
	err := m.db.Get(&ret, `
		SELECT *
		FROM risks
		WHERE id = $1
	`, riskId)
	return &ret, err
}
