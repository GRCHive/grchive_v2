CREATE EXTENSION IF NOT EXISTS pgcrypto;
CREATE OR REPLACE FUNCTION create_machine_state_for_inventory()
    RETURNS trigger AS
$$
    DECLARE
        new_state_id BIGINT;
    BEGIN
        INSERT INTO machine_state (
            engagement_id,
            unique_id
        )
        SELECT engagement_id, gen_random_uuid()::TEXT
        FROM inventory
        WHERE id = NEW.inventory_id
        RETURNING id INTO new_state_id;

        NEW.state_id = new_state_id;
        RETURN NEW;
    END;
$$ LANGUAGE plpgsql;

DROP TRIGGER IF EXISTS trigger_create_machine_state ON inventory_servers;
CREATE TRIGGER trigger_create_machine_state
    BEFORE INSERT ON inventory_servers
    FOR EACH ROW
    EXECUTE FUNCTION create_machine_state_for_inventory();

DROP TRIGGER IF EXISTS trigger_create_machine_state ON inventory_desktops;
CREATE TRIGGER trigger_create_machine_state
    BEFORE INSERT ON inventory_desktops
    FOR EACH ROW
    EXECUTE FUNCTION create_machine_state_for_inventory();

DROP TRIGGER IF EXISTS trigger_create_machine_state ON inventory_laptops;
CREATE TRIGGER trigger_create_machine_state
    BEFORE INSERT ON inventory_laptops
    FOR EACH ROW
    EXECUTE FUNCTION create_machine_state_for_inventory();

DROP TRIGGER IF EXISTS trigger_create_machine_state ON inventory_mobile;
CREATE TRIGGER trigger_create_machine_state
    BEFORE INSERT ON inventory_mobile
    FOR EACH ROW
    EXECUTE FUNCTION create_machine_state_for_inventory();

DROP TRIGGER IF EXISTS trigger_create_machine_state ON inventory_embedded;
CREATE TRIGGER trigger_create_machine_state
    BEFORE INSERT ON inventory_embedded
    FOR EACH ROW
    EXECUTE FUNCTION create_machine_state_for_inventory();
