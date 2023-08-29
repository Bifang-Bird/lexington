//go:build wireinject
// +build wireinject

package inject

import (
	"codeup.aliyun.com/6145b2b428003bdc3daa97c8/go-simba/go-simba-pkg.git/rabbitmq"
	pkgConsumer "codeup.aliyun.com/6145b2b428003bdc3daa97c8/go-simba/go-simba-pkg.git/rabbitmq/consumer"
	pkgPublisher "codeup.aliyun.com/6145b2b428003bdc3daa97c8/go-simba/go-simba-pkg.git/rabbitmq/publisher"
	pkg "github.com/Bifang-Bird/simbapkg/pkg"
	"github.com/Bifang-Bird/simbapkg/pkg/dbFactory"

	"github.com/Bifang-Bird/simbapkg/pkg/mq/publish"
	"lexington/src/app/config"
	"lexington/src/domain/events/subscribe"
	"lexington/src/domain/service"
	grpcserver "lexington/src/infras_adapter/grpc"
	"lexington/src/infras_adapter/mq"
	"lexington/src/infras_adapter/repo"
	uc "lexington/src/usecases"

	"github.com/google/wire"
	"google.golang.org/grpc"
)

func InitApp(
	cfg *config.Config,
	grpcServer *grpc.Server,
	num int,
) (*App, func(), error) {
	panic(wire.Build(
		New,
		dbEngineFunc,
		rabbitMQFuncPool,
		pkgConsumer.EventConsumersSet,
		pkgPublisher.EventPublisherSet,
		pkgPublisher.EventDeplayPublisherSet,
		publish.PublisherSet,
		repo.DomainRepositorySet,
		service.DomainServiceSet,
		uc.UseCaseSet,
		grpcserver.SmsGRPCServerSet,
		subscribe.MqEventHandlerSet,
		mq.MQServerSet,
	))
}

func dbEngineFunc(cfg *config.Config) (pkg.DB, func(), error) {
	var ds = cfg.DataSource
	db, err := dbFactory.GetDb(ds)
	if err != nil {
		return nil, nil, err
	}
	return db, func() { db.Close() }, nil
}

func rabbitMQFuncPool(cfg *config.Config) (*rabbitmq.ChannelPool, error) {
	url := rabbitmq.RabbitMQConnStr(cfg.RabbitMQ.URL)
	if url == "" {
	}
	channelPool := new(rabbitmq.ChannelPool)
	channelPool.InitPool(url)
	return channelPool, nil
}
