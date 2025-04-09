package request

type ReqLogin struct {
	UsernameOrEmail string `json:"username_or_email" validate:"required"`
	Password        string `json:"password" validate:"required,min=8,max=20"`
}

var ReqLoginErrorMessage = map[string]string{
	"username_or_email": "invalid username or email",
	"password":          "invalid password",
}

type ReqToken struct {
	TemporaryToken string `json:"temporary_token" validate:"required"`
}
