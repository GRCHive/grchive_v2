CREATE OR REPLACE FUNCTION verify_gl_account_no_loop()
    RETURNS trigger AS
$$
    BEGIN
        IF EXISTS (
            WITH RECURSIVE loop_query(id, parent_account_id) AS (
                SELECT gl.id, gl.parent_account_id
                FROM gl_accounts AS gl
                WHERE gl.parent_account_id = NEW.id

                UNION ALL
                SELECT gl.id, gl.parent_account_id
                FROM gl_accounts AS gl
                INNER JOIN loop_query AS lq
                    ON lq.parent_account_id = gl.id
            )
            SELECT 1
            FROM loop_query
            WHERE id = NEW.id
            LIMIT 1
        )
        THEN
            RAISE EXCEPTION 'GL Account loop detected.';
        ELSE
            RETURN NEW;
        END IF;
    END;
$$ LANGUAGE plpgsql;

DROP TRIGGER IF EXISTS trigger_verify_gl_account_no_loop ON gl_accounts;
CREATE CONSTRAINT TRIGGER trigger_verify_gl_account_no_loop
    AFTER INSERT OR UPDATE ON gl_accounts INITIALLY DEFERRED
    FOR EACH ROW
    EXECUTE FUNCTION verify_gl_account_no_loop();

