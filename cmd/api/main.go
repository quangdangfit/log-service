package main

import (
	grpcServer "github.com/quangdangfit/log-service/internal/server/grpc"
	"github.com/quangdangfit/log-service/internal/server/queue"
	"github.com/quangdangfit/log-service/pkg/config"
	"github.com/quangdangfit/log-service/pkg/logger"
	"github.com/quangdangfit/log-service/pkg/mongodb"
	"github.com/quangdangfit/log-service/pkg/rabbitmq"
	"github.com/quangdangfit/log-service/pkg/validation"
)

func main() {
	cfg := config.LoadConfig()
	logger.Initialize(cfg.Environment)
	validator := validation.New()

	db, err := mongodb.New(&mongodb.Config{
		URL:      cfg.DatabaseURI,
		Database: cfg.DatabaseName,
	})
	if err != nil {
		logger.Fatal("Cannot connect to database", err)
	}

	rabbitMQ, err := rabbitmq.New(cfg.RabbitMQURL, config.QueueName)
	if err != nil {
		logger.Fatalf("Could not initialize rabbitmq: %v", err)
	}

	consumer := queue.NewConsumer(validator, db, rabbitMQ)
	go consumer.Run()

	grpcSvr := grpcServer.NewServer(validator, db)
	if err = grpcSvr.Run(); err != nil {
		logger.Fatal(err)
	}
}
