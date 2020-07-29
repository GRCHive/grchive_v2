INSERT INTO permissions (human_name, description)
VALUES
    ('org.rel.risks.controls.list', 'Ability to list all relationships between risks and controls for an engagement.'),
    ('org.rel.risks.controls.create', 'Ability to create new relationships between risks and controls for an engagement.'),
    ('org.rel.risks.controls.delete', 'Ability to delete a relationship between a risk and control.');

DO $$
BEGIN
    PERFORM update_admin_roles_with_all_permissions(r.id)
    FROM roles AS r
    WHERE r.is_admin_role = true;
END $$;
