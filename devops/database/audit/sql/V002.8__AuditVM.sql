DO $$
BEGIN
    PERFORM audit.create_audit_trail_triggers_for_table('virtual_machines', 'id');
END $$;
