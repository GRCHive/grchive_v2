CREATE TABLE systems (
    id BIGSERIAL PRIMARY KEY,
    engagement_id BIGINT NOT NULL REFERENCES engagements(id) ON DELETE CASCADE,
    unique_id VARCHAR NOT NULL,
    name VARCHAR NOT NULL,
    description TEXT NOT NULL DEFAULT '',
    purpose TEXT NOT NULL DEFAULT '',
    UNIQUE(engagement_id, unique_id)
);
