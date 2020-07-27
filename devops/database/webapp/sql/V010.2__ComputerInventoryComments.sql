DO $$
BEGIN
    PERFORM create_comment_thread_for_table('inventory_servers', 'id');
    PERFORM create_comment_thread_for_table('inventory_desktops', 'id');
    PERFORM create_comment_thread_for_table('inventory_laptops', 'id');
    PERFORM create_comment_thread_for_table('inventory_mobile', 'id');
    PERFORM create_comment_thread_for_table('inventory_embedded', 'id');
END $$;
