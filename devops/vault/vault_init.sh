VAULT_ADDR=${VAULT_PROTOCOL}://${VAULT_HOST}:8200

vault secrets enable -address="${VAULT_ADDR}" transit
vault secrets enable -address="${VAULT_ADDR}" -path=secret kv-v2

vault policy write -address=$VAULT_ADDR webapp policies/webapp-policy.hcl

vault auth enable -address=$VAULT_ADDR approle
vault write -address=$VAULT_ADDR auth/approle/role/webapp \
    token_policies="webapp" \
    token_num_uses=0 \
    secret_id_num_uses=0
