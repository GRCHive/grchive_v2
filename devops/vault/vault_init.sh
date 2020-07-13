VAULT_ADDR=${VAULT_PROTOCOL}://${VAULT_HOST}:8200

vault secrets enable -address="${VAULT_ADDR}" transit
vault secrets enable -address="${VAULT_ADDR}" -path=secret kv-v2
