package models

import "time"

type User struct {
	ID          int64     `json:"id" gorm:"primaryKey"`
	Username    string    `json:"username"`
	Password    string    `json:"password"`
	Email       string    `json:"email"`
	RoleID      int64     `json:"role_id"`
	BranchID    int64     `json:"branch_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Permissions []Permissions
}

func (User) Tablename() string {
	return "users"
}
