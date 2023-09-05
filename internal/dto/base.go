package dto

type PaginationReq struct {
	Page  int `json:"page,omitempty" validate:"gte=0"`
	Limit int `json:"limit,omitempty" validate:"gte=0,lte=1000"`
}

type PaginationRes struct {
	CurrentPage int   `json:"current_page,omitempty"`
	TotalPage   int   `json:"total_page,omitempty"`
	Skip        int64 `json:"skip,omitempty"`
	Limit       int   `json:"limit,omitempty"`
	Total       int64 `json:"total,omitempty"`
}

type SortReq struct {
	Field string `json:"field,omitempty"`
	Desc  bool   `json:"desc,omitempty"`
}
