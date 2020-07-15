CREATE TABLE user_orgs (
    user_id BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    org_id BIGINT NOT NULL REFERENCES organizations(id) ON DELETE CASCADE,
    UNIQUE(user_id, org_id)
);

CREATE INDEX ON user_orgs(org_id);

CREATE OR REPLACE FUNCTION sync_user_orgs_with_org_owner()
    RETURNS trigger AS
$$
BEGIN
    INSERT INTO user_orgs (user_id, org_id)
    VALUES (NEW.owner_user_id, NEW.id)
    ON CONFLICT (user_id, org_id) DO NOTHING;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_sync_user_orgs_with_org_owner
    AFTER INSERT OR UPDATE ON organizations
    FOR EACH ROW
    EXECUTE PROCEDURE sync_user_orgs_with_org_owner();
