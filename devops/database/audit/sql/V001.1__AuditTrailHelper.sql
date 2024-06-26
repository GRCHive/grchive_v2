CREATE OR REPLACE FUNCTION audit.jsonb_diff(a jsonb, b jsonb)
    RETURNS jsonb AS
$$
    DECLARE
        ret jsonb := '{}'::jsonb;
        merged_keys text[];
        key text;
        a_val text;
        b_val text;
    BEGIN
        WITH union_query AS (
            SELECT jsonb_object_keys(a)
            UNION SELECT jsonb_object_keys(b)
        )
        SELECT ARRAY_AGG(jsonb_object_keys)
        FROM union_query
        INTO merged_keys;

        FOREACH key IN ARRAY merged_keys
        LOOP
            SELECT a ->> key INTO a_val;
            SELECT b ->> key INTO b_val;

            IF (a_val != b_val) OR (a_val IS NULL AND b_val IS NOT NULL) OR (a_val IS NOT NULL AND b_val IS NULL) THEN
                ret = ret || jsonb_build_object(key, 
                    jsonb_build_object(
                        'new', b ->> key,
                        'old', a ->> key
                    )
                );
            END IF;
        END LOOP;

        RETURN ret;
    END;
$$ LANGUAGE plpgsql;

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

CREATE OR REPLACE FUNCTION audit.create_audit_trail_triggers_for_table(
    table_name VARCHAR(64),
    pk1_col TEXT
)
    RETURNS void AS
$BODY$
    BEGIN
        EXECUTE format($DEL$
            CREATE OR REPLACE FUNCTION public.audit_delete_%1$s()
                RETURNS trigger AS
            $$
                BEGIN
                    PERFORM audit.generic_log_audit_trail(
                        '%1$s',
                        '%2$s',
                        'DELETE',
                        to_jsonb(OLD),
                        '{}'::jsonb
                    );
                    RETURN OLD;
                END;
            $$ LANGUAGE plpgsql;
            DROP TRIGGER IF EXISTS trigger_audit_trail_delete ON %1$s;
            CREATE CONSTRAINT TRIGGER trigger_audit_trail_delete
                AFTER DELETE ON %1$s INITIALLY DEFERRED
                FOR EACH ROW
                EXECUTE FUNCTION public.audit_delete_%1$s();
        $DEL$
            , table_name
            , pk1_col
        );

        EXECUTE format($UP$
            CREATE OR REPLACE FUNCTION public.audit_update_%1$s()
                RETURNS trigger AS
            $$
                BEGIN
                    PERFORM audit.generic_log_audit_trail(
                        '%1$s',
                        '%2$s',
                        'UPDATE',
                        to_jsonb(OLD),
                        to_jsonb(NEW)
                    );
                    RETURN NEW;
                END;
            $$ LANGUAGE plpgsql;
            DROP TRIGGER IF EXISTS trigger_audit_trail_update ON %1$s;
            CREATE CONSTRAINT TRIGGER trigger_audit_trail_update
                AFTER UPDATE ON %1$s INITIALLY DEFERRED
                FOR EACH ROW
                EXECUTE FUNCTION public.audit_update_%1$s();
        $UP$
            , table_name
            , pk1_col
        );

        EXECUTE format($INS$
            CREATE OR REPLACE FUNCTION public.audit_insert_%1$s()
                RETURNS trigger AS
            $$
                BEGIN
                    PERFORM audit.generic_log_audit_trail(
                        '%1$s',
                        '%2$s',
                        'INSERT',
                        '{}'::jsonb,
                        to_jsonb(NEW)
                    );
                    RETURN NEW;
                END;
            $$ LANGUAGE plpgsql;
            DROP TRIGGER IF EXISTS trigger_audit_trail_insert ON %1$s;
            CREATE CONSTRAINT TRIGGER trigger_audit_trail_insert
                AFTER INSERT ON %1$s INITIALLY DEFERRED
                FOR EACH ROW
                EXECUTE FUNCTION public.audit_insert_%1$s();
        $INS$
            , table_name
            , pk1_col
        );
    END;
$BODY$ LANGUAGE plpgsql;
