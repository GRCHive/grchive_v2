package main

import (
	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"gitlab.com/grchive/grchive-v2/shared/zipfs"
	"os"
)

type WebsiteApplication struct {
	cfg                *Config
	log                zerolog.Logger
	websiteClientZipFs *zipfs.ZipFS

	// Raw text of the template file that contains what Javascript/CSS to load.
	pages *Pages
}

func (w *WebsiteApplication) loadConfig(fname string) {
	w.cfg = w.loadConfigFromFile(fname)
}

func (w *WebsiteApplication) setupLogging() {
	// TODO: #9cn9hj
	// Setup logger to output raw JSON to a central repository for machine analysis.
	// This setup should be shared across all apps probably.
	console := zerolog.ConsoleWriter{Out: os.Stdout}
	multi := zerolog.MultiLevelWriter(console)
	w.log = zerolog.New(multi).With().Timestamp().Logger()
}

func (w *WebsiteApplication) run() {
	defer w.websiteClientZipFs.Close()

	r := gin.New()
	r.Use(logger.SetLogger(logger.Config{
		Logger: &w.log,
	}))
	r.Use(gin.Recovery())

	r.Static("/static/assets", "static/assets")
	r.StaticFile("/robots.txt", "src/website/robots.txt")
	r.StaticFile("/sitemap.xml", "src/website/sitemap.xml")

	r.StaticFS("/static/client/", w.websiteClientZipFs)

	r.GET("/", w.renderHomePage)
	r.Run()
}
