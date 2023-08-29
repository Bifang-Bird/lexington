package mq

import (
	"context"
	"encoding/json"
	event "lexington/src/domain/events"
	"lexington/src/domain/events/subscribe"

	"github.com/google/wire"
	amqp "github.com/rabbitmq/amqp091-go"
	"golang.org/x/exp/slog"
)

var _ event.MQInterface = (*mqServer)(nil)

var MQServerSet = wire.NewSet(NewMQServer)

type mqServer struct {
	orderHandler subscribe.MqEventHandler
}

func NewMQServer(orderHandler subscribe.MqEventHandler) event.MQInterface {
	return &mqServer{orderHandler: orderHandler}
}

func (a *mqServer) Worker(ctx context.Context, messages <-chan amqp.Delivery) {
	for delivery := range messages {
		slog.Info("processDeliveries", "delivery_tag", delivery.DeliveryTag, "delivery_type", delivery.Type)

		switch delivery.Type {
		case "sms-timeout":
			var payload event.MqMessageBodyDemo

			err := json.Unmarshal(delivery.Body, &payload)
			if err != nil {
				slog.Error("failed to Unmarshal message", err)
			}
			err = a.orderHandler.DemoHandle(ctx, &payload)

			if err != nil {
				if err = delivery.Reject(false); err != nil {
					slog.Error("failed to delivery.Reject", err)
				}

				slog.Error("failed to process delivery", err)
			} else {
				err = delivery.Ack(false)
				if err != nil {
					slog.Error("failed to acknowledge delivery", err)
				}
			}
		default:
			slog.Info("default")
		}
	}
	slog.Info("Deliveries channel closed")
}
