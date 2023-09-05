package grpc

import (
	"context"

	"github.com/quangdangfit/log-service/internal/dto"
	"github.com/quangdangfit/log-service/internal/service"
	"github.com/quangdangfit/log-service/pkg/logger"
	pb "github.com/quangdangfit/log-service/proto/gen/go/log"
)

var (
	LogLevelMapToString = map[pb.LogLevel]string{
		pb.LogLevel_LOG_LEVEL_DEBUG:   "DEBUG",
		pb.LogLevel_LOG_LEVEL_INFO:    "INFO",
		pb.LogLevel_LOG_LEVEL_WARNING: "WARNING",
		pb.LogLevel_LOG_LEVEL_ERROR:   "ERROR",
		pb.LogLevel_LOG_LEVEL_FATAL:   "FATAL",
	}
)

type LogHandler struct {
	pb.UnimplementedLogServiceServer

	service service.ILogService
}

func NewLogHandler(service service.ILogService) *LogHandler {
	return &LogHandler{
		service: service,
	}
}

func (h *LogHandler) AddLog(ctx context.Context, req *pb.AddLogReq) (*pb.AddLogRes, error) {
	log, err := h.service.AddLog(ctx, &dto.AddLogReq{
		ServiceName: req.ServiceName,
		Level:       LogLevelMapToString[req.Level],
		Content:     req.Content,
		TraceId:     req.TraceId,
	})
	if err != nil {
		logger.Error("Failed to register ", err)
		return nil, err
	}

	return &pb.AddLogRes{
		Log: log.ToProto(),
	}, nil
}

func (h *LogHandler) GetLogs(ctx context.Context, req *pb.GetLogsReq) (*pb.GetLogsRes, error) {
	r := &dto.GetLogsReq{
		ServiceName: req.ServiceName,
		Level:       LogLevelMapToString[req.Level],
		TraceId:     req.TraceId,
		CreatedFrom: req.CreatedFrom,
		CreatedTo:   req.CreatedTo,
	}
	if req.Pagination != nil {
		r.Pagination = &dto.PaginationReq{
			Page:  int(req.Pagination.Page),
			Limit: int(req.Pagination.Limit),
		}
	}
	if req.Sort != nil {
		r.Sort = &dto.SortReq{
			Field: req.Sort.Field,
			Desc:  req.Sort.Desc,
		}
	}

	logs, pagination, err := h.service.GetLogs(ctx, r)
	if err != nil {
		logger.Error("Failed to register ", err)
		return nil, err
	}

	return &pb.GetLogsRes{
		Logs: logs.ToProto(),
		Pagination: &pb.PaginationRes{
			CurrentPage: int64(pagination.CurrentPage),
			TotalPage:   int64(pagination.TotalPage),
			Skip:        pagination.Skip,
			Limit:       int64(pagination.Limit),
			Total:       pagination.Total,
		},
	}, nil
}
