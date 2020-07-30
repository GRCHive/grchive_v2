DO $$
BEGIN
    PERFORM create_comment_thread_for_table('virtual_machines', 'id');
END $$;
