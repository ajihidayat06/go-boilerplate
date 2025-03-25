package request

type ReqRoles struct {
	Code string `json:"code" validate:"required"`
	Name string `json:"name" validate:"required"`
}

var ReqRolesErrorMessage = map[string]string{
	"code": "code required",
	"name": "name required",
}

type ReqRoleUpdate struct {
	ID           int64  `json:"id" validate:"required"`
	Code         string `json:"code" validate:"required"`
	Name         string `json:"name" validate:"required"`
	AbstractRequest
}

var ReqRoleUpdateErrorMessage = map[string]string{
	"id":         "id required",
	"code":       "code required",
	"name":       "name required",
	"updated_at": "updated_at required",
}
