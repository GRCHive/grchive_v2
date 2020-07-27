CREATE TABLE inventory (
    id BIGSERIAL PRIMARY KEY,
    engagement_id BIGINT NOT NULL REFERENCES engagements(id) ON DELETE CASCADE,
    unique_id VARCHAR NOT NULL,
    name VARCHAR NOT NULL,
    description TEXT NOT NULL DEFAULT '',
    brand TEXT DEFAULT '',
    model TEXT DEFAULT '',
    UNIQUE(engagement_id, unique_id)
);

CREATE TABLE inventory_servers (
    id BIGSERIAL PRIMARY KEY,
    inventory_id BIGINT NOT NULL REFERENCES inventory(id) ON DELETE CASCADE,
    physical_location TEXT NOT NULL DEFAULT '',
    operating_system TEXT,
    hypervisor TEXT,
    static_external_ip INET
);

CREATE TABLE inventory_desktops (
    id BIGSERIAL PRIMARY KEY,
    inventory_id BIGINT NOT NULL REFERENCES inventory(id) ON DELETE CASCADE,
    physical_location TEXT NOT NULL DEFAULT '',
    operating_system TEXT NOT NULL DEFAULT ''
);

CREATE TABLE inventory_laptops (
    id BIGSERIAL PRIMARY KEY,
    inventory_id BIGINT NOT NULL REFERENCES inventory(id) ON DELETE CASCADE,
    operating_system TEXT NOT NULL DEFAULT ''
);

CREATE TABLE inventory_mobile (
    id BIGSERIAL PRIMARY KEY,
    inventory_id BIGINT NOT NULL REFERENCES inventory(id) ON DELETE CASCADE,
    operating_system TEXT NOT NULL DEFAULT ''
);

CREATE TABLE inventory_embedded (
    id BIGSERIAL PRIMARY KEY,
    inventory_id BIGINT NOT NULL REFERENCES inventory(id) ON DELETE CASCADE,
    operating_system TEXT DEFAULT ''
);

CREATE TABLE inventory_owners (
    id BIGSERIAL PRIMARY KEY,
    inventory_id BIGINT NOT NULL REFERENCES inventory(id) ON DELETE CASCADE,
    user_id BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE
);

CREATE INDEX ON inventory_owners(inventory_id);
