package main

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/grchive/grchive-v2/shared/backend"
	"gitlab.com/grchive/grchive-v2/shared/backend/roles"
	"net/http"
)

func (w *WebappApplication) apiv1EnsureAuth(c *gin.Context) {
	// For now just ensure we have a valid session.
	// TODO: #9jjxxu
	// Support API keys for non browser users.
	sess := w.sessionStore.GetLoginSession(c)
	if !sess.IsLoggedIn() {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	err := w.middleware.StoreCurrentUserIntoContext(c, sess.GetSessionUser())
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Next()
}

func (w *WebappApplication) registerApiv1(r *gin.Engine) {
	apiR := r.Group("/api/v1", w.apiv1EnsureAuth)
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

								w.addCommentEndpoints(
									backend.RIGLAccount,
									singleAccountR,
									w.acl.ACLUserHasPermissions(roles.PGLView),
								)
							}
						}
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
