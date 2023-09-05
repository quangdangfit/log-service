package model

import (
	"context"
	"errors"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Log struct {
	Id          primitive.ObjectID `bson:"_id" json:"id"`
	ServiceName string             `bson:"service_name" json:"service_name"`
	Level       string             `bson:"level" json:"level"`
	Content     []byte             `bson:"content" json:"content"`
	TraceId     string             `bson:"trace_id" json:"trace_id"`
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
}

func (l *Log) BeforeCreate(_ context.Context) error {
	if l.TraceId == "" {
		return errors.New("required log trace_id")
	}

	if len(l.Content) == 0 {
		return errors.New("required log content")
	}

	if l.Level == "" {
		return errors.New("required log level")
	}

	l.Id = primitive.NewObjectID()
	l.CreatedAt = time.Now().UTC()
	l.Level = strings.ToUpper(l.Level)
	return nil
}
