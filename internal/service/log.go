package service

import (
	"context"

	"github.com/quangdangfit/log-service/internal/dto"
	"github.com/quangdangfit/log-service/internal/repository"
	"github.com/quangdangfit/log-service/pkg/paging"
	"github.com/quangdangfit/log-service/pkg/validation"
)

//go:generate mockery --name=ILogService
type ILogService interface {
	AddLog(ctx context.Context, req *dto.AddLogReq) (*dto.Log, error)
	GetLogs(c context.Context, req *dto.GetLogsReq) (dto.Logs, *paging.Pagination, error)
}

type LogService struct {
	validator validation.Validation
	repo      repository.ILogRepository
}

func NewLogService(
	validator validation.Validation,
	repo repository.ILogRepository,
) *LogService {
	return &LogService{
		validator: validator,
		repo:      repo,
	}
}

func (s *LogService) AddLog(ctx context.Context, req *dto.AddLogReq) (*dto.Log, error) {
	if err := s.validator.ValidateStruct(req); err != nil {
		return nil, err
	}

	log, err := s.repo.AddLog(ctx, req)
	if err != nil {
		return nil, err
	}

	return log, nil
}

func (s *LogService) GetLogs(ctx context.Context, req *dto.GetLogsReq) (dto.Logs, *paging.Pagination, error) {
	if err := s.validator.ValidateStruct(req); err != nil {
		return nil, nil, err
	}

	logs, pagination, err := s.repo.GetLogs(ctx, req)
	if err != nil {
		return nil, nil, err
	}

	return logs, pagination, nil
}
