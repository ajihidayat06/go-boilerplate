package errorutils

import (
	"errors"
	"fmt"
	"go-boilerplate/internal/utils"
	"go-boilerplate/pkg/logger"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

const (
	ErrMessageForbidden             = "Anda tidak memiliki hak akses untuk mengakses menu ini."
	ErrMessageInvalidToken          = "Token tidak sesuai"
	ErrMessageExpiredToken          = "Token kadaluwarsa"
	ErrMessageInvalidOrExpiredToken = "Token tidak sesuai atau kadaluwarsa"
	ErrMessageDataNotFound          = "Data tidak ditemukan"
	ErrMessageInvalidRequestData    = "Data tidak valid"
)

var (
	ErrInvalidCredentials  = errors.New("invalid credentials")
	ErrDataNotFound        = errors.New("Data tidak ditemukan")
	ErrInternalServerError = errors.New("internal server error")
	ErrPasswordNotValid    = errors.New("password must contain at least 8 characters, including uppercase, lowercase, numbers, and special characters")
)

func HandleRepoError(err error) error {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return ErrDataNotFound
	}
	return errors.New(fmt.Sprintf(`%s : %s`, ErrInternalServerError.Error(), err.Error()))
}

func HandleUsecaseError(c *fiber.Ctx, err error, msg string) error {
	if errors.Is(err, ErrDataNotFound) {
		logger.Error(ErrMessageDataNotFound, err)
		return utils.SetResponseNotFound(c, ErrMessageDataNotFound, err)
	}

	if errors.Is(err, ErrInternalServerError) {
		logger.Error(ErrInternalServerError.Error(), err)
		return utils.SetResponseInternalServerError(c, ErrInternalServerError.Error(), err)
	}

	logger.Error(msg, err)
	return utils.SetResponseBadRequest(c, msg, err)
}
