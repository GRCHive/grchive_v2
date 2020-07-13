package auth

import (
	"net/http"
)

type VaultTokenManager interface {
	Initialize(url string, client *http.Client)
	Token() string
}
