INSERT INTO permissions (human_name, description)
VALUES
    ('org.databases.list', 'Ability to list all of the databases for an engagement.'),
    ('org.databases.create', 'Ability to create new databases for an engagement.'),
    ('org.databases.view', 'Ability to view a database.'),
    ('org.databases.delete', 'Ability to delete a database.'),
    ('org.databases.update', 'Ability to update a database.');

DO $$
BEGIN
    PERFORM update_admin_roles_with_all_permissions(r.id)
    FROM roles AS r
    WHERE r.is_admin_role = true;
END $$;

