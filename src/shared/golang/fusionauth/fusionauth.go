package fusionauth

import (
	"fmt"
	"github.com/FusionAuth/go-client/pkg/fusionauth"
	"net/http"
	"net/url"
)

type Config struct {
	Host           string
	Port           int32
	ExternalHost   string
	ClientId       string
	ClientSecret   string
	ApiKey         string
	RedirectDomain string
}

type FusionAuthClient struct {
	*fusionauth.FusionAuthClient

	cfg *Config
}

func NewClient(cfg *Config) *FusionAuthClient {
	client := FusionAuthClient{
		FusionAuthClient: fusionauth.NewClient(
			&http.Client{},
			&url.URL{
				Scheme: "http",
				Host:   fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
			},
			cfg.ApiKey,
		),
		cfg: cfg,
	}
	return &client
}
