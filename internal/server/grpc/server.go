package grpc

import (
	"fmt"
	"net"

	"google.golang.org/grpc"

	logGRPC "github.com/quangdangfit/log-service/internal/port/gprc"
	"github.com/quangdangfit/log-service/pkg/config"
	"github.com/quangdangfit/log-service/pkg/logger"
	"github.com/quangdangfit/log-service/pkg/middleware"
	"github.com/quangdangfit/log-service/pkg/mongodb"
	"github.com/quangdangfit/log-service/pkg/validation"
)

type Server struct {
	engine    *grpc.Server
	cfg       *config.Config
	validator validation.Validation
	db        mongodb.DB
}

func NewServer(validator validation.Validation, db mongodb.DB) *Server {
	cfg := config.GetConfig()

	interceptor := middleware.NewInternalAuthInterceptor(cfg.SecretAPIKey, config.AuthIgnoreMethods)

	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			interceptor.Unary(),
		),
	)

	return &Server{
		engine:    grpcServer,
		cfg:       config.GetConfig(),
		validator: validator,
		db:        db,
	}
}

func (s Server) Run() error {
	logGRPC.RegisterHandlers(s.engine, s.db, s.validator)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.cfg.Port))
	logger.Info("gRPC server is listening on PORT: ", s.cfg.Port)
	if err != nil {
		logger.Error("Failed to listen: ", err)
		return err
	}

	// Start grpc server
	err = s.engine.Serve(lis)
	if err != nil {
		logger.Fatal("Failed to serve grpc: ", err)
		return err
	}

	return nil
}
