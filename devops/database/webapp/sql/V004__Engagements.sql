CREATE TABLE engagements (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR NOT NULL,
    description TEXT DEFAULT '',
    org_id BIGINT REFERENCES organizations(id) ON DELETE CASCADE,
    created_time TIMESTAMPTZ NOT NULL,
    start_time TIMESTAMPTZ,
    end_time TIMESTAMPTZ,
    UNIQUE(name, org_id)
);

CREATE TABLE engagement_roles (
    id BIGSERIAL PRIMARY KEY,
    engagement_id BIGINT NOT NULL REFERENCES engagements(id) ON DELETE CASCADE,
    role_id BIGINT NOT NULL REFERENCES roles(id) ON DELETE CASCADE,
    org_id BIGINT NOT NULL REFERENCES organizations(id) ON DELETE CASCADE,
    UNIQUE(org_id, engagement_id, role_id)
);

CREATE OR REPLACE FUNCTION verify_engagement_role_valid_for_org()
    RETURNS trigger AS
$$
    DECLARE
        verify_role_id BIGINT;
        verify_engagement_id BIGINT;
    BEGIN
        -- Need to check that the (engagement_id, org_id) and (role_id, org_id) combo are both valid.
        SELECT id
        FROM roles
        WHERE org_id = NEW.org_id AND id = NEW.role_id
        INTO verify_role_id;

        SELECT id
        FROM engagements
        WHERE org_id = NEW.org_id AND id = NEW.engagement_id
        INTO verify_engagement_id;

        IF (verify_role_id IS NULL) OR (verify_engagement_id IS NULL) THEN
            RAISE EXCEPTION 'Engagement can not be assigned to a role created for an organization it is not a part of.';
        END IF;

        RETURN NEW;
    END;
$$ LANGUAGE plpgsql;

DROP TRIGGER IF EXISTS trigger_verify_engagement_role_valid ON engagement_roles;
CREATE CONSTRAINT TRIGGER trigger_verify_engagement_role_valid
    AFTER INSERT OR UPDATE ON engagement_roles INITIALLY DEFERRED
    FOR EACH ROW
    EXECUTE FUNCTION verify_engagement_role_valid_for_org();

DO $$
BEGIN
END $$;
