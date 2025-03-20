package request

type ReqCategory struct {
    Name string `json:"name" validate:"required"`
}

var ReqCategoryErrorMessage = map[string]string{
    "name": "name required",
}

type ReqCategoryUpdate struct {
    ID        int64  `json:"id" validate:"required"`
    Name      string `json:"name" validate:"required"`
    AbstractRequest
}

var ReqCategoryUpdateErrorMessage = map[string]string{
    "id":         "id required",
    "name":       "name required",
    "updated_at": "updated_at required",
}