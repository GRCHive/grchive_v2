export enum Permission {
	POrgProfileView              = "org.profile.view",
	POrgProfileCreate            = "org.profile.create",
	POrgProfileUpdate            = "org.profile.update",
	POrgRolesView                = "org.roles.view",
	POrgRolesUpdate              = "org.roles.update",
	POrgRolesDelete              = "org.roles.delete",
	POrgRolesCreate              = "org.roles.create",
	POrgRolesList                = "org.roles.list",
	POrgUsersView                = "org.users.view",
	POrgUsersList                = "org.users.list",
	POrgEngagementList           = "org.engagements.list",
	POrgEngagementCreate         = "org.engagements.create",
	POrgEngagementView           = "org.engagements.view",
	POrgEngagementDelete         = "org.engagements.delete",
	POrgEngagementUpdate         = "org.engagements.update",
	POrgEngagementClose          = "org.engagements.close",
	POrgEngagementReopen         = "org.engagements.reopen",
	PRisksView                   = "org.risks.view",
	PRisksUpdate                 = "org.risks.update",
	PRisksDelete                 = "org.risks.delete",
	PRisksCreate                 = "org.risks.create",
	PRisksList                   = "org.risks.list",
	PControlsView                = "org.controls.view",
	PControlsUpdate              = "org.controls.update",
	PControlsDelete              = "org.controls.delete",
	PControlsCreate              = "org.controls.create",
	PControlsList                = "org.controls.list",
	PCommentsUpdate              = "org.comments.update",
	PCommentsDelete              = "org.comments.delete",
	PCommentsCreate              = "org.comments.create",
	PCommentsList                = "org.comments.list",
}

export interface Role {
    Id           : number
    OrgId        : number
    Name         : string
    Description  : string
}
