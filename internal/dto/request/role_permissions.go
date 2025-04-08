package request

type ReqRolePermission struct {
	ID           int64  `json:"id"`
	RoleID       int64  `json:"role_id" validate:"required"`
	PermissionID int64  `json:"permission_id" validate:"required"`
	Scope        string `json:"scope"`
	AbstractRequest
}

var ReqRolePermissionErrorMessage = map[string]string{
	"RoleID.required":       "Role ID is required",
	"PermissionID.required": "Permission ID is required",
}
