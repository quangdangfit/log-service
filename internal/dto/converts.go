package dto

import (
	"time"

	pb "github.com/quangdangfit/log-service/proto/gen/go/log"
)

type Logs []*Log

func (l *Log) ToProto() *pb.LogInfo {
	return &pb.LogInfo{
		Id:          l.Id,
		ServiceName: l.ServiceName,
		Level:       l.Level,
		Content:     l.Content,
		TraceId:     l.TraceId,
		CreatedAt:   l.CreatedAt.Format(time.RFC3339),
	}
}

func (lgs Logs) ToProto() []*pb.LogInfo {
	rs := make([]*pb.LogInfo, len(lgs))
	for i, log := range lgs {
		if log != nil {
			rs[i] = log.ToProto()
		}
	}
	return rs
}
