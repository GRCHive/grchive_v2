DO $$
BEGIN
    PERFORM audit.create_audit_trail_triggers_for_table('user_orgs', 'org_id', 'id');
    PERFORM audit.create_audit_trail_triggers_for_table('roles', 'org_id', 'id');
    PERFORM audit.create_audit_trail_triggers_for_table('user_roles', 'org_id', 'id');
END $$;
