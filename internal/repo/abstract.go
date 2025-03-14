package repo

import (
	"context"
	"go-boilerplate/internal/constanta"
	"gorm.io/gorm"
)

type AbstractRepo struct {
	db *gorm.DB
}

func (a *AbstractRepo) SetDb(tx *gorm.DB) {
	a.db = tx
}

// Method untuk check scope (own atau all) menggunakan Fiber's c.Locals
func (a *AbstractRepo) withCheckScope(c context.Context) func(db *gorm.DB) *gorm.DB {
	scope := c.Value(constanta.Scope)
	userID := c.Value(constanta.AuthUserID)

	return func(db *gorm.DB) *gorm.DB {
		if scope == constanta.ScopeOwn && userID != nil {
			return a.db.Where("created_by = ?", userID)
		}
		return a.db
	}
}

func (a *AbstractRepo) paginate(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	offset := (page - 1) * pageSize
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(offset).Limit(pageSize)
	}
}
