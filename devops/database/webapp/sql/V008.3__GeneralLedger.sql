-- This table is mainly so we can support the automatic creation of a comment thread for the entire GL.
CREATE TABLE general_ledger (
    id BIGSERIAL PRIMARY KEY,
    engagement_id BIGINT NOT NULL UNIQUE REFERENCES engagements(id) ON DELETE CASCADE
);

ALTER TABLE gl_accounts
ADD COLUMN gl_id BIGINT REFERENCES general_ledger(id) ON DELETE CASCADE;

CREATE INDEX ON gl_accounts(gl_id);

DO $$
BEGIN
    PERFORM create_comment_thread_for_table('general_ledger', 'id');
END $$;

INSERT INTO general_ledger (engagement_id)
SELECT e.id
FROM engagements AS e;

SET "grchive.current_org_id" TO '-1';
SET "grchive.current_engagement_id" TO '-1';
SET "grchive.current_thread_id" TO '-1';
SET "grchive.ip_address" TO '127.0.0.1';
SET "grchive.current_user_id" TO '-1';

UPDATE gl_accounts
SET gl_id = subquery.id
FROM (
    SELECT *
    FROM general_ledger
) AS subquery
WHERE gl_accounts.engagement_id = subquery.engagement_id;

CREATE OR REPLACE FUNCTION create_general_ledger_for_engagement()
    RETURNS trigger AS
$$
    BEGIN
        INSERT INTO general_ledger (engagement_id)
        VALUES (NEW.id);

        RETURN NEW;
    END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_create_general_ledger_for_engagement
    AFTER INSERT ON engagements
    FOR EACH ROW
    EXECUTE FUNCTION create_general_ledger_for_engagement();
