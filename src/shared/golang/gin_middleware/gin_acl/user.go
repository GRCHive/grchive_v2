package gin_acl

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/grchive/grchive-v2/shared/gin_middleware/gin_backend_utility"
	"net/http"
)

func (acl *ACLClient) ACLUserEmailVerified(c *gin.Context) {
	user, err := acl.Middleware.GetCurrentUserFromContext(c)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, &gin_backend_utility.WebappError{
			Err:     err,
			Context: "ACLUserEmailVerified - Get user from context.",
			Code:    gin_backend_utility.GECBadRequest,
			Message: gin_backend_utility.GEMBadRequest,
		})
		return
	}

	if !user.EmailVerified {
		c.AbortWithError(http.StatusUnauthorized, &gin_backend_utility.WebappError{
			Err:     nil,
			Context: "ACLUserEmailVerified - User email not verified.",
			Code:    gin_backend_utility.GECUnauthorized,
			Message: gin_backend_utility.GEMUnauthorizedVerify,
		})
		return
	}

	c.Next()
}

func (acl *ACLClient) ACLUserHasValidRole(c *gin.Context) {
	roles, err := acl.Middleware.GetCurrentUserRolesFromContext(c)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, &gin_backend_utility.WebappError{
			Err:     err,
			Context: "ACLUserHasValidRole - Get roles.",
			Code:    gin_backend_utility.GECBadRequest,
			Message: gin_backend_utility.GEMBadRequest,
		})
		return
	}

	if len(roles) == 0 {
		c.AbortWithError(http.StatusUnauthorized, &gin_backend_utility.WebappError{
			Err:     nil,
			Context: "ACLUserHasValidRole - No roles.",
			Code:    gin_backend_utility.GECUnauthorized,
			Message: gin_backend_utility.GEMUnauthorizedRoles,
		})
		return
	}

	c.Next()
}
