package inject

import (
	"codeup.aliyun.com/6145b2b428003bdc3daa97c8/go-simba/go-simba-pkg.git/rabbitmq"
	pkgConsumer "codeup.aliyun.com/6145b2b428003bdc3daa97c8/go-simba/go-simba-pkg.git/rabbitmq/consumer"
	"codeup.aliyun.com/6145b2b428003bdc3daa97c8/go-simba/go-simba-proto.git/gen"
	"github.com/Bifang-Bird/simbapkg/pkg"
	"lexington/src/app/config"
	"lexington/src/domain"
	"lexington/src/domain/events"
	UC "lexington/src/usecases"
)

type App struct {
	Cfg                   *config.Config
	DB                    pkg.DB
	AMPPool               *rabbitmq.ChannelPool
	Consumer              []pkgConsumer.EventConsumer
	PaymentPub            pkg.MqPublisher
	DomainRepo            domain.DomainRepo
	PayOrderDomainService UC.DomainService
	UC                    UC.UseCase
	SmsGRPCServer         gen.SmsServiceServer
	MqServer              event.MQInterface
}

func New(
	cfg *config.Config,
	db pkg.DB,
	amqppool *rabbitmq.ChannelPool,
	consumer []pkgConsumer.EventConsumer,
	paymentPub pkg.MqPublisher,
	domainRepo domain.DomainRepo,
	domainService UC.DomainService,
	uc UC.UseCase,
	paymentGRPCServer gen.SmsServiceServer,
	mqServer event.MQInterface,
) *App {
	return &App{
		Cfg:                   cfg,
		DB:                    db,
		DomainRepo:            domainRepo,
		PayOrderDomainService: domainService,
		UC:                    uc,
		SmsGRPCServer:         paymentGRPCServer,
		AMPPool:               amqppool,
		Consumer:              consumer,
		PaymentPub:            paymentPub,
		MqServer:              mqServer,
	}
}
