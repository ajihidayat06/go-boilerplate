package request

// req user
type ReqUser struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (ReqUser) GenerateReqUserErrorMessage() map[string]string {
	reqUserErrorMessage := make(map[string]string)
	reqUserErrorMessage["name"] = "name required"
	reqUserErrorMessage["email"] = "email not valid"
	return reqUserErrorMessage
}
