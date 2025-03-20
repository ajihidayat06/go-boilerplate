package request

// req user
type ReqUser struct {
	ID    int64  `json:"id"`
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
}

var ReqUserErrorMessage = map[string]string{
	"name":  "name required",
	"email": "email not valid",
}

type ReqUserUpdate struct {
	ID    int64  `json:"id" validate:"required"`
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
	AbstractRequest
}

var ReqUserUpdateErrorMessage = map[string]string{
	"id":    "id required",
	"name":  "name required",
	"email": "email not valid",
}

func (r *ReqUserUpdate) ValidateRequestUpdate() error {
	if err := r.ValidateUpdatedAt(); err != nil {
		return err
	}
	return nil
}
