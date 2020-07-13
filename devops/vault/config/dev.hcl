storage "postgresql" {
    connection_url = "postgres://vault:${VAULT_USER_PASSWORD}@${POSTGRES_HOST}:5432/vault?sslmode=disable&timezone=UTC"
    table = "vault_kv_store"
}

listener "tcp" {
    address     = "0.0.0.0:8200"
    tls_disable = 1
}
