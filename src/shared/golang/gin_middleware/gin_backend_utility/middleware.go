package gin_backend_utility

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gitlab.com/grchive/grchive-v2/shared/backend"
	"gitlab.com/grchive/grchive-v2/shared/backend/orgs"
	"gitlab.com/grchive/grchive-v2/shared/backend/roles"
	"gitlab.com/grchive/grchive-v2/shared/backend/users"
	"net/http"
)

type MiddlewareClient struct {
	Itf *backend.BackendInterface
}

// Store user into context differently because we aren't necessarily going to be getting the
// user ID via path parameters which is what LoadResourceIntoContext assumes.
func (m *MiddlewareClient) StoreCurrentUserIntoContext(c *gin.Context, user *users.User) error {
	if user == nil {
		return errors.New("Can not set nil currenet user into context.")
	}

	c.Set(currentUserContextKey, user)
	return nil
}

func (m *MiddlewareClient) GetCurrentUserFromContext(c *gin.Context) (*users.User, error) {
	val, ok := c.Get(currentUserContextKey)
	if !ok {
		return nil, errors.New("Could not find current user.")
	}

	return val.(*users.User), nil
}

func (m *MiddlewareClient) LoadResourceIntoContext(resource backend.ResourceIdentifier, queryStr string) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := c.Param(queryStr)
		obj, err := m.Itf.GetResource(resource, key)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, &WebappError{
				Err:     err,
				Context: "LoadResourceIntoContext - Get Resource",
			})
			return
		}
		c.Set(resource.Str(), obj)
		c.Next()
	}
}

func (m *MiddlewareClient) LoadCurrentUserRolesIntoContext(c *gin.Context) {
	org, err := m.GetResourceFromContext(c, backend.RIOrganization)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, &WebappError{
			Err:     err,
			Context: "LoadCurrentUserRolesIntoContext - Get org",
		})
		return
	}

	user, err := m.GetCurrentUserFromContext(c)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, &WebappError{
			Err:     err,
			Context: "LoadCurrentUserRolesIntoContext - Get current user",
		})
		return
	}

	roles, err := m.Itf.Roles.GetUserRolesForOrg(user.Id, org.(*orgs.Organization).Id)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, &WebappError{
			Err:     err,
			Context: "LoadCurrentUserRolesIntoContext - Get user roles for org",
		})
		return
	}

	c.Set(currentUserRolesContextKey, roles)
	c.Next()
}

func (m *MiddlewareClient) GetCurrentUserRolesFromContext(c *gin.Context) ([]*roles.Role, error) {
	val, ok := c.Get(currentUserRolesContextKey)
	if !ok {
		return nil, errors.New("Could not find current user roles in context.")
	}
	return val.([]*roles.Role), nil
}

func (m *MiddlewareClient) GetResourceFromContext(c *gin.Context, resource backend.ResourceIdentifier) (interface{}, error) {
	val, ok := c.Get(resource.Str())
	if !ok {
		return nil, errors.New("Could not find resource.")
	}
	return val, nil
}
