package constanta

const (
	MenuGroupUser     = "user"
	MenuGroupCategory = "category"
	MenuGroupRole     = "role"
	MenuGroupPermissions    = "permissions"
)

const (
	MenuUserActionWrite = MenuGroupUser + ":" + AuthActionWrite
	MenuUserActionRead  = MenuGroupUser + ":" + AuthActionRead

	MenuCategoryActionWrite = MenuGroupCategory + ":" + AuthActionWrite
	MenuCategoryActionRead  = MenuGroupCategory + ":" + AuthActionRead

	MenuRoleActionWrite = MenuGroupRole + ":" + AuthActionWrite
	MenuRoleActionRead  = MenuGroupRole + ":" + AuthActionRead

	MenuPermissionsActionWrite = MenuGroupPermissions + ":" + AuthActionWrite
	MenuPermissionsActionRead  = MenuGroupPermissions + ":" + AuthActionRead
)
