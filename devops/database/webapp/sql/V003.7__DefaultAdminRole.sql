ALTER TABLE roles 
ADD COLUMN is_admin_role BOOLEAN NOT NULL DEFAULT false;

CREATE UNIQUE INDEX ON roles(org_id, is_admin_role);

CREATE OR REPLACE FUNCTION create_default_admin_role_for_org()
    RETURNS trigger AS
$$
    DECLARE
        admin_role_id BIGINT;
    BEGIN
        INSERT INTO roles (org_id, name, description, is_admin_role)
        VALUES (NEW.id, 'Admin', 'Default admin role.', true)
        RETURNING id INTO admin_role_id;

        INSERT INTO role_permissions (role_id, permission_id)
        SELECT admin_role_id, permissions.id
        FROM permissions;
        
        INSERT INTO user_roles (user_id, role_id, org_id)
        VALUES (NEW.owner_user_id, admin_role_id, NEW.id);

        RETURN NEW;
    END;
$$ LANGUAGE plpgsql;

DROP TRIGGER IF EXISTS trigger_create_default_admin_role_for_org ON organizations;
CREATE TRIGGER trigger_create_default_admin_role_for_org
    AFTER INSERT ON organizations
    FOR EACH ROW
    EXECUTE FUNCTION create_default_admin_role_for_org();
