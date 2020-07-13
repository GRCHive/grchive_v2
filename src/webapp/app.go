package main

import (
	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gitlab.com/grchive/grchive-v2/shared/vault"
	"gitlab.com/grchive/grchive-v2/shared/vault/auth"
	"gitlab.com/grchive/grchive-v2/shared/zipfs"
	"os"
)

type WebappApplication struct {
	// cfg contains values loaded either from the
	// config file or from environment variables.
	cfg *Config

	vaultClient *vault.Client

	log   zerolog.Logger
	pages *Pages

	clientZipFs *zipfs.ZipFS
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
	r := gin.New()
	r.Use(logger.SetLogger(logger.Config{
		Logger: &w.log,
	}))
	r.Use(gin.Recovery())

	r.Static("/static/assets", "static/assets")
	r.StaticFile("/robots.txt", "src/website/robots.txt")
	r.StaticFS("/static/client/", w.clientZipFs)

	r.GET("/", w.renderApp)
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
}

func (w *WebappApplication) Close() {
	w.clientZipFs.Close()
}
