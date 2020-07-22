CREATE TABLE control_types (
    id SERIAL PRIMARY KEY,
    name VARCHAR NOT NULL UNIQUE,
    description TEXT DEFAULT ''
);

INSERT INTO control_types (id, name)
VALUES 
    (1, 'Access'),
    (2, 'Authorization'),
    (3, 'Configuration'),
    (4, 'GITC'),
    (5, 'Interface'),
    (6, 'Management Review'),
    (7, 'Reconciliation'),
    (8, 'Report/IPE');

CREATE TABLE control_frequency_types (
    id SERIAL PRIMARY KEY,
    name VARCHAR NOT NULL UNIQUE
);

INSERT INTO control_frequency_types (id, name)
VALUES
    (1, 'Ad-Hoc'),
    (2, 'Interval'),
    (3, 'Other');

CREATE TABLE controls (
    id BIGSERIAL PRIMARY KEY,
    engagement_id BIGINT NOT NULL REFERENCES engagements(id) ON DELETE CASCADE,
    name VARCHAR NOT NULL,
    human_id VARCHAR(256) NOT NULL,
    description TEXT NOT NULL DEFAULT '',
    likelihood SMALLINT NOT NULL DEFAULT 0,
    likelihood_justification TEXT,
    control_type INTEGER NOT NULL REFERENCES control_types(id) ON DELETE RESTRICT,
    owner_id BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    is_manual BOOLEAN NOT NULL DEFAULT false,
    freq_type INTEGER NOT NULL REFERENCES control_frequency_types(id) ON DELETE RESTRICT,
    freq_other TEXT NOT NULL DEFAULT '',
    freq_interval TEXT NOT NULL DEFAULT '',
    UNIQUE(engagement_id, human_id)
);
