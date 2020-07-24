INSERT INTO permissions (human_name, description)
VALUES
    ('org.vendors.list', 'Ability to list all of the vendors for an engagement.'),
    ('org.vendors.create', 'Ability to create new vendors for an engagement.'),
    ('org.vendors.view', 'Ability to view a vendor.'),
    ('org.vendors.delete', 'Ability to delete a vendor.'),
    ('org.vendors.update', 'Ability to update a vendor.'),
    ('org.vendors.products.list', 'Ability to list all of the products for a vendor.'),
    ('org.vendors.products.create', 'Ability to create new product for a vendor.'),
    ('org.vendors.products.view', 'Ability to view a vendor product.'),
    ('org.vendors.products.delete', 'Ability to delete a vendor product.'),
    ('org.vendors.products.update', 'Ability to update a vendor product.');

DO $$
BEGIN
    PERFORM update_admin_roles_with_all_permissions(r.id)
    FROM roles AS r
    WHERE r.is_admin_role = true;
END $$;
