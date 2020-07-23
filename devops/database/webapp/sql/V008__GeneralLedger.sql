CREATE TABLE gl_account_types (
    id SERIAL PRIMARY KEY,
    name VARCHAR UNIQUE
);

INSERT INTO gl_account_types (id, name)
VALUES
    (1, 'Asset'),
    (2, 'Liability'),
    (3, 'Equity'),
    (4, 'Revenue'),
    (5, 'Expense'),
    (6, 'Contra');

CREATE TABLE gl_accounts (
    id BIGSERIAL PRIMARY KEY,
    engagement_id BIGINT NOT NULL REFERENCES engagements(id) ON DELETE CASCADE,
    name VARCHAR NOT NULL,
    description TEXT NOT NULL DEFAULT '',
    account_id VARCHAR NOT NULL,
    account_type INTEGER NOT NULL REFERENCES gl_account_types(id) ON DELETE RESTRICT,
    parent_account_id BIGINT REFERENCES gl_accounts(id) ON DELETE CASCADE,
    financially_relevant BOOLEAN NOT NULL DEFAULT false,
    UNIQUE(engagement_id, account_id)
);
