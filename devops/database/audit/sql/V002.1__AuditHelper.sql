CREATE OR REPLACE FUNCTION audit.generic_log_audit_trail(
    table_name VARCHAR(64),
    pk1_col TEXT,
    table_action VARCHAR(64),
    table_old_val JSONB,
    table_new_val JSONB
)
    RETURNS void AS
$$
    DECLARE
        jdiff JSONB; 
        table_pk1 VARCHAR;
    BEGIN
        SELECT audit.jsonb_diff(table_old_val, table_new_val)
        INTO jdiff;

        SELECT COALESCE(
            jsonb_strip_nulls(table_old_val)->pk1_col::VARCHAR,
            jsonb_strip_nulls(table_new_val)->pk1_col::VARCHAR
        )
        INTO table_pk1;

        INSERT INTO audit.audit_trail (
            org_id,
            engagement_id,
            thread_id,
            table_name,
            table_pk1,
            action,
            performed_at,
            ip_address, 
            user_identifier,
            table_column,
            table_old_value,
            table_new_value
        )
        SELECT
            current_setting('grchive.current_org_id')::BIGINT,
            current_setting('grchive.current_engagement_id')::BIGINT,
            current_setting('grchive.current_thread_id')::BIGINT,
            table_name,
            table_pk1,
            table_action,
            NOW(),
            current_setting('grchive.ip_address')::INET,
            current_setting('grchive.current_user_id')::BIGINT,
            jd.key,
            jd.value -> 'old',
            jd.value -> 'new'
        FROM jsonb_each(jdiff) AS jd;
    END;
$$ LANGUAGE plpgsql;


