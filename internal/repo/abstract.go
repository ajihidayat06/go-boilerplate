package repo

import (
	"context"
	"go-boilerplate/internal/constanta"
	"go-boilerplate/internal/models"
	"gorm.io/gorm"
	"strings"
)

type AbstractRepo struct {
	db *gorm.DB
}

func (a *AbstractRepo) SetDb(tx *gorm.DB) {
	a.db = tx
}

// Method untuk check scope (own atau all)
func (a *AbstractRepo) withCheckScope(c context.Context) func(db *gorm.DB) *gorm.DB {
	scope := c.Value(constanta.Scope)
	userID := c.Value(constanta.AuthUserID)

	return func(db *gorm.DB) *gorm.DB {
		if scope == constanta.ScopeOwn && userID != nil {
			return db.Where("created_by = ?", userID)
		}
		return db
	}
}

// Pagination function
func (a *AbstractRepo) paginate(page, pageSize int) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

// Apply filters to the query
func (a *AbstractRepo) applyFilters(filters map[string][2]interface{}) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		for key, val := range filters {
			operator := val[0].(string)
			value := val[1]

			// Jika operatornya IN, kita harus memastikan value bisa digunakan dalam IN
			if operator == "IN" {
				switch v := value.(type) {
				case []string, []int, []int64, []float64:
					db = db.Where(key+" IN (?)", v)
				default:
					continue // Skip filter jika tipe tidak sesuai
				}
			} else {
				db = db.Where(key+" "+operator+" ?", value)
			}
		}
		return db
	}
}

func (a *AbstractRepo) orderBy(field string, sortBy string) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if strings.TrimSpace(strings.ToUpper(sortBy)) == "ASC" {
			return db.Order(field + " ASC")
		}
		return db.Order(field + " DESC")
	}
}

// Combine filters and pagination
func (a *AbstractRepo) applyFiltersAndPaginationAndOrder(params *models.GetListStruct) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		db = a.applyFilters(params.Filters)(db)
		db = a.paginate(params.Page, params.Limit)(db)
		db = a.orderBy(params.OrderBy, params.SortBy)(db)
		return db
	}
}

func (a *AbstractRepo) filterByRole(role string, userID int64) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if role == "admin" {
			return db // Admin bisa akses semua
		}
		return db.Where("created_by = ?", userID)
	}
}
