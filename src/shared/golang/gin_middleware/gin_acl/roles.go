package gin_acl

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/grchive/grchive-v2/shared/backend/roles"
	"gitlab.com/grchive/grchive-v2/shared/gin_middleware/gin_backend_utility"
	"net/http"
)

func (acl *ACLClient) ACLUserHasPermissions(permissions ...roles.Permission) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, err := acl.Middleware.GetCurrentUserFromContext(c)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, &gin_backend_utility.WebappError{
				Err:     err,
				Context: "ACLUserHasPermissions - Get Current User",
			})
			return
		}

		err = acl.Middleware.Itf.Roles.UserHasPermissions(user.Id, permissions...)
		if err != nil {
			c.AbortWithError(http.StatusUnauthorized, &gin_backend_utility.WebappError{
				Err:     err,
				Context: "ACLUserHasPermissions - UserHasPermissions",
			})
			return
		}

		c.Next()
	}

}
