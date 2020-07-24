DO $$
BEGIN
    PERFORM audit.create_audit_trail_triggers_for_table('vendors', 'id');
    PERFORM audit.create_audit_trail_triggers_for_table('vendor_products', 'id');
END $$;
