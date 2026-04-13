path "database/creds/changelock-readwrite" {
  capabilities = ["read"]
}
path "kv/data/changelock/*" {
  capabilities = ["read"]
}
