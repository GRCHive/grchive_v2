CREATE TABLE user_roles (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    role_id BIGINT NOT NULL REFERENCES roles(id) ON DELETE CASCADE,
    org_id BIGINT NOT NULL REFERENCES organizations(id) ON DELETE CASCADE
);

CREATE OR REPLACE FUNCTION verify_user_role_valid_for_org()
    RETURNS trigger AS
$$
    DECLARE
        verify_role_id BIGINT;
        verify_user_id BIGINT;
    BEGIN
        -- Need to check that the (user_id, org_id) and (role_id, org_id) combo are both valid.
        SELECT id
        FROM roles
        WHERE org_id = NEW.org_id AND id = NEW.role_id
        INTO verify_role_id;

        SELECT user_id
        FROM user_orgs
        WHERE org_id = NEW.org_id AND user_id = NEW.user_id
        INTO verify_user_id;

        IF (verify_role_id IS NULL) OR (verify_user_id IS NULL) THEN
            RAISE EXCEPTION 'User can not be assigned to a role created for an organization they are not a part of.';
        END IF;

        RETURN NEW;
    END;
$$ LANGUAGE plpgsql;

DROP TRIGGER IF EXISTS trigger_verify_user_role_valid ON user_roles;
CREATE CONSTRAINT TRIGGER trigger_verify_user_role_valid
    AFTER INSERT OR UPDATE ON user_roles INITIALLY DEFERRED
    FOR EACH ROW
    EXECUTE FUNCTION verify_user_role_valid_for_org();
