package constanta

const (
	MenuGroupUser     = "user"
	MenuGroupCategory = "category"
	MenuGroupRole     = "role"
)

const (
	MenuUserActionWrite = MenuGroupUser + ":" + AuthActionWrite
	MenuUserActionRead  = MenuGroupUser + ":" + AuthActionRead

	MenuCategoryActionWrite = MenuGroupCategory + ":" + AuthActionWrite
	MenuCategoryActionRead  = MenuGroupCategory + ":" + AuthActionRead

	MenuRoleActionWrite = MenuGroupRole + ":" + AuthActionWrite
	MenuRoleActionRead  = MenuGroupRole + ":" + AuthActionRead
)
