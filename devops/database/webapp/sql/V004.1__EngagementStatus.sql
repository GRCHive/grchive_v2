CREATE TABLE engagement_status (
    id BIGSERIAL PRIMARY KEY,
    engagement_id BIGINT NOT NULL REFERENCES engagements(id) ON DELETE CASCADE,
    is_closed BOOLEAN NOT NULL,
    change_time TIMESTAMPTZ NOT NULL,
    change_user_id BIGINT REFERENCES users(id) ON DELETE SET NULL
);

CREATE INDEX ON engagement_status(engagement_id);
