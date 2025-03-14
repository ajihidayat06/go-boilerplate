package request

type ReqLogin struct {
	UsenameOrEmail string `json:"usename_or_email" validate:"required"`
	Password       string `json:"password" validate:"required,min=8,max=20"`
}

var ReqLoginErrorMessage = map[string]string{
	"usename_or_email": "invalid username or email",
	"password":         "invalid password",
}
