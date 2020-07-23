DO $$
BEGIN
    PERFORM create_comment_thread_for_table('gl_accounts', 'id');
END $$;
