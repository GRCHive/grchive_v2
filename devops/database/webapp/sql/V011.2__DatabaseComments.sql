DO $$
BEGIN
    PERFORM create_comment_thread_for_table('databases', 'id');
END $$;
