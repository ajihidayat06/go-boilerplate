package errorutils

import (
	"context"
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
	ErrMessageDataUpdated           = "data telah diperbarui, silakan muat ulang data terbaru"
	ErrMessaageDataAlreadyExists    = "data sudah ada"
	ErrMessaageDataRequired         = "data tidak boleh kosong"
	ErrMessageUserNotLogin          = "silahkan login terlebih dahulu"
	ErrMessageInternalServerError   = "terjadi kesalahan pada server, silahkan hubungi admin"
)

var (
	ErrInvalidCredentials  = errors.New("invalid credentials")
	ErrDataNotFound        = errors.New("data tidak ditemukan")
	ErrInternalServerError = errors.New("internal server error")
	ErrPasswordNotValid    = errors.New("password must contain at least 8 characters, including uppercase, lowercase, numbers, and special characters")
	ErrDataDataUpdated     = errors.New(ErrMessageDataUpdated)
)

type CustomError struct {
	Message    string
	FieldError string
	Err        error
}

func (e *CustomError) Error() string {
	if e.FieldError != "" {
		return fmt.Sprintf("%s [%s]", e.Message, e.FieldError)
	}
	return e.Message
}

func HandleCustomError(ctx context.Context, baseErr error, msg string, fieldError ...string) error {
	var fe string
	if len(fieldError) > 0 {
		fe = fieldError[0]
	}

	// Kalau baseErr nil, tetap buat error berdasarkan msg dan fieldError
	if baseErr == nil {
		baseErr = fmt.Errorf("%s", msg)
	}

	wrappedErr := &CustomError{
		Message:    msg,
		FieldError: fe,
		Err:        baseErr,
	}

	logger.LogWithCaller(ctx, msg, baseErr, 2)

	return wrappedErr
}

func HandleRepoError(ctx context.Context, err error) error {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		logger.LogWithCaller(ctx, ErrMessageDataNotFound, err, 2)
		return ErrDataNotFound
	}

	logger.LogWithCaller(ctx, ErrInternalServerError.Error(), err, 2)
	return ErrInternalServerError
}

func HandleUsecaseError(c *fiber.Ctx, err error, msg string) error {
	ctx := utils.GetContext(c)

	if errors.Is(err, ErrDataNotFound) {
		logger.LogWithCaller(ctx, ErrMessageDataNotFound, err, 2)
		return utils.SetResponseNotFound(c, ErrMessageDataNotFound, err)
	}

	if errors.Is(err, ErrInternalServerError) {
		logger.LogWithCaller(ctx, ErrInternalServerError.Error(), err, 2)
		return utils.SetResponseInternalServerError(c, ErrMessageInternalServerError, err)
	}

	logger.LogWithCaller(ctx, msg, err, 2)
	return utils.SetResponseBadRequest(c, msg, err)
}
