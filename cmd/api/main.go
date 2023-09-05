package main

import (
	grpcServer "github.com/quangdangfit/log-service/internal/server/grpc"
	"github.com/quangdangfit/log-service/pkg/config"
	"github.com/quangdangfit/log-service/pkg/logger"
	"github.com/quangdangfit/log-service/pkg/mongodb"
	"github.com/quangdangfit/log-service/pkg/validation"
)

func main() {
	cfg := config.GetConfig()
	logger.Initialize(cfg.Environment)
	db, err := mongodb.New(&mongodb.Config{
		URL:      cfg.DatabaseURI,
		Database: cfg.DatabaseName,
	})
	if err != nil {
		logger.Fatal("Cannot connect to database", err)
	}

	validator := validation.New()

	grpcSvr := grpcServer.NewServer(validator, db)
	if err = grpcSvr.Run(); err != nil {
		logger.Fatal(err)
	}
}
