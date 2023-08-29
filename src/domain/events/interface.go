package event

import (
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
)

type MQInterface interface {
	Worker(ctx context.Context, messages <-chan amqp.Delivery)
}
