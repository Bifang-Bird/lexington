package event

import (
	shared "github.com/Bifang-Bird/simbapkg/pkg/shared_kernel"
)

// MqMessageBodyDemo
// @Description: 接受消息结构体
type MqMessageBodyDemo struct {
	shared.DomainEvent
}

func (e MqMessageBodyDemo) Identity() string {
	return "MqMessageBodyDemo"
}
