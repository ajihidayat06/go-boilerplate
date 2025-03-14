package usecase

import (
	"gorm.io/gorm"
)

func processWithTx(db *gorm.DB, fn func(tx *gorm.DB) error) error {
	tx := db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	// Jika terjadi panic, rollback transaksi dan lanjutkan panic.
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		}
	}()

	err := fn(tx)
	if err != nil {
		tx.Rollback()
		return err
	}

	if commitErr := tx.Commit().Error; commitErr != nil {
		return commitErr
	}

	return nil
}
