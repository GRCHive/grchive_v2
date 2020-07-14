package main

import (
	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gitlab.com/grchive/grchive-v2/shared/fusionauth"
	"gitlab.com/grchive/grchive-v2/shared/gin_middleware/redirect_response"
	"gitlab.com/grchive/grchive-v2/shared/vault"
	"gitlab.com/grchive/grchive-v2/shared/vault/auth"
	"gitlab.com/grchive/grchive-v2/shared/zipfs"
	"gitlab.com/grchive/grchive-v2/webapp/session"
	"os"
)

type WebappApplication struct {
	// cfg contains values loaded either from the
	// config file or from environment variables.
	cfg *Config

	log   zerolog.Logger
	pages *Pages

	clientZipFs *zipfs.ZipFS

	vaultClient *vault.Client
	fusionauth  *fusionauth.FusionAuthClient

	sessionStore *session.SessionStore
}

func (w *WebappApplication) SetupLogging() {
	// TODO: #9cn9hj
	// Setup logger to output raw JSON to a central repository for machine analysis.
	// This setup should be shared across all apps probably.
	console := zerolog.ConsoleWriter{Out: os.Stdout}
	multi := zerolog.MultiLevelWriter(console)
	w.log = zerolog.New(multi).With().Timestamp().Logger()
	log.Logger = w.log
}

func (w *WebappApplication) Run() {
	w.sessionStore = session.CreateSessionStore(w.cfg.Grchive.SessionAuthKey, w.cfg.Grchive.SessionEncryptKey)

	r := gin.New()
	r.Use(logger.SetLogger(logger.Config{
		Logger: &w.log,
	}))
	r.Use(gin.Recovery())
	r.Use(redirect_response.GinHTTPRedirectStatusCodes)

	r.Static("/static/assets", "static/assets")
	r.StaticFile("/robots.txt", "src/website/robots.txt")
	r.StaticFS("/static/client/", w.clientZipFs)

	r.Use(w.sessionStore.ValidateLogin(w.fusionauth))
	loginR := r.Group("/", w.sessionStore.RedirectIfLoggedIn("/"))
	loginR.GET("/login", w.sessionStore.PopulateLoginState, w.renderLogin)
	loginR.GET("/oauth2callback", w.handleOauthCallback)

	appR := r.Group("/", w.sessionStore.RedirectIfLoggedOut("/login"))
	appR.GET("/", w.renderApp)

	r.Run()
}

func (w *WebappApplication) InitializeBackend() {
	w.vaultClient = vault.NewClient(&vault.Config{
		Protocol: w.cfg.Vault.Protocol,
		Host:     w.cfg.Vault.Host,
		Port:     w.cfg.Vault.Port,
	}, &auth.AppRoleTokenManager{
		RoleId:   w.cfg.Vault.RoleId,
		SecretId: w.cfg.Vault.SecretId,
	})

	w.fusionauth = fusionauth.NewClient(&fusionauth.Config{
		Host:           w.cfg.FusionAuth.Host,
		Port:           w.cfg.FusionAuth.Port,
		ExternalHost:   w.cfg.FusionAuth.ExternalHost,
		ClientId:       w.cfg.FusionAuth.ClientId,
		ClientSecret:   w.cfg.FusionAuth.ClientSecret,
		ApiKey:         w.cfg.FusionAuth.ApiKey,
		RedirectDomain: w.cfg.Grchive.Domain,
	})
}

func (w *WebappApplication) Close() {
	w.clientZipFs.Close()
}
