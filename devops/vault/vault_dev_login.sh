#!/bin/bash
export VAULT_SKIP_VERIFY=1
VAULT_TOKEN=$1
vault login -address="${VAULT_HOST}:8200" token=${VAULT_TOKEN}
