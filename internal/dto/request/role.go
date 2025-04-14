package request

type ReqRoles struct {
	Code            string              `json:"code"`
	Name            string              `json:"name"`
	RolePermissions []ReqRolePermission `json:"role_permissions" validate:"required,dive"`
}

var ReqRolesErrorMessage = map[string]string{
	"Code":         "code required",
	"Name":         "name required",
	"PermissionID": "Permission ID required",
}

type ReqRoleUpdate struct {
	ID   int64  `json:"id" validate:"required"`
	Code string `json:"code" validate:"required"`
	Name string `json:"name" validate:"required"`
	AbstractRequest
}

var ReqRoleUpdateErrorMessage = map[string]string{
	"id":         "id required",
	"code":       "code required",
	"name":       "name required",
	"updated_at": "updated_at required",
}
