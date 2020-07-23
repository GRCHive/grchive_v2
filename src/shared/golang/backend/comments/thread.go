package comments

func (m *CommentManager) GetThreadIdForRisk(riskId int64, engagementId int64, orgId int64) (int64, error) {
	id := int64(-1)
	err := m.db.Get(&id, `
		SELECT t.thread_id
		FROM risks_comment_threads AS t
		INNER JOIN risks AS r
			ON r.id = t.resource_id
		INNER JOIN engagements AS e
			ON e.id = r.engagement_id
		WHERE
			r.id = $1 AND
			e.id = $2 AND
			e.org_id = $3
	`, riskId, engagementId, orgId)
	return id, err
}

func (m *CommentManager) GetThreadIdForControl(controlId int64, engagementId int64, orgId int64) (int64, error) {
	id := int64(-1)
	err := m.db.Get(&id, `
		SELECT t.thread_id
		FROM controls_comment_threads AS t
		INNER JOIN controls AS r
			ON r.id = t.resource_id
		INNER JOIN engagements AS e
			ON e.id = r.engagement_id
		WHERE
			r.id = $1 AND
			e.id = $2 AND
			e.org_id = $3
	`, controlId, engagementId, orgId)
	return id, err
}

func (m *CommentManager) GetThreadIdForGeneralLedger(ledgerId int64, engagementId int64, orgId int64) (int64, error) {
	id := int64(-1)
	err := m.db.Get(&id, `
		SELECT t.thread_id
		FROM general_ledger_comment_threads AS t
		INNER JOIN general_ledger AS r
			ON r.id = t.resource_id
		INNER JOIN engagements AS e
			ON e.id = r.engagement_id
		WHERE
			r.id = $1 AND
			e.id = $2 AND
			e.org_id = $3
	`, ledgerId, engagementId, orgId)
	return id, err
}
