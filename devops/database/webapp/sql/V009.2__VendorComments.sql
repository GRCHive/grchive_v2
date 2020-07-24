DO $$
BEGIN
    PERFORM create_comment_thread_for_table('vendors', 'id');
    PERFORM create_comment_thread_for_table('vendor_products', 'id');
END $$;
