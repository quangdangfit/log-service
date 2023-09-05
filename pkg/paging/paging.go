package paging

import (
	"math"
)

const (
	DefaultPageSize = 20
)

// Pagination struct
type Pagination struct {
	CurrentPage int   `json:"current_page"`
	TotalPage   int   `json:"total_page"`
	Skip        int64 `json:"skip"`
	Limit       int   `json:"limit"`
	Total       int64 `json:"total"`
}

// New paging object
func New(page int, pageSize int, total int64) *Pagination {
	var pageInfo Pagination
	limit := DefaultPageSize
	if pageSize > 0 && pageSize <= limit {
		pageInfo.Limit = pageSize
	} else {
		pageInfo.Limit = limit
	}

	totalPage := int(math.Ceil(float64(total) / float64(pageInfo.Limit)))
	pageInfo.Total = total
	pageInfo.TotalPage = totalPage
	if page < 1 || totalPage == 0 {
		page = 1
	}

	pageInfo.CurrentPage = page
	pageInfo.Skip = int64((page - 1) * pageInfo.Limit)
	return &pageInfo
}
