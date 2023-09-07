package rabbitmq

import (
	"time"

	"github.com/streadway/amqp"

	"github.com/quangdangfit/log-service/pkg/logger"
)

type MQ interface {
	Consume() (chan amqp.Delivery, chan error)
}

// WaitTimeReconnect constants
const (
	WaitTimeReconnect = 5
)

// Topic : Topic
type Topic struct {
	Name     string
	Exchange string
}

// Queue : Queue
type Queue struct {
	Name   string
	Topics []Topic
}

// Declaration : Declaration
type Declaration struct {
	Exchanges []string
	Queues    []Queue
}

type mqImpl struct {
	url              string
	connection       *amqp.Connection
	channel          *amqp.Channel
	declaration      *Declaration
	connectionClosed bool
	channelClosed    bool
	chanErr          chan *amqp.Error
}

/**========================================================================
 *                           INTERFACE IMPLEMENTATION
 *========================================================================**/

func New(url string, declarationFile string) (MQ, error) {
	mq := &mqImpl{
		url: url,
	}

	_, err := mq.newConnection()
	if err != nil {
		return nil, err
	}

	return mq, nil
}

func (mq *mqImpl) Consume() (chan amqp.Delivery, chan error) {
	chanMsg := make(chan amqp.Delivery)
	chanErr := make(chan error)
	mq.consuming(chanMsg, chanErr)
	go func() {
		for {
			closedErr := <-mq.chanErr
			if closedErr != nil {
				logger.Errorf("[RabbitMQ] connection is closed, err: %v. Reconnecting...", closedErr)
				err := mq.reconnect()
				if err != nil {
					logger.Errorf("[RabbitMQ] failed to reconnect, err: %v", err)
					continue
				}
				mq.consuming(chanMsg, chanErr)
			}
		}
	}()
	return chanMsg, chanErr
}

/**========================================================================
 *                           PRIVATE METHODS
 *========================================================================**/

func (mq *mqImpl) newConnection() (*amqp.Connection, error) {
	conn, err := amqp.Dial(mq.url)
	for err != nil {
		logger.Errorf(
			"[RabbitMQ] failed to create new connection to AMQP, err: %v. Sleep %d seconds to reconnect",
			err,
			WaitTimeReconnect,
		)
		time.Sleep(WaitTimeReconnect * time.Second)
		conn, err = amqp.Dial(mq.url)
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	mq.channel = ch
	mq.connection = conn
	mq.chanErr = make(chan *amqp.Error)
	mq.connection.NotifyClose(mq.chanErr)

	return conn, nil
}

func (mq *mqImpl) reconnect() (err error) {
	err = mq.closeAll()
	if err != nil {
		logger.Infof("[RabbitMQ] failed to close connection, err: %v", err)
	}

	var conn *amqp.Connection
	for {
		conn, err = amqp.Dial(mq.url)
		if err == nil {
			break
		}

		logger.Infof(
			"[RabbitMQ] failed to create new connection to AMQP: %s. Sleep %d seconds to reconnect.",
			err,
			WaitTimeReconnect,
		)
		time.Sleep(WaitTimeReconnect * time.Second)
	}

	ch, err := conn.Channel()
	if err != nil {
		return err
	}

	mq.channel = ch
	mq.connection = conn
	mq.chanErr = make(chan *amqp.Error)
	conn.NotifyClose(mq.chanErr)

	logger.Info("[RabbitMQ] reconnect rabbitMQ successfully!!!")
	return nil
}

func (mq *mqImpl) consuming(chanMsg chan amqp.Delivery, chanErr chan error) {
	for _, queue := range mq.declaration.Queues {
		go func(qName string) {
			// Consume: queue, consumer, autoAck, exclusive, noLocal, noWait, args
			msgs, err := mq.channel.Consume(qName, "", false, false, false, false, nil)
			if err != nil {
				chanErr <- err
			}

			forever := make(chan bool)
			go func() {
				for d := range msgs {
					chanMsg <- d
				}
			}()
			<-forever
		}(queue.Name)
	}
}

// closeAll : Close connection and channel
func (mq *mqImpl) closeAll() (err error) {
	if !mq.connection.IsClosed() {
		err = mq.connection.Close()
		if err != nil {
			return err
		}
	}
	return nil
}
