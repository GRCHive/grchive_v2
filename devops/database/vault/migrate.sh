#!/bin/bash
flyway \
    -user=vault \
    -password=$VAULT_USER_PASSWORD \
    -url="jdbc:postgresql://$POSTGRES_HOST/vault" \
    -locations="filesystem:$PWD/sql" \
    migrate
