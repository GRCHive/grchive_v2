package main

import (
	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gitlab.com/grchive/grchive-v2/shared/backend"
	"gitlab.com/grchive/grchive-v2/shared/fusionauth"
	"gitlab.com/grchive/grchive-v2/shared/gin_middleware/redirect_response"
	"gitlab.com/grchive/grchive-v2/shared/vault"
	"gitlab.com/grchive/grchive-v2/shared/vault/auth"
	"gitlab.com/grchive/grchive-v2/shared/zipfs"
	"gitlab.com/grchive/grchive-v2/webapp/session"
	"os"
	"time"
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

	backend struct {
		db  *sqlx.DB
		itf *backend.BackendInterface
	}
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
	w.sessionStore = session.CreateSessionStore(
		w.cfg.Grchive.SessionAuthKey,
		w.cfg.Grchive.SessionEncryptKey,
		w.backend.itf,
	)
	w.sessionStore.SecureCookies = w.cfg.EnableReleaseMode

	r := gin.New()
	r.RedirectTrailingSlash = true
	r.RedirectFixedPath = true

	r.Use(logger.SetLogger(logger.Config{
		Logger: &w.log,
	}))
	r.Use(gin.Recovery())
	r.Use(redirect_response.GinHTTPRedirectStatusCodes)

	r.Static("/static/assets", "static/assets")
	r.StaticFile("/robots.txt", "src/website/robots.txt")
	r.StaticFS("/static/client/", w.clientZipFs)

	r.Use(w.sessionStore.ValidateLogin(w.fusionauth))
	r.Use(w.sessionStore.SyncEmailVerification(w.fusionauth))

	loginR := r.Group("/", w.sessionStore.RedirectIfLoggedIn("/app/"))
	loginR.GET("/", w.sessionStore.RedirectIfLoggedOut("/login"))
	loginR.GET("/login", w.sessionStore.PopulateLoginState, w.renderLogin)
	loginR.GET("/oauth2callback", w.handleOauthCallback)
	r.GET("/logout", w.renderLogout)

	appR := r.Group("/app", w.sessionStore.RedirectIfLoggedOut("/login"))
	appR.GET("/*path", w.renderApp)

	w.registerApiv1(r)
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
		TenantId:       w.cfg.FusionAuth.TenantId,
		RedirectDomain: w.cfg.Grchive.Domain,
	})

	w.backend.db = sqlx.MustConnect("postgres", w.cfg.Database.JdbcUrl())
	w.backend.db.SetMaxOpenConns(10)
	w.backend.db.SetMaxIdleConns(5)
	w.backend.db.SetConnMaxLifetime(5 * time.Minute)

	w.backend.itf = backend.CreateBackendInterface(w.backend.db)
}

func (w *WebappApplication) Close() {
	w.clientZipFs.Close()
}
