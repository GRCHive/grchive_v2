INSERT INTO permissions (human_name, description)
VALUES
    ('org.users.list', 'Ability to list all users in an org.'),
    ('org.users.view', 'Ability to view a given user in the org.');

DO $$
BEGIN
    PERFORM update_admin_roles_with_all_permissions(r.id)
    FROM roles AS r
    WHERE r.is_admin_role = true;
END $$;
