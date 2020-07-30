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
	PGLView                      = "org.gl.view",
	PGLUpdate                    = "org.gl.update",
	PGLDelete                    = "org.gl.delete",
	PGLCreate                    = "org.gl.create",
	PGLList                      = "org.gl.list",
	PVendorsView                 = "org.vendors.view",
	PVendorsUpdate               = "org.vendors.update",
	PVendorsDelete               = "org.vendors.delete",
	PVendorsCreate               = "org.vendors.create",
	PVendorsList                 = "org.vendors.list",
	PVendorProductsView          = "org.vendors.products.view",
	PVendorProductsUpdate        = "org.vendors.products.update",
	PVendorProductsDelete        = "org.vendors.products.delete",
	PVendorProductsCreate        = "org.vendors.products.create",
	PVendorProductsList          = "org.vendors.products.list",
	PServersView                 = "org.inventory.servers.view",
	PServersUpdate               = "org.inventory.servers.update",
	PServersDelete               = "org.inventory.servers.delete",
	PServersCreate               = "org.inventory.servers.create",
	PServersList                 = "org.inventory.servers.list",
	PDesktopsView                = "org.inventory.desktops.view",
	PDesktopsUpdate              = "org.inventory.desktops.update",
	PDesktopsDelete              = "org.inventory.desktops.delete",
	PDesktopsCreate              = "org.inventory.desktops.create",
	PDesktopsList                = "org.inventory.desktops.list",
	PLaptopsView                 = "org.inventory.laptops.view",
	PLaptopsUpdate               = "org.inventory.laptops.update",
	PLaptopsDelete               = "org.inventory.laptops.delete",
	PLaptopsCreate               = "org.inventory.laptops.create",
	PLaptopsList                 = "org.inventory.laptops.list",
	PMobileView                  = "org.inventory.mobile.view",
	PMobileUpdate                = "org.inventory.mobile.update",
	PMobileDelete                = "org.inventory.mobile.delete",
	PMobileCreate                = "org.inventory.mobile.create",
	PMobileList                  = "org.inventory.mobile.list",
	PEmbeddedView                = "org.inventory.embedded.view",
	PEmbeddedUpdate              = "org.inventory.embedded.update",
	PEmbeddedDelete              = "org.inventory.embedded.delete",
	PEmbeddedCreate              = "org.inventory.embedded.create",
	PEmbeddedList                = "org.inventory.embedded.list",
	PDatabasesView               = "org.databases.view",
	PDatabasesUpdate             = "org.databases.update",
	PDatabasesDelete             = "org.databases.delete",
	PDatabasesCreate             = "org.databases.create",
	PDatabasesList               = "org.databases.list",
	PSystemsView                 = "org.systems.view",
	PSystemsUpdate               = "org.systems.update",
	PSystemsDelete               = "org.systems.delete",
	PSystemsCreate               = "org.systems.create",
	PSystemsList                 = "org.systems.list",
	PRelRisksControlsDelete      = "org.rel.risks.controls.delete",
	PRelRisksControlsCreate      = "org.rel.risks.controls.create",
	PRelRisksControlsList        = "org.rel.risks.controls.list",
	PVMView                      = "org.vm.view",
	PVMUpdate                    = "org.vm.update",
	PVMDelete                    = "org.vm.delete",
	PVMCreate                    = "org.vm.create",
	PVMList                      = "org.vm.list",
    PNull = "null",
}

export interface Role {
    Id           : number
    OrgId        : number
    Name         : string
    Description  : string
}

export interface PermissionExpression {
    toString(): string
}

export class SinglePermission implements PermissionExpression {
    permission : Permission
    constructor(permission : Permission) {
        this.permission = permission
    }

    toString() : string {
        return `'${this.permission}'`
    }
}

export class PermissionExpressionBase {
    permissions : PermissionExpression[]

    constructor(permissions: PermissionExpression[]) {
        this.permissions = permissions
    }
}

export class PermissionAnd extends PermissionExpressionBase implements PermissionExpression {
    constructor(permissions: PermissionExpression[]) {
        super(permissions)
    }

    toString() : string {
        return '[' + this.permissions.map((ele: PermissionExpression) => ele.toString()).join(', and') + ']'
    }
}

export class PermissionOr extends PermissionExpressionBase implements PermissionExpression {
    constructor(permissions: PermissionExpression[]) {
        super(permissions)
    }

    toString() : string {
        return '[' + this.permissions.map((ele: PermissionExpression) => ele.toString()).join(', or') + ']'
    }
}
