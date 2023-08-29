package subscribe

import (
	"context"
	"github.com/google/wire"
	"golang.org/x/exp/slog"
	event "lexington/src/domain/events"
	UC "lexington/src/usecases"
)

type MqEventHandlerImpl struct {
	uc UC.UseCase
}

var _ MqEventHandler = (*MqEventHandlerImpl)(nil)

var MqEventHandlerSet = wire.NewSet(NewMqEventHandlerImpl)

func NewMqEventHandlerImpl(
	uc UC.UseCase,
) MqEventHandler {
	return &MqEventHandlerImpl{
		uc: uc,
	}
}

func (h *MqEventHandlerImpl) DemoHandle(ctx context.Context, e *event.MqMessageBodyDemo) error {
	slog.Info("handle", "delivery_message", e)

	return nil
}
