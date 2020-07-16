CREATE TABLE audit.audit_trail (
    id BIGSERIAL PRIMARY KEY,
    org_id BIGINT,
    table_name VARCHAR(64) NOT NULL,
    table_pk1 VARCHAR NOT NULL,
    action VARCHAR(64) NOT NULL,
    performed_at TIMESTAMPTZ NOT NULL,
    ip_address INET NOT NULL,
    user_identifier BIGINT NOT NULL,
    table_column VARCHAR(64) NOT NULL,
    -- { value: ?? }
    table_old_value JSONB,
    table_new_value JSONB
);

CREATE INDEX ON audit.audit_trail(org_id);
CREATE INDEX ON audit.audit_trail(table_name, table_pk1);
CREATE INDEX ON audit.audit_trail(user_identifier);
CREATE INDEX ON audit.audit_trail(ip_address);
CREATE INDEX ON audit.audit_trail(action);
