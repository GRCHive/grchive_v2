package redirect_response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// This response writer will change the original http.ResponseWriter so that
// it writes a status code of 301 -> 308 and 302 -> 307.
type RedirectResponseWriter struct {
	gin.ResponseWriter
}

func (this RedirectResponseWriter) WriteHeader(statusCode int) {
	if statusCode == http.StatusMovedPermanently {
		this.ResponseWriter.WriteHeader(http.StatusPermanentRedirect)
	} else if statusCode == http.StatusFound {
		this.ResponseWriter.WriteHeader(http.StatusTemporaryRedirect)
	} else {
		this.ResponseWriter.WriteHeader(statusCode)
	}
}
