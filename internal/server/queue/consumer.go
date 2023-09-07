package queue

import (
	logQueue "github.com/quangdangfit/log-service/internal/port/queue"
	"github.com/quangdangfit/log-service/pkg/config"
	"github.com/quangdangfit/log-service/pkg/mongodb"
	"github.com/quangdangfit/log-service/pkg/rabbitmq"
	"github.com/quangdangfit/log-service/pkg/validation"
)

type Consumer struct {
	mq        rabbitmq.MQ
	cfg       *config.Config
	validator validation.Validation
	db        mongodb.DB
}

func NewConsumer(validator validation.Validation, db mongodb.DB, rabbitMQ rabbitmq.MQ) *Consumer {
	return &Consumer{
		mq:        rabbitMQ,
		cfg:       config.GetConfig(),
		validator: validator,
		db:        db,
	}
}

func (c Consumer) Run() {
	logQueue.ConsumeMessages(c.db, c.validator, c.mq)
}
