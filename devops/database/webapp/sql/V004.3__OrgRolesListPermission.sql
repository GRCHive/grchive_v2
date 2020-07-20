INSERT INTO permissions (human_name, description)
VALUES
    ('org.role.list', 'Ability to view all of the roles created for the organization.');

DO $$
BEGIN
    PERFORM update_admin_roles_with_all_permissions(r.id)
    FROM roles AS r
    WHERE r.is_admin_role = true;
END $$;
