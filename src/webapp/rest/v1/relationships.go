package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"gitlab.com/grchive/grchive-v2/shared/backend"
	"gitlab.com/grchive/grchive-v2/shared/gin_middleware/gin_backend_utility"
	"net/http"
	"strconv"
)

func (w *WebappApplication) addRelationshipEndpoints(
	r *gin.RouterGroup,
	from backend.ResourceIdentifier,
	to ...backend.ResourceIdentifier,
) {
	relR := r.Group("/relationships")

	for _, toResource := range to {
		toCrud := backend.ResourceToCrudPermissions(toResource)
		relCrud := backend.ResourceRelationshipToCrudPermissions(from, toResource)

		rr := relR.Group(backend.ResourceToEndpointName(toResource))

		rr.GET("/",
			w.acl.ACLUserHasPermissions(relCrud.List, toCrud.List),
			w.createResourceRelationshipList(from, toResource),
		)

		rr.POST("/",
			w.acl.ACLUserHasPermissions(relCrud.Create),
			w.createResourceRelationshipCreate(from, toResource),
		)

		rr.DELETE("/:resId",
			w.acl.ACLUserHasPermissions(relCrud.Delete),
			w.createResourceRelationshipDelete(from, toResource),
		)
	}
}

func (w *WebappApplication) createResourceRelationshipList(from backend.ResourceIdentifier, to backend.ResourceIdentifier) gin.HandlerFunc {
	return func(c *gin.Context) {
		resource, err := w.middleware.GetResourceFromContext(c, from)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, &gin_backend_utility.WebappError{
				Err:     err,
				Context: "createResourceRelationshipList - Obtain 'from' resource in context",
				Code:    gin_backend_utility.GECBadRequest,
				Message: gin_backend_utility.GEMBadRequest,
			})
			return
		}

		resourceId := backend.GetResourceId(resource)

		var rel []backend.RelationshipWrapper
		rel, err = w.backend.itf.ListRelationships(from, resourceId, to)

		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, &gin_backend_utility.WebappError{
				Err:     err,
				Context: "createResourceRelationshipList - Get relationships",
			})
			return
		}

		c.JSON(http.StatusOK, rel)
	}
}

func (w *WebappApplication) createResourceRelationshipCreate(from backend.ResourceIdentifier, to backend.ResourceIdentifier) gin.HandlerFunc {
	return func(c *gin.Context) {
		fromResource, err := w.middleware.GetResourceFromContext(c, from)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, &gin_backend_utility.WebappError{
				Err:     err,
				Context: "createResourceRelationshipCreate - Obtain 'from' resource in context",
				Code:    gin_backend_utility.GECBadRequest,
				Message: gin_backend_utility.GEMBadRequest,
			})
			return
		}
		fromResourceId := backend.GetResourceId(fromResource)
		fromEngId := backend.GetResourceEngagementId(fromResource)

		toResourceGenericId := backend.GenericResourceIdentifier{}
		err = c.BindJSON(&toResourceGenericId)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, &gin_backend_utility.WebappError{
				Err:     err,
				Context: "createResourceRelationshipCreate - Read resource ID from JSON body.",
				Code:    gin_backend_utility.GECBadRequest,
				Message: gin_backend_utility.GEMBadRequest,
			})
			return
		}

		// Make sure this resource belongs to the same org/engagement.
		// We do this just by checking the engagment id on both resources are equal (sicne the fromResource has its
		// engagement id -> org id relationship checked in our middleware stack).
		toResource, err := w.backend.itf.GetResourceFromGenericId(toResourceGenericId)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, &gin_backend_utility.WebappError{
				Err:     err,
				Context: "createResourceRelationshipCreate - Get 'to' resource",
				Code:    gin_backend_utility.GECBadRequest,
				Message: gin_backend_utility.GEMBadRequest,
			})
			return
		}
		toResourceId := backend.GetResourceId(toResource)
		toEngId := backend.GetResourceEngagementId(toResource)

		if fromEngId != toEngId {
			c.AbortWithError(http.StatusBadRequest, &gin_backend_utility.WebappError{
				Err:     nil,
				Context: "createResourceRelationshipCreate - Can't create inter-engagement relationships.",
				Code:    gin_backend_utility.GECBadRequest,
				Message: gin_backend_utility.GEMBadRequest,
			})
			return
		}

		err = w.backend.itf.WrapDatabaseTx(w.middleware.GetAuditTrailId(c), func(tx *sqlx.Tx) error {
			return w.backend.itf.CreateRelationship(tx, from, fromResourceId, to, toResourceId)
		})

		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, &gin_backend_utility.WebappError{
				Err:     err,
				Context: "createResourceRelationshipCreate - Failed to create relationships",
			})
			return
		}

		c.JSON(http.StatusOK, backend.RelationshipWrapper{
			Explicit: true,
			Data:     toResource,
		})
	}
}

func (w *WebappApplication) createResourceRelationshipDelete(from backend.ResourceIdentifier, to backend.ResourceIdentifier) gin.HandlerFunc {
	return func(c *gin.Context) {
		fromResource, err := w.middleware.GetResourceFromContext(c, from)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, &gin_backend_utility.WebappError{
				Err:     err,
				Context: "createResourceRelationshipDelete - Obtain 'from' resource in context",
				Code:    gin_backend_utility.GECBadRequest,
				Message: gin_backend_utility.GEMBadRequest,
			})
			return
		}
		fromResourceId := backend.GetResourceId(fromResource)

		toResourceId, err := strconv.ParseInt(c.Param("resId"), 10, 64)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, &gin_backend_utility.WebappError{
				Err:     err,
				Context: "createResourceRelationshipDelete - Get to resource id",
				Code:    gin_backend_utility.GECBadRequest,
				Message: gin_backend_utility.GEMBadRequest,
			})
			return
		}

		err = w.backend.itf.WrapDatabaseTx(w.middleware.GetAuditTrailId(c), func(tx *sqlx.Tx) error {
			return w.backend.itf.DeleteRelationship(tx, from, fromResourceId, to, toResourceId)
		})

		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, &gin_backend_utility.WebappError{
				Err:     err,
				Context: "createResourceRelationshipDelete - Failed to delete relationship",
			})
			return
		}

		c.Status(http.StatusNoContent)
	}
}
