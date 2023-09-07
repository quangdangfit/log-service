package queue

import (
	"context"
	"encoding/json"

	"github.com/quangdangfit/log-service/internal/dto"
	"github.com/quangdangfit/log-service/internal/service"
	"github.com/quangdangfit/log-service/pkg/logger"
)

type LogHandler struct {
	service service.ILogService
}

func NewLogHandler(service service.ILogService) *LogHandler {
	return &LogHandler{
		service: service,
	}
}

func (h *LogHandler) AddLog(ctx context.Context, data []byte) error {
	var req dto.AddLogReq
	err := json.Unmarshal(data, &req)
	if err != nil {
		logger.Error("Failed to parse data ", err)
		return err
	}

	_, err = h.service.AddLog(ctx, &req)
	if err != nil {
		logger.Error("Failed to add log ", err)
		return err
	}

	return nil
}
