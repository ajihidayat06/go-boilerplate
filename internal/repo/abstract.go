package repo

import "gorm.io/gorm"

type AbstractRepo struct {
	db *gorm.DB
}

func (a *AbstractRepo) SetDb(tx *gorm.DB) {
	a.db = tx
}
