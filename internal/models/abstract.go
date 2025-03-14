package models

type GetListStruct struct {
	Filters map[string]interface{}
	Page    int
	Limit   int
}
