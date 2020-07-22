package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gitlab.com/grchive/grchive-v2/shared/backend"
	"gitlab.com/grchive/grchive-v2/shared/backend/engagements"
	"gitlab.com/grchive/grchive-v2/shared/backend/orgs"
	"gitlab.com/grchive/grchive-v2/shared/backend/risks"
	"gitlab.com/grchive/grchive-v2/shared/backend/roles"
	"gitlab.com/grchive/grchive-v2/shared/gin_middleware/gin_backend_utility"
	"net/http"
)

func (w *WebappApplication) addCommentEndpoints(resource backend.ResourceIdentifier, r *gin.RouterGroup, handlers ...gin.HandlerFunc) {
	commentR := r.Group("/comments", handlers...)
	{
		commentR.GET("/",
			w.acl.ACLUserHasPermissions(roles.PCommentsList),
			w.createResourceCommentsList(resource),
		)
	}
}

func (w *WebappApplication) createResourceCommentsList(resource backend.ResourceIdentifier) gin.HandlerFunc {
	return func(c *gin.Context) {
		org, err := w.middleware.GetResourceFromContext(c, backend.RIOrganization)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, &gin_backend_utility.WebappError{
				Err:     err,
				Context: "createResourceCommentsList - Obtain org in context",
			})
			return
		}

		engagement, err := w.middleware.GetResourceFromContext(c, backend.RIEngagement)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, &gin_backend_utility.WebappError{
				Err:     err,
				Context: "createResourceCommentsList - Obtain engagement in context",
			})
			return
		}

		rsc, err := w.middleware.GetResourceFromContext(c, resource)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, &gin_backend_utility.WebappError{
				Err:     err,
				Context: "createResourceCommentsList - Obtain resource in context",
			})
			return
		}

		torg := org.(*orgs.Organization)
		teng := engagement.(*engagements.Engagement)

		var threadId int64

		switch resource {
		case backend.RIRisk:
			trsc := rsc.(*risks.Risk)
			threadId, err = w.backend.itf.Comments.GetThreadIdForRisk(trsc.Id, teng.Id, torg.Id)
		default:
			err = errors.New("Unsupported resource for comments.")
		}

		if err != nil {
			c.AbortWithError(http.StatusBadRequest, &gin_backend_utility.WebappError{
				Err:     err,
				Context: "createResourceCommentsList - Obtain thread id",
			})
			return
		}

		comments, err := w.backend.itf.Comments.GetCommentsForThread(threadId)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, &gin_backend_utility.WebappError{
				Err:     err,
				Context: "createResourceCommentsList - Obtain comments",
			})
			return
		}

		c.JSON(http.StatusOK, comments)
	}
}
