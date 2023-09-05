package repository

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/quangdangfit/log-service/internal/dto"
	"github.com/quangdangfit/log-service/internal/model"
	"github.com/quangdangfit/log-service/pkg/config"
	"github.com/quangdangfit/log-service/pkg/mongodb"
	"github.com/quangdangfit/log-service/pkg/paging"
	"github.com/quangdangfit/log-service/pkg/utils"
)

//go:generate mockery --name=ILogRepository
type ILogRepository interface {
	AddLog(ctx context.Context, in *dto.AddLogReq) (*dto.Log, error)
	GetLogs(ctx context.Context, req *dto.GetLogsReq) (dto.Logs, *paging.Pagination, error)
}

type LogRepo struct {
	db mongodb.DB
}

func NewLogRepository(db mongodb.DB) *LogRepo {
	return &LogRepo{db: db}
}

func (r *LogRepo) collectionName() string {
	now := time.Now().UTC()
	year := now.Year()
	month := now.Month()
	return fmt.Sprintf("logs_%d_%d", year, month)
}

func (r *LogRepo) AddLog(ctx context.Context, in *dto.AddLogReq) (*dto.Log, error) {
	ctx, cancel := context.WithTimeout(ctx, config.DatabaseTimeout)
	defer cancel()

	log := &model.Log{
		ServiceName: in.ServiceName,
		Level:       in.Level,
		TraceId:     in.TraceId,
		Content:     in.Content,
	}

	err := log.BeforeCreate(ctx)
	if err != nil {
		return nil, err
	}

	err = r.db.Insert(ctx, r.collectionName(), log)
	if err != nil {
		return nil, err
	}

	rs := &dto.Log{
		Id:          log.Id.Hex(),
		ServiceName: log.ServiceName,
		Level:       log.Level,
		Content:     log.Content,
		TraceId:     log.TraceId,
		CreatedAt:   log.CreatedAt,
	}

	return rs, nil
}

func (r *LogRepo) GetLogs(ctx context.Context, req *dto.GetLogsReq) (dto.Logs, *paging.Pagination, error) {
	ctx, cancel := context.WithTimeout(ctx, config.DatabaseTimeout)
	defer cancel()

	logs := make([]*model.Log, 0)
	filter := r.filter(req)

	total, err := r.db.Count(ctx, r.collectionName(), filter)
	if err != nil {
		return nil, nil, err
	}

	if req.Pagination == nil {
		req.Pagination = &dto.PaginationReq{
			Page:  1,
			Limit: config.DefaultPageLimit,
		}
	}

	pagination := paging.New(req.Pagination.Page, req.Pagination.Limit, total)

	err = r.db.Find(
		ctx,
		r.collectionName(),
		&logs,
		mongodb.WithFilter(filter),
		mongodb.WithPaging(pagination),
		mongodb.WithSorter(r.buildSort(req.Sort)),
	)
	if err != nil {
		return nil, nil, err
	}

	rs := make([]*dto.Log, len(logs))
	utils.Copy(&rs, &logs)
	return rs, pagination, nil
}

func (r *LogRepo) filter(req *dto.GetLogsReq) bson.M {
	filter := bson.M{}
	if req.ServiceName != "" {
		filter["service_name"] = req.ServiceName
	}

	if req.Level != "" {
		filter["level"] = req.Level
	}

	if req.TraceId != "" {
		filter["trace_id"] = req.TraceId
	}

	if req.CreatedTo == 0 {
		req.CreatedTo = time.Now().Unix()
	}

	if req.CreatedFrom == 0 || req.CreatedFrom > req.CreatedTo {
		req.CreatedFrom = time.Now().Add(-24 * time.Hour).Unix()
	}

	filter["created_at"] = bson.M{
		"$gte": time.Unix(req.CreatedFrom, 0),
		"$lte": time.Unix(req.CreatedTo, 0),
	}

	return filter
}

func (r *LogRepo) buildSort(req *dto.SortReq) bson.D {
	sort := bson.D{{"_id", -1}}
	if req != nil {
		if req.Desc {
			sort = bson.D{{req.Field, -1}}
		} else {
			sort = bson.D{{req.Field, 1}}
		}
	}
	return sort
}
