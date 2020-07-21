DO $$
BEGIN
    PERFORM audit.create_audit_trail_triggers_for_table('risks', 'id');
END $$;
