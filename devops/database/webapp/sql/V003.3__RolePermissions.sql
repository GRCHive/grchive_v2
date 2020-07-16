CREATE TABLE permissions (
    id BIGSERIAL PRIMARY KEY,
    human_name VARCHAR NOT NULL UNIQUE,
    description TEXT
);

INSERT INTO permissions (human_name, description)
VALUES
    ('org.profile.view', 'Ability to view the profile of the organization.'),
    ('org.profile.create', 'Ability to create suborgs in the org.'),
    ('org.profile.update', 'Ability to modify the organization''s profile.'),
    ('org.roles.view', 'Ability to view all of the roles created for the organization.'),
    ('org.roles.update', 'Ability to edit roles created for the organization.'),
    ('org.roles.delete', 'Ability to delete roles created for the organization.'),
    ('org.roles.create', 'Ability to create new roles for the organization.');
