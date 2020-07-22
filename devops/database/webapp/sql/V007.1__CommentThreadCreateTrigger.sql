CREATE OR REPLACE FUNCTION create_comment_thread_for_table(
    table_name VARCHAR(64),
    pk1_col TEXT
)
    RETURNS void AS
$BODY$
    DECLARE
        comment_thread_table_name VARCHAR(64);
    BEGIN
        comment_thread_table_name := table_name || '_comment_threads';

        EXECUTE format($COM$
            DROP TABLE IF EXISTS %3$s;
            CREATE TABLE %3$s (
                thread_id BIGINT NOT NULL UNIQUE REFERENCES comment_threads(id) ON DELETE CASCADE,
                resource_id BIGINT NOT NULL UNIQUE REFERENCES %1$s(%2$s) ON DELETE CASCADE 
            );
            
            CREATE OR REPLACE FUNCTION create_comment_thread_%1$s()
                RETURNS trigger AS
            $$
                DECLARE
                    new_thread_id BIGINT;
                BEGIN
                    INSERT INTO comment_threads VALUES (default) RETURNING id INTO new_thread_id;

                    INSERT INTO %3$s (thread_id, resource_id)
                    VALUES (new_thread_id, NEW.%2$s);
                    RETURN NEW;
                END;
            $$ LANGUAGE plpgsql;
            DROP TRIGGER IF EXISTS trigger_create_comment_thread ON %1$s;
            CREATE TRIGGER trigger_create_comment_thread
                AFTER INSERT ON %1$s
                FOR EACH ROW
                EXECUTE FUNCTION create_comment_thread_%1$s();

            CREATE OR REPLACE FUNCTION delete_comment_thread_%1$s()
                RETURNS trigger AS
            $$
                BEGIN
                    DELETE FROM comment_threads AS ct
                        USING %3$s AS rt
                    WHERE rt.resource_id = OLD.%2$s
                        AND rt.thread_id = ct.id;
                    RETURN OLD;
                END;
            $$ LANGUAGE plpgsql;
            DROP TRIGGER IF EXISTS trigger_delete_comment_thread ON %1$s;
            CREATE TRIGGER trigger_delete_comment_thread
                BEFORE DELETE ON %1$s
                FOR EACH ROW
                EXECUTE FUNCTION delete_comment_thread_%1$s();
        $COM$
            , table_name
            , pk1_col
            , comment_thread_table_name
        );
    END;
$BODY$ LANGUAGE plpgsql;

DO $$
BEGIN
    PERFORM create_comment_thread_for_table('risks', 'id');
    PERFORM create_comment_thread_for_table('controls', 'id');
END $$;
