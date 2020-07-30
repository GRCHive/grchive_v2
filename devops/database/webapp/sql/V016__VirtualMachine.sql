CREATE TABLE virtual_machines (
    id BIGSERIAL PRIMARY KEY,
    engagement_id BIGINT NOT NULL REFERENCES engagements(id) ON DELETE CASCADE,
    unique_id VARCHAR NOT NULL,
    name VARCHAR NOT NULL,
    description TEXT NOT NULL DEFAULT '',
    vcpus INTEGER NOT NULL DEFAULT 1,
    memory_bytes BIGINT NOT NULL DEFAULT 0,
    state_id BIGINT NOT NULL REFERENCES machine_state(id) ON DELETE CASCADE,
    base_state_id BIGINT REFERENCES machine_state(id) ON DELETE SET NULL,
    UNIQUE(engagement_id, unique_id)
);
