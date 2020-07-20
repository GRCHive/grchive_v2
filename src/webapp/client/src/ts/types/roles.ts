export enum Permission {
	POrgProfileView              = "org.profile.view",
	POrgProfileCreate            = "org.profile.create",
	POrgProfileUpdate            = "org.profile.update",
	POrgRolesView                = "org.roles.view",
	POrgRolesUpdate              = "org.roles.update",
	POrgRolesDelete              = "org.roles.delete",
	POrgRolesCreate              = "org.roles.create",
	POrgRolesList                = "org.roles.list",
	POrgEngagementList           = "org.engagements.list",
	POrgEngagementCreate         = "org.engagements.create",
	POrgEngagementView           = "org.engagements.view",
	POrgEngagementDelete         = "org.engagements.delete",
	POrgEngagementUpdate         = "org.engagements.update",
	POrgEngagementClose          = "org.engagements.close",
	POrgEngagementReopen         = "org.engagements.reopen",
}

export interface Role {
    Id           : number
    OrgId        : number
    Name         : string
    Description  : string
}
