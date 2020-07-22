package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"gitlab.com/grchive/grchive-v2/shared/backend"
	"gitlab.com/grchive/grchive-v2/shared/backend/comments"
	"gitlab.com/grchive/grchive-v2/shared/backend/engagements"
	"gitlab.com/grchive/grchive-v2/shared/backend/orgs"
	"gitlab.com/grchive/grchive-v2/shared/backend/risks"
	"gitlab.com/grchive/grchive-v2/shared/backend/roles"
	"gitlab.com/grchive/grchive-v2/shared/gin_middleware/gin_backend_utility"
	"net/http"
	"time"
)

func (w *WebappApplication) addCommentEndpoints(resource backend.ResourceIdentifier, r *gin.RouterGroup, handlers ...gin.HandlerFunc) {
	commentR := r.Group("/comments", handlers...)
	{
		commentR.GET("/",
			w.acl.ACLUserHasPermissions(roles.PCommentsList),
			w.createResourceCommentsList(resource),
		)

		commentR.POST("/",
			w.acl.ACLUserHasPermissions(roles.PCommentsCreate),
			w.createResourceCommentsCreate(resource),
		)

	}
}

func (w *WebappApplication) getCommentThreadIdFromContext(resource backend.ResourceIdentifier, c *gin.Context) (int64, error) {
	org, err := w.middleware.GetResourceFromContext(c, backend.RIOrganization)
	if err != nil {
		return -1, err
	}

	engagement, err := w.middleware.GetResourceFromContext(c, backend.RIEngagement)
	if err != nil {
		return -1, err
	}

	rsc, err := w.middleware.GetResourceFromContext(c, resource)
	if err != nil {
		return -1, err
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

	return threadId, err
}

func (w *WebappApplication) createResourceCommentsList(resource backend.ResourceIdentifier) gin.HandlerFunc {
	return func(c *gin.Context) {
		threadId, err := w.getCommentThreadIdFromContext(resource, c)
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

func (w *WebappApplication) createResourceCommentsCreate(resource backend.ResourceIdentifier) gin.HandlerFunc {
	return func(c *gin.Context) {
		threadId, err := w.getCommentThreadIdFromContext(resource, c)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, &gin_backend_utility.WebappError{
				Err:     err,
				Context: "createResourceCommentsCreate - Obtain thread id",
			})
			return
		}

		currentUser := w.sessionStore.GetLoginSession(c).GetSessionUser()
		comment := comments.Comment{
			ThreadId: threadId,
			UserId:   currentUser.Id,
			PostTime: time.Now().UTC(),
			Content:  "",
		}

		err = c.BindJSON(&comment.Content)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, &gin_backend_utility.WebappError{
				Err:     err,
				Context: "createResourceCommentsCreate - Read comment from JSON body.",
			})
			return
		}

		err = w.backend.itf.WrapDatabaseTx(w.middleware.GetAuditTrailId(c), func(tx *sqlx.Tx) error {
			return w.backend.itf.Comments.CreateComment(tx, &comment)
		})

		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, &gin_backend_utility.WebappError{
				Err:     err,
				Context: "createResourceCommentsCreate - Create comment",
			})
			return
		}

		c.JSON(http.StatusOK, comment)
	}
}
