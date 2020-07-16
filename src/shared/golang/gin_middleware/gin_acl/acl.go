package gin_acl

import (
	"gitlab.com/grchive/grchive-v2/shared/gin_middleware/gin_backend_utility"
)

type ACLClient struct {
	Middleware *gin_backend_utility.MiddlewareClient
}
