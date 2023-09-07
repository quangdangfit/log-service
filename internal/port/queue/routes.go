package queue

import (
	"context"
	"log"

	"github.com/streadway/amqp"

	"github.com/quangdangfit/log-service/internal/repository"
	"github.com/quangdangfit/log-service/internal/service"
	"github.com/quangdangfit/log-service/pkg/config"
	"github.com/quangdangfit/log-service/pkg/logger"
	"github.com/quangdangfit/log-service/pkg/mongodb"
	"github.com/quangdangfit/log-service/pkg/rabbitmq"
	"github.com/quangdangfit/log-service/pkg/validation"
)

func ConsumeMessages(db mongodb.DB, validator validation.Validation, mq rabbitmq.MQ) {
	logRepo := repository.NewLogRepository(db)
	logSvc := service.NewLogService(validator, logRepo)
	logHandler := NewLogHandler(logSvc)

	msgC, errC := mq.Consume()
	go func() {
		for {
			select {
			case msg := <-msgC:
				routes(logHandler, msg)
			case err := <-errC:
				log.Println("err", err.Error())
			}
		}
	}()

}

func routes(logHandler *LogHandler, msg amqp.Delivery) {
	switch msg.RoutingKey {
	case config.RoutingKeyAddLog:
		_ = logHandler.AddLog(context.Background(), msg.Body)
	default:
		logger.Error("Unknown routing key %s", msg.RoutingKey)
	}
}
