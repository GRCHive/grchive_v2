#!/bin/bash
flyway \
    -user=$POSTGRES_USER \
    -password=$POSTGRES_PASSWORD \
    -url="jdbc:postgresql://$POSTGRES_HOST/grchive" \
    -locations="filesystem:$PWD/sql" \
    migrate
