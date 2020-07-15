package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"gitlab.com/grchive/grchive-v2/shared/backend/orgs"
	"net/http"
)

func (w *WebappApplication) apiv1CreateNewOrg(c *gin.Context) {
	org := orgs.Organization{}
	err := c.BindJSON(&org)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	currentUser := w.sessionStore.GetLoginSession(c).GetSessionUser()
	// This endpoint shouldn't allow creating a sub-org as this won't have
	// the proper ACL checks in place.
	org.ParentOrgId = nil
	org.OwnerUserId = currentUser.Id

	err = w.backend.itf.WrapDatabaseTx(func(tx *sqlx.Tx) error {
		return w.backend.itf.Orgs.CreateOrg(tx, &org)
	})

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, org)
}
