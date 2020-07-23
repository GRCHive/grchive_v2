package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"gitlab.com/grchive/grchive-v2/shared/backend"
	"gitlab.com/grchive/grchive-v2/shared/backend/comments"
	"gitlab.com/grchive/grchive-v2/shared/backend/roles"
	"gitlab.com/grchive/grchive-v2/shared/gin_middleware/gin_backend_utility"
	"net/http"
	"strconv"
	"time"
)

func (w *WebappApplication) addCommentEndpoints(resource backend.ResourceIdentifier, r *gin.RouterGroup, handlers ...gin.HandlerFunc) {
	commentR := r.Group("/comments", append(handlers, w.middleware.LoadCommentThreadIdIntoContext(resource))...)
	{
		commentR.GET("/",
			w.acl.ACLUserHasPermissions(roles.PCommentsList),
			w.createResourceCommentsList(resource),
		)

		commentR.POST("/",
			w.acl.ACLUserHasPermissions(roles.PCommentsCreate),
			w.createResourceCommentsCreate(resource),
		)

		singleCommentR := commentR.Group("/:commentId")
		{
			singleCommentR.PUT("/",
				w.acl.ACLUserHasPermissions(roles.PCommentsUpdate),
				w.createResourceCommentsUpdate(resource),
			)

			singleCommentR.DELETE("/",
				w.acl.ACLUserHasPermissions(roles.PCommentsDelete),
				w.createResourceCommentsDelete(resource),
			)
		}
	}
}

func (w *WebappApplication) getCommentFromContext(resource backend.ResourceIdentifier, c *gin.Context) (*comments.Comment, error) {
	threadId, err := w.middleware.GetCommentThreadIdFromContext(c)
	if err != nil {
		return nil, err
	}

	commentIdStr := c.Param("commentId")
	commentId, err := strconv.ParseInt(commentIdStr, 10, 64)
	if err != nil {
		return nil, err
	}

	// We need to grab the current comment to verify that this comment actually lives
	// in the thread the user is accessing since our comment IDs are unique across all threads.
	return w.backend.itf.Comments.GetCommentById(commentId, threadId)
}

func (w *WebappApplication) createResourceCommentsList(resource backend.ResourceIdentifier) gin.HandlerFunc {
	return func(c *gin.Context) {
		threadId, err := w.middleware.GetCommentThreadIdFromContext(c)
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
		threadId, err := w.middleware.GetCommentThreadIdFromContext(c)
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

func (w *WebappApplication) createResourceCommentsUpdate(resource backend.ResourceIdentifier) gin.HandlerFunc {
	return func(c *gin.Context) {
		comment, err := w.getCommentFromContext(resource, c)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, &gin_backend_utility.WebappError{
				Err:     err,
				Context: "createResourceCommentsUpdate - Get current comment",
			})
			return
		}

		err = c.BindJSON(&comment.Content)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, &gin_backend_utility.WebappError{
				Err:     err,
				Context: "createResourceCommentsUpdate - Read comment from JSON body.",
			})
			return
		}

		err = w.backend.itf.WrapDatabaseTx(w.middleware.GetAuditTrailId(c), func(tx *sqlx.Tx) error {
			return w.backend.itf.Comments.UpdateComment(tx, comment)
		})

		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, &gin_backend_utility.WebappError{
				Err:     err,
				Context: "createResourceCommentsUpdate - Update comment",
			})
			return
		}

		c.JSON(http.StatusOK, comment)
	}
}

func (w *WebappApplication) createResourceCommentsDelete(resource backend.ResourceIdentifier) gin.HandlerFunc {
	return func(c *gin.Context) {
		comment, err := w.getCommentFromContext(resource, c)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, &gin_backend_utility.WebappError{
				Err:     err,
				Context: "createResourceCommentsDelete - Get current comment",
			})
			return
		}

		err = w.backend.itf.WrapDatabaseTx(w.middleware.GetAuditTrailId(c), func(tx *sqlx.Tx) error {
			return w.backend.itf.Comments.DeleteComment(tx, comment.Id)
		})

		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, &gin_backend_utility.WebappError{
				Err:     err,
				Context: "createResourceCommentsDelete - Delete comment",
			})
			return
		}

		c.Status(http.StatusNoContent)
	}
}
