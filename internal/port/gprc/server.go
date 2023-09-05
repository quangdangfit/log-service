package grpc

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/quangdangfit/log-service/internal/repository"
	"github.com/quangdangfit/log-service/internal/service"
	"github.com/quangdangfit/log-service/pkg/mongodb"
	"github.com/quangdangfit/log-service/pkg/validation"
	pb "github.com/quangdangfit/log-service/proto/gen/go/log"
)

func RegisterHandlers(svr *grpc.Server, db mongodb.DB, validator validation.Validation) {
	logRepo := repository.NewLogRepository(db)
	logSvc := service.NewLogService(validator, logRepo)
	logHandler := NewLogHandler(logSvc)

	pb.RegisterLogServiceServer(svr, logHandler)
	reflection.Register(svr)
}
