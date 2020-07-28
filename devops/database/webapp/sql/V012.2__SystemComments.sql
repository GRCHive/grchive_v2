DO $$
BEGIN
    PERFORM create_comment_thread_for_table('systems', 'id');
END $$;
