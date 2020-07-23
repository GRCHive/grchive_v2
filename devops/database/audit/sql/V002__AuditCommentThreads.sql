ALTER TABLE audit.audit_trail
ADD COLUMN thread_id BIGINT NOT NULL DEFAULT -1;

CREATE INDEX ON audit.audit_trail(thread_id);

ALTER TABLE audit.audit_trail
ALTER COLUMN thread_id DROP DEFAULT;
