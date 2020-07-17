INSERT INTO permissions (human_name, description)
VALUES
    ('org.engagements.list', 'Ability to view all of the engagements for the org. If not, only allow listing of assigned engagements.'),
    ('org.engagements.create', 'Ability to create new engagements for the org. The role must also be assigned to the engagement for this permission to take effect.'),
    ('org.engagements.view', 'Ability to view an engagement. The role must also be assigned to the engagement for this permission to take effect.'),
    ('org.engagements.delete', 'Ability to delete an engagement. The role must also be assigned to the engagement for this permission to take effect.'),
    ('org.engagements.update', 'Ability to update information about an engagement. The role must also be assigned to the engagement for this permission to take effect.'),
    ('org.engagements.close', 'Ability to close an engagement. The role must also be assigned to the engagement for this permission to take effect.'),
    ('org.engagements.reopen', 'Ability to reopen an engagement once closed. The role must also be assigned to the engagement for this permission to take effect.');

CREATE OR REPLACE FUNCTION update_admin_roles_with_all_permissions(admin_role_id BIGINT)
    RETURNS void AS
$$
    BEGIN
        INSERT INTO role_permissions (role_id, permission_id)
        SELECT admin_role_id, permissions.id
        FROM permissions
        ON CONFLICT (role_id, permission_id) DO NOTHING;
    END;
$$ LANGUAGE plpgsql;

DO $$
BEGIN
    PERFORM update_admin_roles_with_all_permissions(r.id)
    FROM roles AS r
    WHERE r.is_admin_role = true;
END $$;
