DO $$
BEGIN
    PERFORM audit.create_audit_trail_triggers_for_table('inventory', 'id');
    PERFORM audit.create_audit_trail_triggers_for_table('inventory_servers', 'id');
    PERFORM audit.create_audit_trail_triggers_for_table('inventory_desktops', 'id');
    PERFORM audit.create_audit_trail_triggers_for_table('inventory_laptops', 'id');
    PERFORM audit.create_audit_trail_triggers_for_table('inventory_mobile', 'id');
    PERFORM audit.create_audit_trail_triggers_for_table('inventory_embedded', 'id');
    PERFORM audit.create_audit_trail_triggers_for_table('inventory_owners', 'id');
END $$;
