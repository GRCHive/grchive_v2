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
                    ON gl.parent_account_id = lq.id
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
