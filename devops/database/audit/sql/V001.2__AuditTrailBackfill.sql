DO $$
BEGIN
    PERFORM audit.create_audit_trail_triggers_for_table('users', 'id');
    PERFORM audit.create_audit_trail_triggers_for_table('user_orgs', 'id');
    PERFORM audit.create_audit_trail_triggers_for_table('roles', 'id');
    PERFORM audit.create_audit_trail_triggers_for_table('user_roles', 'id');
    PERFORM audit.create_audit_trail_triggers_for_table('organizations', 'id');
    PERFORM audit.create_audit_trail_triggers_for_table('engagements', 'id');
    PERFORM audit.create_audit_trail_triggers_for_table('engagement_roles', 'id');
    PERFORM audit.create_audit_trail_triggers_for_table('engagement_status', 'id');
END $$;
