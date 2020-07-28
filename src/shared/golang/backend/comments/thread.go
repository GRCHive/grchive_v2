package comments

import (
	"fmt"
)

func (m *CommentManager) GetThreadIdForInventoryResource(resourceTable string, resourceId int64, engagementId int64, orgId int64) (int64, error) {
	id := int64(-1)
	err := m.db.Get(&id, fmt.Sprintf(`
		SELECT t.thread_id
		FROM %[1]s_comment_threads AS t
		INNER JOIN %[1]s AS r
			ON r.id = t.resource_id
		INNER JOIN inventory AS inv
			ON inv.id = r.inventory_id
		INNER JOIN engagements AS e
			ON e.id = inv.engagement_id
		WHERE
			r.id = $1 AND
			e.id = $2 AND
			e.org_id = $3
	`, resourceTable), resourceId, engagementId, orgId)
	return id, err
}

func (m *CommentManager) GetThreadIdForGenericResource(resourceTable string, resourceId int64, engagementId int64, orgId int64) (int64, error) {
	id := int64(-1)
	err := m.db.Get(&id, fmt.Sprintf(`
		SELECT t.thread_id
		FROM %[1]s_comment_threads AS t
		INNER JOIN %[1]s AS r
			ON r.id = t.resource_id
		INNER JOIN engagements AS e
			ON e.id = r.engagement_id
		WHERE
			r.id = $1 AND
			e.id = $2 AND
			e.org_id = $3
	`, resourceTable), resourceId, engagementId, orgId)
	return id, err
}

func (m *CommentManager) GetThreadIdForRisk(riskId int64, engagementId int64, orgId int64) (int64, error) {
	return m.GetThreadIdForGenericResource("risks", riskId, engagementId, orgId)
}

func (m *CommentManager) GetThreadIdForControl(controlId int64, engagementId int64, orgId int64) (int64, error) {
	return m.GetThreadIdForGenericResource("controls", controlId, engagementId, orgId)
}

func (m *CommentManager) GetThreadIdForGeneralLedger(ledgerId int64, engagementId int64, orgId int64) (int64, error) {
	return m.GetThreadIdForGenericResource("general_ledger", ledgerId, engagementId, orgId)
}

func (m *CommentManager) GetThreadIdForGLAccount(accId int64, engagementId int64, orgId int64) (int64, error) {
	return m.GetThreadIdForGenericResource("gl_accounts", accId, engagementId, orgId)
}

func (m *CommentManager) GetThreadIdForVendor(vendorId int64, engagementId int64, orgId int64) (int64, error) {
	return m.GetThreadIdForGenericResource("vendors", vendorId, engagementId, orgId)
}

func (m *CommentManager) GetThreadIdForServer(serverId int64, engagementId int64, orgId int64) (int64, error) {
	return m.GetThreadIdForInventoryResource("inventory_servers", serverId, engagementId, orgId)
}

func (m *CommentManager) GetThreadIdForDesktop(desktopId int64, engagementId int64, orgId int64) (int64, error) {
	return m.GetThreadIdForInventoryResource("inventory_desktops", desktopId, engagementId, orgId)
}

func (m *CommentManager) GetThreadIdForLaptop(laptopId int64, engagementId int64, orgId int64) (int64, error) {
	return m.GetThreadIdForInventoryResource("inventory_laptops", laptopId, engagementId, orgId)
}

func (m *CommentManager) GetThreadIdForMobile(mobileId int64, engagementId int64, orgId int64) (int64, error) {
	return m.GetThreadIdForInventoryResource("inventory_mobile", mobileId, engagementId, orgId)
}

func (m *CommentManager) GetThreadIdForEmbedded(embeddedId int64, engagementId int64, orgId int64) (int64, error) {
	return m.GetThreadIdForInventoryResource("inventory_embedded", embeddedId, engagementId, orgId)
}

func (m *CommentManager) GetThreadIdForDatabase(dbId int64, engagementId int64, orgId int64) (int64, error) {
	return m.GetThreadIdForGenericResource("databases", dbId, engagementId, orgId)
}

func (m *CommentManager) GetThreadIdForSystem(systemId int64, engagementId int64, orgId int64) (int64, error) {
	return m.GetThreadIdForGenericResource("systems", systemId, engagementId, orgId)
}
