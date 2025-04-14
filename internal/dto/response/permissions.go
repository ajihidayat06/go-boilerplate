package response

import (
	"go-boilerplate/internal/models"
	"time"
)

type PermissionsResponse struct {
	ID        int64     `json:"id"`
	Code      string    `json:"code"`
	Name      string    `json:"name"`
	GroupMenu string    `json:"group_menu"`
	Action    string    `json:"action"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedBy int64     `json:"created_by"`
	UpdatedBy int64     `json:"updated_by"`
}

func SetPermissionsRespons(permissions models.Permissions) PermissionsResponse {
	return PermissionsResponse{
		ID:        permissions.ID,
		Code:      permissions.Code,
		Name:      permissions.Name,
		GroupMenu: permissions.GroupMenu,
		Action:    permissions.Action,
		CreatedAt: permissions.CreatedAt,
		UpdatedAt: permissions.UpdatedAt,
		CreatedBy: permissions.CreatedBy,
		UpdatedBy: permissions.UpdatedBy,
	}
}

type RolePermissionsResponse struct {
	ID            int64               `json:"id"`
	RoleID        int64               `json:"role_id"`
	PermissionsID int64               `json:"permissions_id"`
	AccessScope   string              `json:"access_scope"`
	CreatedAt     time.Time           `json:"created_at"`
	UpdatedAt     time.Time           `json:"updated_at"`
	CreatedBy     int64               `json:"created_by"`
	UpdatedBy     int64               `json:"updated_by"`
	Permissions   PermissionsResponse `json:"permissions" gorm:"foreignKey:PermissionsID"`
}
