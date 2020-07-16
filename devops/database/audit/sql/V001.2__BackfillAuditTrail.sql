DO $$
BEGIN
    PERFORM audit.create_audit_trail_triggers_for_table('users', NULL, 'id');
    PERFORM audit.create_audit_trail_triggers_for_table('organizations', 'parent_org_id', 'id');
END $$;
