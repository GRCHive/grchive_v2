package main

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/grchive/grchive-v2/shared/zipfs"
	"html/template"
	"io/ioutil"
)

type templateKey int
type templateMap map[templateKey]*template.Template

const (
	GettingStartedPageTemplateKey templateKey = iota
	ContactUsPageTemplateKey
	LandingPageTemplateKey
	LearnMorePageTemplateKey
)

type Pages struct {
	jsCssLoaderTemplate string
	allTemplates        templateMap
}

func (p *Pages) loadPage(key templateKey, tmplFname string) {
	p.allTemplates[key] = template.Must(template.Must(
		template.New("").
			Delims("[[", "]]").
			ParseFiles(
				"src/website/templates/base.tmpl",
				tmplFname,
			)).Parse(p.jsCssLoaderTemplate))
}

func (w *WebsiteApplication) loadPages() {
	w.pages = &Pages{
		allTemplates: make(templateMap),
	}

	var err error
	w.websiteClientZipFs, err = zipfs.CreateZipFSFromZipFile("", "websiteClient", "src/website/client/websiteClient.zip")
	if err != nil {
		w.log.Fatal().Err(err).Msg("Failed to open website client zip.")
	}

	{
		tmplFile, err := w.websiteClientZipFs.Open("/main.tmpl")
		if err != nil {
			w.log.Fatal().Err(err).Msg("Failed to open CSS JS loader template.")
		}

		data, err := ioutil.ReadAll(tmplFile)
		if err != nil {
			w.log.Fatal().Err(err).Msg("Failed to read CSS JS loader template.")
		}

		w.pages.jsCssLoaderTemplate = string(data)
	}

	w.pages.loadPage(GettingStartedPageTemplateKey, "src/website/templates/gettingStarted.tmpl")
	w.pages.loadPage(LandingPageTemplateKey, "src/website/templates/index.tmpl")
	w.pages.loadPage(LearnMorePageTemplateKey, "src/website/templates/learnMore.tmpl")
	w.pages.loadPage(ContactUsPageTemplateKey, "src/website/templates/contactUs.tmpl")
}

func (w *WebsiteApplication) renderGenericPage(key templateKey, r *gin.Context) {
	tmpl, ok := w.pages.allTemplates[key]
	if !ok {
		w.log.Error().Msgf("Failed to find template: %d.", key)
		return
	}

	tmpl.ExecuteTemplate(r.Writer, "base", map[string]interface{}{
		"UseAnalytics": w.cfg.EnableAnalytics,
	})
}

func (w *WebsiteApplication) renderHomePage(r *gin.Context) {
	w.renderGenericPage(LandingPageTemplateKey, r)
}
