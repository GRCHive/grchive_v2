package vm

func (m *VirtualMachineManager) ListVMForEngagement(engagementId int64) ([]*VirtualMachine, error) {
	vms := []*VirtualMachine{}
	err := m.db.Select(&vms, `
		SELECT
			vm.id AS "id",
			vm.engagement_id AS "engagement_id",
			vm.unique_id AS "unique_id",
			vm.name AS "name",
			vm.description AS "description",
			vm.vcpus AS "vcpus",
			vm.memory_bytes AS "memory_bytes",
			vm.base_state_id AS "base_state_id",
			ms.id AS "state.id",
			ms.engagement_id AS "state.engagement_id",
			ms.unique_id AS "state.unique_id",
			ms.hypervisor AS "state.hypervisor",
			ms.operating_system AS "state.operating_system"
		FROM virtual_machines AS vm
		INNER JOIN machine_state AS ms
			ON ms.id = vm.state_id
		WHERE vm.engagement_id = $1
	`, engagementId)
	return vms, err
}
