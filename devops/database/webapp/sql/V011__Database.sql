CREATE TABLE database_types (
    id SERIAL PRIMARY KEY,
    name VARCHAR NOT NULL UNIQUE
);

INSERT INTO database_types (id, name)
VALUES 
    (1, 'Other'),
    (2, 'PostgreSQL'),
    (3, 'MySQL'),
    (4, 'MariaDB'),
    (5, 'SAP HANA'),
    (6, 'IBM DB2'),
    (7, 'Oracle'),
    (8, 'Microsoft SQL'),
    (9, 'Microsoft Access'),
    (10, 'SQLite'),
    (11, 'H2');

CREATE TABLE databases (
    id BIGSERIAL PRIMARY KEY,
    engagement_id BIGINT NOT NULL REFERENCES engagements(id) ON DELETE CASCADE,
    unique_id VARCHAR NOT NULL,
    name VARCHAR NOT NULL,
    description TEXT NOT NULL DEFAULT '',
    type_id INTEGER NOT NULL REFERENCES database_types(id) ON DELETE RESTRICT,
    other_type TEXT,
    version TEXT NOT NULL DEFAULT '',
    UNIQUE(engagement_id, unique_id)
);
