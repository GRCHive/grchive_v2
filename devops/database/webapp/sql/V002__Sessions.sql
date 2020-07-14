CREATE TABLE user_sessions (
    id VARCHAR(36) PRIMARY KEY,
    session_expiration_time TIMESTAMPTZ NOT NULL,
    jwt_expiration_time TIMESTAMPTZ NOT NULL,
    access_token TEXT NOT NULL,
    refresh_token TEXT NOT NULL,
    user_id BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE
);

CREATE INDEX ON user_sessions(user_id);
