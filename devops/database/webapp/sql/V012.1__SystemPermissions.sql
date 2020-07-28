INSERT INTO permissions (human_name, description)
VALUES
    ('org.systems.list', 'Ability to list all of the systems for an engagement.'),
    ('org.systems.create', 'Ability to create new systems for an engagement.'),
    ('org.systems.view', 'Ability to view a system.'),
    ('org.systems.delete', 'Ability to delete a system.'),
    ('org.systems.update', 'Ability to update a system.');

DO $$
BEGIN
    PERFORM update_admin_roles_with_all_permissions(r.id)
    FROM roles AS r
    WHERE r.is_admin_role = true;
END $$;
