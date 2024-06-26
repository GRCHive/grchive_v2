package main

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/grchive/grchive-v2/shared/backend"
	"gitlab.com/grchive/grchive-v2/shared/backend/roles"
	"gitlab.com/grchive/grchive-v2/shared/gin_middleware/gin_backend_utility"
	"net/http"
)

func (w *WebappApplication) apiv1EnsureAuth(c *gin.Context) {
	// For now just ensure we have a valid session.
	// TODO: #9jjxxu
	// Support API keys for non browser users.
	sess := w.sessionStore.GetLoginSession(c)
	err := sess.ValidateLogin(w.fusionauth)
	if err != nil || !sess.IsLoggedIn() {
		c.AbortWithError(http.StatusUnauthorized, &gin_backend_utility.WebappError{
			Err:     err,
			Context: "apiv1EnsureAuth: Check login.",
			Code:    gin_backend_utility.GECUnauthorized,
			Message: gin_backend_utility.GEMUnauthorizedLogin,
		})
		return
	}

	err = w.middleware.StoreCurrentUserIntoContext(c, sess.GetSessionUser())
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, &gin_backend_utility.WebappError{
			Err:     err,
			Context: "Store user into context.",
		})
		return
	}

	c.Next()
}

func (w *WebappApplication) registerApiv1(r *gin.Engine) {
	apiR := r.Group("/api/v1",
		gin_backend_utility.ApiMiddlewareErrorHandling,
		w.apiv1EnsureAuth,
	)
	{
		userR := apiR.Group("/users")
		{
			currentR := userR.Group("/current")
			currentR.GET("/", w.apiv1GetCurrentUser)
			currentR.PUT("/", w.apiv1UpdateCurrentUser)
			currentR.GET("/orgs", w.apiv1GetCurrentUserOrgs)
			userOrgR := currentR.Group("/orgs/:orgId",
				w.middleware.LoadResourceIntoContext(backend.RIOrganization, "orgId"),
			)
			{
				userOrgR.GET("/permissions", w.apiv1CheckCurrentUserPermissions)

				userOrgEngR := userOrgR.Group("/engagement/:engId",
					w.middleware.LoadResourceIntoContext(backend.RIEngagement, "engId"),
				)
				{
					userOrgEngR.GET("/permissions", w.apiv1CheckCurrentUserPermissions)
				}
			}

			currentR.POST("/verify", w.apiv1ResendEmailVerification)
		}

		orgsR := apiR.Group("/orgs", w.acl.ACLUserEmailVerified)
		{
			orgsR.POST("/", w.apiv1CreateNewOrg)

			singleOrgR := orgsR.Group("/:orgId",
				w.middleware.LoadResourceIntoContext(backend.RIOrganization, "orgId"),
				w.middleware.LoadCurrentUserRolesIntoContext,
				w.acl.ACLUserHasValidRole,
			)
			{
				singleOrgR.GET("/",
					w.acl.ACLUserHasPermissions(roles.POrgProfileView),
					w.apiv1GetOrg)

				singleOrgR.PUT("/",
					w.acl.ACLUserHasPermissions(roles.POrgProfileUpdate),
					w.apiv1UpdateOrg)

				singleOrgR.POST("/",
					w.acl.ACLUserHasPermissions(roles.POrgProfileCreate),
					w.apiv1CreateSuborg)

				singleOrgR.GET("/suborgs",
					w.acl.ACLUserHasPermissions(roles.POrgProfileView),
					w.apiv1GetSuborgs)

				singleOrgR.GET("/parents",
					w.acl.ACLUserHasPermissions(roles.POrgProfileView),
					w.apiv1GetParentOrgs)

				orgUsersR := singleOrgR.Group("/users")
				{
					orgUsersR.GET("/",
						w.acl.ACLUserHasPermissions(roles.POrgUsersList),
						w.apiv1ListOrgUsers)

					singleOrgUserR := orgUsersR.Group("/:userId",
						w.middleware.LoadResourceIntoContext(backend.RIUser, "userId"),
						w.middleware.CheckResourcePartOfOrg(backend.RIUser),
					)

					{
						singleOrgUserR.GET("/",
							w.acl.ACLUserHasPermissions(roles.POrgUsersView),
							w.apiv1GetOrgUser)
					}
				}

				engagementsR := singleOrgR.Group("/engagements")
				{
					engagementsR.GET("/",
						w.acl.ACLBranchOnPermissions(
							w.apiv1ListAllOrgEngagements,
							w.apiv1ListAllOrgEngagementsForCurrentUser,
							roles.POrgEngagementList,
						))

					engagementsR.POST("/",
						w.acl.ACLUserHasPermissions(roles.POrgEngagementCreate),
						w.apiv1CreateEngagement)

					singleEngR := engagementsR.Group("/:engId",
						w.middleware.LoadResourceIntoContext(backend.RIEngagement, "engId"),
						w.middleware.CheckResourcePartOfOrg(backend.RIEngagement),
					)

					singleEngR.GET("/",
						w.acl.ACLUserHasPermissions(roles.POrgEngagementView),
						w.apiv1GetEngagement)

					singleEngR.PUT("/",
						w.acl.ACLUserHasPermissions(roles.POrgEngagementUpdate),
						w.apiv1UpdateEngagement)

					statsR := singleEngR.Group("/stats", w.acl.ACLUserHasPermissions(roles.POrgEngagementView))
					{
						statsR.GET("/scoping", w.apiv1GetEngagementScopingStats)
					}

					risksR := singleEngR.Group("/risks")
					{
						risksR.GET("/",
							w.acl.ACLUserHasPermissions(roles.PRisksList),
							w.apiv1ListRisks)

						risksR.POST("/",
							w.acl.ACLUserHasPermissions(roles.PRisksCreate),
							w.apiv1CreateRisk)

						singleRiskR := risksR.Group("/:riskId",
							w.middleware.LoadResourceIntoContext(backend.RIRisk, "riskId"),
							w.middleware.CheckResourcePartOfEngagement(backend.RIRisk),
						)

						{
							singleRiskR.GET("/",
								w.acl.ACLUserHasPermissions(roles.PRisksView),
								w.apiv1GetRisk)

							singleRiskR.PUT("/",
								w.acl.ACLUserHasPermissions(roles.PRisksUpdate),
								w.apiv1UpdateRisk)

							singleRiskR.DELETE("/",
								w.acl.ACLUserHasPermissions(roles.PRisksDelete),
								w.apiv1DeleteRisk)

							w.addCommentEndpoints(
								backend.RIRisk,
								singleRiskR,
								w.acl.ACLUserHasPermissions(roles.PRisksView),
							)

							w.addRelationshipEndpoints(
								singleRiskR,
								backend.RIRisk,
								backend.RIControl,
							)
						}
					}

					controlsR := singleEngR.Group("/controls")
					{
						controlsR.GET("/",
							w.acl.ACLUserHasPermissions(roles.PControlsList),
							w.apiv1ListControls)

						controlsR.POST("/",
							w.acl.ACLUserHasPermissions(roles.PControlsCreate),
							w.apiv1CreateControl)

						singleControlR := controlsR.Group("/:controlId",
							w.middleware.LoadResourceIntoContext(backend.RIControl, "controlId"),
							w.middleware.CheckResourcePartOfEngagement(backend.RIControl),
						)

						{
							singleControlR.GET("/",
								w.acl.ACLUserHasPermissions(roles.PControlsView),
								w.apiv1GetControl)

							singleControlR.PUT("/",
								w.acl.ACLUserHasPermissions(roles.PControlsUpdate),
								w.apiv1UpdateControl)

							singleControlR.DELETE("/",
								w.acl.ACLUserHasPermissions(roles.PControlsDelete),
								w.apiv1DeleteControl)

							w.addCommentEndpoints(
								backend.RIControl,
								singleControlR,
								w.acl.ACLUserHasPermissions(roles.PControlsView),
							)

							w.addRelationshipEndpoints(
								singleControlR,
								backend.RIControl,
								backend.RIRisk,
							)
						}
					}

					glR := singleEngR.Group("/gl",
						w.middleware.LoadGeneralLedgerIntoContext,
					)
					{
						w.addCommentEndpoints(
							backend.RIGeneralLedger,
							glR,
							w.acl.ACLUserHasPermissions(roles.PGLList),
						)

						accountsR := glR.Group("/accs")
						{
							accountsR.GET("/",
								w.acl.ACLUserHasPermissions(roles.PGLList),
								w.apiv1ListGLAccounts)

							accountsR.POST("/",
								w.acl.ACLUserHasPermissions(roles.PGLCreate),
								w.apiv1CreateGLAccount)

							singleAccountR := accountsR.Group("/:accId",
								w.middleware.LoadResourceIntoContext(backend.RIGLAccount, "accId"),
								w.middleware.CheckResourcePartOfEngagement(backend.RIGLAccount),
							)

							{
								singleAccountR.GET("/",
									w.acl.ACLUserHasPermissions(roles.PGLView),
									w.apiv1GetGLAccount)

								singleAccountR.DELETE("/",
									w.acl.ACLUserHasPermissions(roles.PGLDelete),
									w.apiv1DeleteGLAccount)

								singleAccountR.PUT("/",
									w.acl.ACLUserHasPermissions(roles.PGLUpdate),
									w.apiv1UpdateGLAccount)

								singleAccountR.GET("/subaccounts",
									w.acl.ACLUserHasPermissions(roles.PGLView),
									w.apiv1GetGLSubaccounts)

								singleAccountR.GET("/parents",
									w.acl.ACLUserHasPermissions(roles.PGLView),
									w.apiv1GetGLParentAccounts)

								w.addCommentEndpoints(
									backend.RIGLAccount,
									singleAccountR,
									w.acl.ACLUserHasPermissions(roles.PGLView),
								)
							}
						}
					}

					vendorsR := singleEngR.Group("/vendors")
					{
						vendorsR.GET("/",
							w.acl.ACLUserHasPermissions(roles.PVendorsList),
							w.apiv1ListVendors)

						vendorsR.POST("/",
							w.acl.ACLUserHasPermissions(roles.PVendorsCreate),
							w.apiv1CreateVendor)

						singleVendorR := vendorsR.Group("/:vendorId",
							w.middleware.LoadResourceIntoContext(backend.RIVendor, "vendorId"),
							w.middleware.CheckResourcePartOfEngagement(backend.RIVendor),
						)

						{
							singleVendorR.GET("/",
								w.acl.ACLUserHasPermissions(roles.PVendorsView),
								w.apiv1GetVendor)

							singleVendorR.PUT("/",
								w.acl.ACLUserHasPermissions(roles.PVendorsUpdate),
								w.apiv1UpdateVendor)

							singleVendorR.DELETE("/",
								w.acl.ACLUserHasPermissions(roles.PVendorsDelete),
								w.apiv1DeleteVendor)

							w.addCommentEndpoints(
								backend.RIVendor,
								singleVendorR,
								w.acl.ACLUserHasPermissions(roles.PVendorsView),
							)

							productsR := singleVendorR.Group("/products")
							{
								productsR.GET("/",
									w.acl.ACLUserHasPermissions(roles.PVendorProductsList),
									w.apiv1ListVendorProducts)

								productsR.POST("/",
									w.acl.ACLUserHasPermissions(roles.PVendorProductsCreate),
									w.apiv1CreateVendorProduct)

								singleProductR := productsR.Group("/:productId",
									w.middleware.LoadResourceIntoContext(backend.RIVendorProduct, "productId"),
									w.middleware.CheckResourcePartOfVendor(backend.RIVendorProduct),
								)

								{
									singleProductR.GET("/",
										w.acl.ACLUserHasPermissions(roles.PVendorsView),
										w.apiv1GetVendorProduct)

									singleProductR.PUT("/",
										w.acl.ACLUserHasPermissions(roles.PVendorsUpdate),
										w.apiv1UpdateVendorProduct)

									singleProductR.DELETE("/",
										w.acl.ACLUserHasPermissions(roles.PVendorsDelete),
										w.apiv1DeleteVendorProduct)
								}
							}
						}
					}

					inventoryR := singleEngR.Group("/inventory")
					{
						w.apiv1CreateInventoryEndpoints(inventoryR, "servers", backend.RIInventoryServer)
						w.apiv1CreateInventoryEndpoints(inventoryR, "desktops", backend.RIInventoryDesktop)
						w.apiv1CreateInventoryEndpoints(inventoryR, "laptops", backend.RIInventoryLaptop)
						w.apiv1CreateInventoryEndpoints(inventoryR, "mobile", backend.RIInventoryMobile)
						w.apiv1CreateInventoryEndpoints(inventoryR, "embedded", backend.RIInventoryEmbedded)
					}

					databaseR := singleEngR.Group("/databases")
					{
						databaseR.GET("/",
							w.acl.ACLUserHasPermissions(roles.PDatabasesList),
							w.apiv1ListDatabases)

						databaseR.POST("/",
							w.acl.ACLUserHasPermissions(roles.PDatabasesCreate),
							w.apiv1CreateDatabase)

						singleDatabaseR := databaseR.Group("/:dbId",
							w.middleware.LoadResourceIntoContext(backend.RIDatabase, "dbId"),
							w.middleware.CheckResourcePartOfEngagement(backend.RIDatabase),
						)

						{
							singleDatabaseR.GET("/",
								w.acl.ACLUserHasPermissions(roles.PDatabasesView),
								w.apiv1GetDatabase)

							singleDatabaseR.PUT("/",
								w.acl.ACLUserHasPermissions(roles.PDatabasesUpdate),
								w.apiv1UpdateDatabase)

							singleDatabaseR.DELETE("/",
								w.acl.ACLUserHasPermissions(roles.PDatabasesDelete),
								w.apiv1DeleteDatabase)

							w.addCommentEndpoints(
								backend.RIDatabase,
								singleDatabaseR,
								w.acl.ACLUserHasPermissions(roles.PDatabasesView),
							)
						}
					}

					systemR := singleEngR.Group("/systems")
					{
						systemR.GET("/",
							w.acl.ACLUserHasPermissions(roles.PSystemsList),
							w.apiv1ListSystems)

						systemR.POST("/",
							w.acl.ACLUserHasPermissions(roles.PSystemsCreate),
							w.apiv1CreateSystem)

						singleSystemR := systemR.Group("/:sysId",
							w.middleware.LoadResourceIntoContext(backend.RISystem, "sysId"),
							w.middleware.CheckResourcePartOfEngagement(backend.RISystem),
						)

						{
							singleSystemR.GET("/",
								w.acl.ACLUserHasPermissions(roles.PSystemsView),
								w.apiv1GetSystem)

							singleSystemR.PUT("/",
								w.acl.ACLUserHasPermissions(roles.PSystemsUpdate),
								w.apiv1UpdateSystem)

							singleSystemR.DELETE("/",
								w.acl.ACLUserHasPermissions(roles.PSystemsDelete),
								w.apiv1DeleteSystem)

							w.addCommentEndpoints(
								backend.RISystem,
								singleSystemR,
								w.acl.ACLUserHasPermissions(roles.PSystemsView),
							)
						}
					}

					vmR := singleEngR.Group("/vm")
					{
						vmR.GET("/",
							w.acl.ACLUserHasPermissions(roles.PVMList),
							w.apiv1ListVM)
					}
				}

				rolesR := singleOrgR.Group("/roles")
				{
					rolesR.GET("/",
						w.acl.ACLBranchOnPermissions(
							w.apiv1ListAllOrgRoles,
							w.apiv1ListAllOrgRolesForCurrentUser,
							roles.POrgRolesList,
						))
				}
			}
		}
	}
}
