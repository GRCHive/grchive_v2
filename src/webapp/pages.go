package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gitlab.com/grchive/grchive-v2/shared/zipfs"
	"html/template"
	"io/ioutil"
	"net/http"
)

type Pages struct {
	jsCssLoaderTemplate string
	appTemplate         *template.Template
}

type templateKey int

const (
	AppPageTempateKey templateKey = iota
)

func (w *WebappApplication) LoadPages() {
	w.pages = &Pages{}

	var err error
	w.clientZipFs, err = zipfs.CreateZipFSFromZipFile("", "webappClient", "src/webapp/client/webappClient.zip")
	if err != nil {
		w.log.Fatal().Err(err).Msg("Failed to open webapp client zip.")
	}

	{
		tmplFile, err := w.clientZipFs.Open("/main.tmpl")
		if err != nil {
			w.log.Fatal().Err(err).Msg("Failed to open CSS JS loader template.")
		}

		data, err := ioutil.ReadAll(tmplFile)
		if err != nil {
			w.log.Fatal().Err(err).Msg("Failed to read CSS JS loader template.")
		}

		w.pages.jsCssLoaderTemplate = string(data)
	}

	w.pages.appTemplate = template.Must(template.Must(
		template.New("").
			Delims("[[", "]]").
			ParseFiles(
				"src/webapp/templates/base.tmpl",
			)).Parse(w.pages.jsCssLoaderTemplate))
}

func (w *WebappApplication) renderApp(r *gin.Context) {
	w.pages.appTemplate.ExecuteTemplate(r.Writer, "base", nil)
}

func (w *WebappApplication) renderLogin(r *gin.Context) {
	r.Redirect(
		http.StatusTemporaryRedirect,
		w.cfg.FusionAuth.ExternalHost+fmt.Sprintf(w.cfg.FusionAuth.LoginEndpoint, w.cfg.FusionAuth.ClientId, w.cfg.Grchive.Domain),
	)
}
