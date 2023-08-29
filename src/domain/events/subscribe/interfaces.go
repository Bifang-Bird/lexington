package subscribe

import (
	"context"
	event "lexington/src/domain/events"
)

type (
	MqEventHandler interface {
		DemoHandle(context.Context, *event.MqMessageBodyDemo) error
	}
)
