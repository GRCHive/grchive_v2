CREATE OR REPLACE FUNCTION create_machine_state_for_vm()
    RETURNS trigger AS
$$
    DECLARE
        new_state_id BIGINT;
    BEGIN
        INSERT INTO machine_state (
            engagement_id,
            unique_id
        )
        VALUES (
            NEW.engagement_id,
            gen_random_uuid()::TEXT
        )
        RETURNING id INTO new_state_id;

        NEW.state_id = new_state_id;
        RETURN NEW;
    END;
$$ LANGUAGE plpgsql;

DROP TRIGGER IF EXISTS trigger_create_machine_state ON virtual_machines;
CREATE TRIGGER trigger_create_machine_state
    BEFORE INSERT ON virtual_machines
    FOR EACH ROW
    EXECUTE FUNCTION create_machine_state_for_vm();

CREATE OR REPLACE FUNCTION delete_machine_state_for_vm()
    RETURNS trigger AS
$$
    BEGIN
        DELETE FROM machine_state
        WHERE id = OLD.state_id;
        RETURN OLD;
    END;
$$ LANGUAGE plpgsql;

DROP TRIGGER IF EXISTS trigger_delete_machine_state ON virtual_machines;
CREATE TRIGGER trigger_delete_machine_state
    AFTER DELETE ON virtual_machines
    FOR EACH ROW
    EXECUTE FUNCTION delete_machine_state_for_vm();
