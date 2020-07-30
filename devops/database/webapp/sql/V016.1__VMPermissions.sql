INSERT INTO permissions (human_name, description)
VALUES
    ('org.vm.list', 'Ability to list all of the virtual machines for an engagement.'),
    ('org.vm.create', 'Ability to create new virtual machines for an engagement.'),
    ('org.vm.view', 'Ability to view a virtual machine.'),
    ('org.vm.delete', 'Ability to delete a virtual machine.'),
    ('org.vm.update', 'Ability to update a virtual machine.');

DO $$
BEGIN
    PERFORM update_admin_roles_with_all_permissions(r.id)
    FROM roles AS r
    WHERE r.is_admin_role = true;
END $$;
