INSERT INTO permissions (human_name, description)
VALUES
    ('org.risks.list', 'Ability to list all of the risks for an engagement.'),
    ('org.risks.create', 'Ability to create new risks for an engagement.'),
    ('org.risks.view', 'Ability to view a risk.'),
    ('org.risks.delete', 'Ability to delete a risk.'),
    ('org.risks.update', 'Ability to update a risk.');

DO $$
BEGIN
    PERFORM update_admin_roles_with_all_permissions(r.id)
    FROM roles AS r
    WHERE r.is_admin_role = true;
END $$;
