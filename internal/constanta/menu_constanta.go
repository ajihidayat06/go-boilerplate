package constanta

const (
	MenuGroupUser = "user"
	MenuGroupCategory = "category"
)

const (
	MenuUserActionWrite = MenuGroupUser + ":" + AuthActionWrite
	MenuUserActionRead  = MenuGroupUser + ":" + AuthActionRead

	MenuCategoryActionWrite = MenuGroupCategory + ":" + AuthActionWrite
	MenuCategoryActionRead  = MenuGroupCategory + ":" + AuthActionRead
)
