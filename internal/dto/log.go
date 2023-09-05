package dto

import (
	"time"
)

type Log struct {
	Id          string    `json:"id"`
	ServiceName string    `json:"service_name"`
	Level       string    `json:"level"`
	Content     []byte    `json:"content"`
	TraceId     string    `json:"trace_id"`
	CreatedAt   time.Time `json:"created_at"`
}

type AddLogReq struct {
	ServiceName string `json:"service_name" validate:"required"`
	Level       string `json:"level" validate:"required"`
	Content     []byte `json:"content" validate:"required"`
	TraceId     string `json:"trace_id" validate:"required"`
}

type GetLogsReq struct {
	ServiceName string         `json:"service_name,omitempty"`
	Level       string         `json:"level,omitempty"`
	TraceId     string         `json:"trace_id,omitempty"`
	CreatedFrom int64          `json:"created_from,omitempty"`
	CreatedTo   int64          `json:"created_to,omitempty"`
	Pagination  *PaginationReq `json:"pagination,omitempty"`
	Sort        *SortReq       `json:"sort,omitempty"`
}
