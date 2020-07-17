package gin_acl

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/grchive/grchive-v2/shared/backend"
	"gitlab.com/grchive/grchive-v2/shared/backend/orgs"
	"gitlab.com/grchive/grchive-v2/shared/backend/roles"
	"gitlab.com/grchive/grchive-v2/shared/gin_middleware/gin_backend_utility"
	"net/http"
)

func (acl *ACLClient) ACLCheckPermissionHandler(c *gin.Context, onSuccess func(), onError func(error), permissions ...roles.Permission) {
	user, err := acl.Middleware.GetCurrentUserFromContext(c)
	if err != nil {
		onError(err)
		return
	}

	org, err := acl.Middleware.GetResourceFromContext(c, backend.RIOrganization)
	if err != nil {
		onError(err)
		return
	}

	err = acl.Middleware.Itf.Roles.UserHasPermissions(user.Id, org.(*orgs.Organization).Id, permissions...)
	if err != nil {
		onError(err)
		return
	}

	onSuccess()
}

func (acl *ACLClient) ACLUserHasPermissions(permissions ...roles.Permission) gin.HandlerFunc {
	return func(c *gin.Context) {
		acl.ACLCheckPermissionHandler(c, func() {
			c.Next()
		}, func(err error) {
			c.AbortWithError(http.StatusUnauthorized, &gin_backend_utility.WebappError{
				Err:     err,
				Context: "Check user has permission",
			})
		}, permissions...)
	}
}

func (acl *ACLClient) ACLBranchOnPermissions(onSuccess gin.HandlerFunc, onFail gin.HandlerFunc, permissions ...roles.Permission) gin.HandlerFunc {
	return func(c *gin.Context) {
		acl.ACLCheckPermissionHandler(c, func() {
			onSuccess(c)
		}, func(err error) {
			onFail(c)
		}, permissions...)
	}
}
