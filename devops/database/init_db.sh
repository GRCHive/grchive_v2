#!/bin/bash
set -xe

# This database will be setup later.
createdb grchive
psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname grchive <<-EOSQL
    CREATE USER $POSTGRES_WEBAPP_USER WITH
        LOGIN
        ENCRYPTED PASSWORD '$POSTGRES_WEBAPP_PASSWORD';
    GRANT CONNECT ON DATABASE grchive TO $POSTGRES_WEBAPP_USER;
    ALTER DEFAULT PRIVILEGES IN SCHEMA public
        GRANT SELECT, INSERT, UPDATE, DELETE ON TABLES TO $POSTGRES_WEBAPP_USER;

    ALTER DEFAULT PRIVILEGES IN SCHEMA public
        GRANT USAGE, SELECT ON SEQUENCES TO $POSTGRES_WEBAPP_USER;
EOSQL

# These databases needs to be setup now.
createdb vault
createdb fusionauth

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname vault <<-EOSQL
    CREATE USER vault WITH
        LOGIN
        ENCRYPTED PASSWORD '$VAULT_USER_PASSWORD';
    GRANT ALL PRIVILEGES ON DATABASE vault TO vault;
EOSQL
