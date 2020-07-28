package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"gitlab.com/grchive/grchive-v2/shared/backend"
	"gitlab.com/grchive/grchive-v2/shared/backend/engagements"
	"gitlab.com/grchive/grchive-v2/shared/backend/systems"
	"gitlab.com/grchive/grchive-v2/shared/gin_middleware/gin_backend_utility"
	"net/http"
)

func (w *WebappApplication) apiv1ListSystems(c *gin.Context) {
	engagement, err := w.middleware.GetResourceFromContext(c, backend.RIEngagement)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, &gin_backend_utility.WebappError{
			Err:     err,
			Context: "apiv1ListSystems - Obtain engagement in context",
			Code:    gin_backend_utility.GECBadRequest,
			Message: gin_backend_utility.GEMBadRequest,
		})
		return
	}

	teng := engagement.(*engagements.Engagement)
	dbs, err := w.backend.itf.Systems.ListSystemsForEngagement(teng.Id)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, &gin_backend_utility.WebappError{
			Err:     err,
			Context: "apiv1ListSystems - List systems",
			Code:    gin_backend_utility.GECBadRequest,
			Message: gin_backend_utility.GEMBadRequest,
		})
	}

	c.JSON(http.StatusOK, dbs)
}

func (w *WebappApplication) apiv1CreateSystem(c *gin.Context) {
	engagement, err := w.middleware.GetResourceFromContext(c, backend.RIEngagement)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, &gin_backend_utility.WebappError{
			Err:     err,
			Context: "apiv1CreateSystem - Obtain engagement in context",
			Code:    gin_backend_utility.GECBadRequest,
			Message: gin_backend_utility.GEMBadRequest,
		})
		return
	}

	teng := engagement.(*engagements.Engagement)

	db := systems.System{}
	err = c.BindJSON(&db)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, &gin_backend_utility.WebappError{
			Err:     err,
			Context: "apiv1CreateSystem - Read risk from JSON body.",
			Code:    gin_backend_utility.GECBadRequest,
			Message: gin_backend_utility.GEMBadRequest,
		})
		return
	}
	db.EngagementId = teng.Id

	err = w.backend.itf.WrapDatabaseTx(w.middleware.GetAuditTrailId(c), func(tx *sqlx.Tx) error {
		return w.backend.itf.Systems.CreateSystem(tx, &db)
	})

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, &gin_backend_utility.WebappError{
			Err:     err,
			Context: "apiv1CreateSystem - Create system",
		})
		return
	}

	c.JSON(http.StatusOK, db)
}

func (w *WebappApplication) apiv1GetSystem(c *gin.Context) {
	db, err := w.middleware.GetResourceFromContext(c, backend.RISystem)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, &gin_backend_utility.WebappError{
			Err:     err,
			Context: "apiv1GetSystem - Obtain system in context",
			Code:    gin_backend_utility.GECBadRequest,
			Message: gin_backend_utility.GEMBadRequest,
		})
		return
	}
	c.JSON(http.StatusOK, db)
}

func (w *WebappApplication) apiv1UpdateSystem(c *gin.Context) {
	currentSystem, err := w.middleware.GetResourceFromContext(c, backend.RISystem)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, &gin_backend_utility.WebappError{
			Err:     err,
			Context: "apiv1UpdateSystem - Obtain system in context",
			Code:    gin_backend_utility.GECBadRequest,
			Message: gin_backend_utility.GEMBadRequest,
		})
		return
	}

	system := systems.System{}
	err = c.BindJSON(&system)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, &gin_backend_utility.WebappError{
			Err:     err,
			Context: "apiv1UpdateSystem - Read system from JSON body.",
			Code:    gin_backend_utility.GECBadRequest,
			Message: gin_backend_utility.GEMBadRequest,
		})
		return
	}

	tCurrentSystem := currentSystem.(*systems.System)
	system.Id = tCurrentSystem.Id
	system.EngagementId = tCurrentSystem.EngagementId

	err = w.backend.itf.WrapDatabaseTx(w.middleware.GetAuditTrailId(c), func(tx *sqlx.Tx) error {
		return w.backend.itf.Systems.UpdateSystem(tx, &system)
	})

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, &gin_backend_utility.WebappError{
			Err:     err,
			Context: "apiv1UpdateSystem - Update system",
		})
		return
	}

	c.JSON(http.StatusOK, system)
}

func (w *WebappApplication) apiv1DeleteSystem(c *gin.Context) {
	db, err := w.middleware.GetResourceFromContext(c, backend.RISystem)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, &gin_backend_utility.WebappError{
			Err:     err,
			Context: "apiv1DeleteSystem - Obtain system in context",
			Code:    gin_backend_utility.GECBadRequest,
			Message: gin_backend_utility.GEMBadRequest,
		})
		return
	}
	tdb := db.(*systems.System)

	err = w.backend.itf.WrapDatabaseTx(w.middleware.GetAuditTrailId(c), func(tx *sqlx.Tx) error {
		return w.backend.itf.Systems.DeleteSystem(tx, tdb.Id)
	})

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, &gin_backend_utility.WebappError{
			Err:     err,
			Context: "apiv1DeleteSystem - Delete system",
		})
		return
	}

	c.Status(http.StatusNoContent)
}
