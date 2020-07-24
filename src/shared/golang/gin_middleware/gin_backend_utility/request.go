package gin_backend_utility

import (
	"github.com/gin-gonic/gin"
)

const requestIdHeaderKey = "Grchive-Request-Id"

func PingPongGrchiveRequestId(c *gin.Context) {
	val := c.GetHeader(requestIdHeaderKey)
	c.Header(requestIdHeaderKey, val)
	c.Next()
}
