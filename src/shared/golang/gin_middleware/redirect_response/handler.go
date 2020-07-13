package redirect_response

import (
	"github.com/gin-gonic/gin"
)

// Redirect non-GET 301 and 302 to 308 and 307 respectively. The 307 and 308
// status codes will pass the original body and method to the redirected path
// which is desirable behavior for us whereas 301 and 302 will generally always
// redirect to a GET.
// See:
// 	301: https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/301
//  302: https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/302
//  307: https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/307
//  308: https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/308
func GinHTTPRedirectStatusCodes(r *gin.Context) {
	if r.Request.Method == "GET" {
		r.Writer = RedirectResponseWriter{r.Writer}
	}
	r.Next()
}
