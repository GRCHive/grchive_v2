INSERT INTO permissions (human_name, description)
VALUES
    ('org.controls.list', 'Ability to list all of the controls for an engagement.'),
    ('org.controls.create', 'Ability to create new controls for an engagement.'),
    ('org.controls.view', 'Ability to view a control.'),
    ('org.controls.delete', 'Ability to delete a control.'),
    ('org.controls.update', 'Ability to update a control.');

DO $$
BEGIN
    PERFORM update_admin_roles_with_all_permissions(r.id)
    FROM roles AS r
    WHERE r.is_admin_role = true;
END $$;
