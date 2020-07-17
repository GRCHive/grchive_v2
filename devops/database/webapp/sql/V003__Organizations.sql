CREATE TABLE organizations (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(256) NOT NULL,
    description TEXT DEFAULT '',
    parent_org_id BIGINT REFERENCES organizations(id) ON DELETE CASCADE,
    owner_user_id BIGINT NOT NULL REFERENCES users(id) ON DELETE RESTRICT,
    UNIQUE(name, parent_org_id)
);

CREATE INDEX ON organizations(parent_org_id);
CREATE INDEX ON organizations(owner_user_id);
CREATE INDEX ON organizations(name);
