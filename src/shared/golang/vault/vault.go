package vault

import (
	"fmt"
	"gitlab.com/grchive/grchive-v2/shared/vault/auth"
	"net/http"
)

type Config struct {
	Protocol string
	Host     string
	Port     int32
}

type Client struct {
	Url      string
	tokenMgr auth.VaultTokenManager
	client   *http.Client
}

func NewClient(c *Config, mgr auth.VaultTokenManager) *Client {
	client := &Client{
		Url:      fmt.Sprintf("%s://%s:%d", c.Protocol, c.Host, c.Port),
		client:   &http.Client{},
		tokenMgr: mgr,
	}

	client.tokenMgr.Initialize(client.Url, client.client)
	return client
}
