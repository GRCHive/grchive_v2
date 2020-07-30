CREATE OR REPLACE FUNCTION delete_machine_state_for_inventory()
    RETURNS trigger AS
$$
    BEGIN
        DELETE FROM machine_state
        WHERE id = OLD.state_id;
        RETURN OLD;
    END;
$$ LANGUAGE plpgsql;

DROP TRIGGER IF EXISTS trigger_delete_machine_state ON inventory_servers;
CREATE TRIGGER trigger_delete_machine_state
    AFTER DELETE ON inventory_servers
    FOR EACH ROW
    EXECUTE FUNCTION delete_machine_state_for_inventory();

DROP TRIGGER IF EXISTS trigger_delete_machine_state ON inventory_desktops;
CREATE TRIGGER trigger_delete_machine_state
    AFTER DELETE ON inventory_desktops
    FOR EACH ROW
    EXECUTE FUNCTION delete_machine_state_for_inventory();

DROP TRIGGER IF EXISTS trigger_delete_machine_state ON inventory_laptops;
CREATE TRIGGER trigger_delete_machine_state
    AFTER DELETE ON inventory_laptops
    FOR EACH ROW
    EXECUTE FUNCTION delete_machine_state_for_inventory();

DROP TRIGGER IF EXISTS trigger_delete_machine_state ON inventory_mobile;
CREATE TRIGGER trigger_delete_machine_state
    AFTER DELETE ON inventory_mobile
    FOR EACH ROW
    EXECUTE FUNCTION delete_machine_state_for_inventory();

DROP TRIGGER IF EXISTS trigger_delete_machine_state ON inventory_embedded;
CREATE TRIGGER trigger_delete_machine_state
    AFTER DELETE ON inventory_embedded
    FOR EACH ROW
    EXECUTE FUNCTION delete_machine_state_for_inventory();
