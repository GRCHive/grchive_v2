package backend

import (
	"gitlab.com/grchive/grchive-v2/shared/backend/inventory"
)

type EngagementScopingStats struct {
	NumRisks                         int
	NumControls                      int
	NumGLAccounts                    int
	NumFinanciallyRelevantGLAccounts int
	NumVendors                       int
	NumVendorProducts                int
	NumSystems                       int
	NumDatabases                     int
	NumInventory                     map[inventory.InventoryType]int
}

func (m *BackendInterface) GetEngagementScopingStats(engagementId int64) (*EngagementScopingStats, error) {
	var err error
	stats := EngagementScopingStats{}

	stats.NumRisks, err = m.Risks.GetNumRisksForEngagement(engagementId)
	if err != nil {
		return nil, err
	}

	stats.NumControls, err = m.Controls.GetNumControlsForEngagement(engagementId)
	if err != nil {
		return nil, err
	}

	glAccountsNotRelevant, err := m.GL.GetNumGLAccountsForEngagement(engagementId, false)
	if err != nil {
		return nil, err
	}

	glAccountsRelevant, err := m.GL.GetNumGLAccountsForEngagement(engagementId, true)
	if err != nil {
		return nil, err
	}

	stats.NumGLAccounts = glAccountsNotRelevant + glAccountsRelevant
	stats.NumFinanciallyRelevantGLAccounts = glAccountsRelevant

	stats.NumVendors, err = m.Vendors.GetNumVendorsForEngagement(engagementId)
	if err != nil {
		return nil, err
	}

	stats.NumVendorProducts, err = m.Vendors.GetNumVendorProductsForEngagement(engagementId)
	if err != nil {
		return nil, err
	}

	stats.NumSystems, err = m.Systems.GetNumSystemsForEngagement(engagementId)
	if err != nil {
		return nil, err
	}

	stats.NumDatabases, err = m.Databases.GetNumDatabasesForEngagement(engagementId)
	if err != nil {
		return nil, err
	}

	stats.NumInventory, err = m.Inventory.GetNumInventoryForEngagement(engagementId)
	if err != nil {
		return nil, err
	}
	return &stats, nil
}
