package errors

import "errors"

const (
	ErrMessageForbidden             = "Anda tidak memiliki hak akses untuk mengakses menu ini."
	ErrMessageInvalidToken          = "Token tidak sesuai"
	ErrMessageExpiredToken          = "Token kadaluwarsa"
	ErrMessageInvalidOrExpiredToken = "Token tidak sesuai atau kadaluwarsa"
)

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
)
