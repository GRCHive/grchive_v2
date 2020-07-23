INSERT INTO permissions (human_name, description)
VALUES
    ('org.gl.list', 'Ability to list all of the general ledger accounts for an engagement.'),
    ('org.gl.create', 'Ability to create new general ledger account for an engagement.'),
    ('org.gl.view', 'Ability to view a general ledger account.'),
    ('org.gl.delete', 'Ability to delete a general ledger account.'),
    ('org.gl.update', 'Ability to update a general ledger account.');

DO $$
BEGIN
    PERFORM update_admin_roles_with_all_permissions(r.id)
    FROM roles AS r
    WHERE r.is_admin_role = true;
END $$;
