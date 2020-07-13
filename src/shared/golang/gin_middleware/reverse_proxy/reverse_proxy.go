package reverse_proxy

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httputil"
)

func GinNoop(c *gin.Context) {
	c.Next()
}

func GinReverseProxy(host string, prefix string) gin.HandlerFunc {
	return func(c *gin.Context) {
		proxy := httputil.ReverseProxy{
			Director: func(req *http.Request) {
				req.URL.Scheme = "http"
				req.URL.Host = host
				req.URL.Path = fmt.Sprintf("%s%s", prefix, c.Param("path"))
			},
		}

		proxy.ServeHTTP(c.Writer, c.Request)
	}
}
