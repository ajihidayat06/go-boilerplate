package request

import (
	"errors"
	"go-boilerplate/pkg/logger"
)

type ReqLogin struct {
	UsernameOrEmail string `json:"username_or_email" validate:"required"`
	Password        string `json:"password" validate:"required,min=8,max=20"`
}

func (r *ReqLogin) ValidateRequest() error {
	// Validasi input
	if r.UsernameOrEmail == "" || r.Password == "" {
		logger.Error("Login attempt with empty username or password", nil)
		return errors.New("Username and password are required")
	}

	return nil
}

var ReqLoginErrorMessage = map[string]string{
	"username_or_email": "invalid username or email",
	"password":          "invalid password",
}

type ReqToken struct {
	TemporaryToken string `json:"temporary_token" validate:"required"`
}

func (r *ReqToken) ValidateRequest() error {
	// Validasi input
	if r.TemporaryToken == "" {
		err := errors.New("Token are required")
		logger.Error("temporary token nil", err)
		return err
	}

	return nil
}
