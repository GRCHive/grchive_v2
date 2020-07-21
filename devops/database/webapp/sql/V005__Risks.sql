CREATE TABLE risks (
    id BIGSERIAL PRIMARY KEY,
    engagement_id BIGINT NOT NULL REFERENCES engagements(id),
    name VARCHAR NOT NULL,
    human_id VARCHAR(256) NOT NULL,
    severity SMALLINT NOT NULL DEFAULT 0,
    severity_justification TEXT,
    UNIQUE(engagement_id, human_id)
);
