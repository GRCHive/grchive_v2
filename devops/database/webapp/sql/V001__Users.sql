CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    full_name VARCHAR(512) NOT NULL,
    email VARCHAR(320) NOT NULL UNIQUE,
    fa_user_id VARCHAR(36) NOT NULL UNIQUE
);

CREATE INDEX ON users(full_name);

DO $$
BEGIN
    PERFORM audit.create_audit_trail_triggers_for_table('users', NULL, 'id');
END $$;
