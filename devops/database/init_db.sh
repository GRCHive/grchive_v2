#!/bin/bash
set -xe

# This database will be setup later.
createdb grchive

# These databases needs to be setup now.
createdb vault
createdb fusionauth

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname vault <<-EOSQL
    CREATE USER vault WITH
        LOGIN
        ENCRYPTED PASSWORD '$VAULT_USER_PASSWORD';
    GRANT ALL PRIVILEGES ON DATABASE vault TO vault;
EOSQL