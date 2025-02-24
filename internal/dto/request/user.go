package request

// req user
type ReqUser struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var ReqUserErrorMessage = map[string]string{
	"name":  "name required",
	"email": "email not valid",
}

type ReqLogin struct {
	UsenameOrEmail string `json:"usename_or_email"`
	Password       string `json:"password"`
}

var ReqLoginErrorMessage = map[string]string{
	"usename_or_email": "invalid username or email",
	"password":         "invalid password",
}
