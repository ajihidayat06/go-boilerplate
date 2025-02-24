package constanta

const (
	MenuGroupUser = "user"
)

const (
	ActionCreate = "create"
	ActionRead   = "read"
	ActionUpdate = "update"
	ActionDelete = "delete"
)

const (
	MenuUserActionCreate = MenuGroupUser + ":" + ActionCreate
	MenuUserActionRead   = MenuGroupUser + ":" + ActionRead
)
