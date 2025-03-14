package request

// req user
type ReqUser struct {
	ID    string `json:"id"`
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
}

var ReqUserErrorMessage = map[string]string{
	"name":  "name required",
	"email": "email not valid",
}
