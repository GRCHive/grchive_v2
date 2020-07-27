package gin_backend_utility

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"net/http"
	"time"
)

func CreateLogger(lg zerolog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestTime := time.Now().UTC()
		path := c.Request.URL.String()
		ip := c.ClientIP()
		method := c.Request.Method
		userAgent := c.Request.UserAgent()
		requestId := c.GetHeader(requestIdHeaderKey)

		c.Next()

		status := c.Writer.Status()
		finishTime := time.Now().UTC()
		requestDuration := finishTime.Sub(requestTime)

		logger := lg.With().
			Time("now", requestTime).
			Str("request-id", requestId).
			Int("status", status).
			Str("method", method).
			Str("path", path).
			Str("ip", ip).
			Dur("latency", requestDuration).
			Str("user-agent", userAgent).
			Logger()

		var evt *zerolog.Event
		if status >= http.StatusInternalServerError {
			evt = logger.Error()
		} else if status >= http.StatusBadRequest {
			evt = logger.Warn()
		} else {
			evt = logger.Info()
		}

		evt.Msg(c.Errors.String())
	}
}
