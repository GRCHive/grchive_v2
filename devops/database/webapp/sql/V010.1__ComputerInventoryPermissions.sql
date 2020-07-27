INSERT INTO permissions (human_name, description)
VALUES
    ('org.inventory.servers.list', 'Ability to list all of the servers for an engagement.'),
    ('org.inventory.servers.create', 'Ability to create new servers for an engagement.'),
    ('org.inventory.servers.view', 'Ability to view a server.'),
    ('org.inventory.servers.delete', 'Ability to delete a server.'),
    ('org.inventory.servers.update', 'Ability to update a server.'),
    ('org.inventory.desktops.list', 'Ability to list all of the desktops for an engagement.'),
    ('org.inventory.desktops.create', 'Ability to create new desktops for an engagement.'),
    ('org.inventory.desktops.view', 'Ability to view a desktop.'),
    ('org.inventory.desktops.delete', 'Ability to delete a desktop.'),
    ('org.inventory.desktops.update', 'Ability to update a desktop.'),
    ('org.inventory.laptops.list', 'Ability to list all of the laptops for an engagement.'),
    ('org.inventory.laptops.create', 'Ability to create new laptops for an engagement.'),
    ('org.inventory.laptops.view', 'Ability to view a laptop.'),
    ('org.inventory.laptops.delete', 'Ability to delete a laptop.'),
    ('org.inventory.laptops.update', 'Ability to update a laptop.'),
    ('org.inventory.mobile.list', 'Ability to list all of the mobile devices for an engagement.'),
    ('org.inventory.mobile.create', 'Ability to create new mobile devices for an engagement.'),
    ('org.inventory.mobile.view', 'Ability to view a mobile device.'),
    ('org.inventory.mobile.delete', 'Ability to delete a mobile device.'),
    ('org.inventory.mobile.update', 'Ability to update a mobile device.'),
    ('org.inventory.embedded.list', 'Ability to list all of the embedded devices for an engagement.'),
    ('org.inventory.embedded.create', 'Ability to create new embedded devices for an engagement.'),
    ('org.inventory.embedded.view', 'Ability to view a embedded device.'),
    ('org.inventory.embedded.delete', 'Ability to delete a embedded device.'),
    ('org.inventory.embedded.update', 'Ability to update a embedded device.');

DO $$
BEGIN
    PERFORM update_admin_roles_with_all_permissions(r.id)
    FROM roles AS r
    WHERE r.is_admin_role = true;
END $$;
