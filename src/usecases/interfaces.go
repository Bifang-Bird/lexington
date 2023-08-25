package usecases

import (
	"context"
	"lexington/src/app/config"
	"codeup.aliyun.com/6145b2b428003bdc3daa97c8/go-simba/go-simba-pkg.git/rabbitmq/publisher"
)

type (
	PaymentPublisher interface {
		Configure(...publisher.Option)
		DelayConfigure(...publisher.DeplayOption)
		Publish(context.Context, []byte, string, config.MqConfig) error
		DelayPublish(context.Context, []byte, string, int64) error
	}

	// DomainService 领域服务
	DomainService interface {
		//todo
	}

	// UseCase 用例
	UseCase interface {
		//todo
	}
)
