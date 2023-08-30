package main

import (
	"codeup.aliyun.com/6145b2b428003bdc3daa97c8/go-simba/go-simba-pkg.git/rabbitmq/consumer"
	"context"
	"fmt"
	"github.com/Bifang-Bird/simbapkg/pkg/dbconfig"
	"lexington/src/app/config"
	"lexington/src/app/inject"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/Bifang-Bird/simbapkg/app"
	"go.uber.org/zap"

	"go.uber.org/automaxprocs/maxprocs"
	"golang.org/x/exp/slog"
	"google.golang.org/grpc"
)

func main() {
	_, err := maxprocs.Set()
	if err != nil {
		slog.Error("failed set max procs", err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	cfg, err := config.NewConfig()
	if err != nil {
		slog.Error("failed get config", err)
	}

	s := app.NewServer()
	//日志初始化
	s.SetInitLogHandler(app.InitLogger).InitLogHandler(&cfg.Log)
	app.Logger.Info("⚡ 初始化应用开始", zap.String("service", cfg.Name), zap.String("version", cfg.App.Version))
	//初始化grpc
	server := s.SetInitGrpcHandler(app.InitGrpcServer).InitGrpcHandler(ctx)
	go func() {
		defer server.GracefulStop()
		<-ctx.Done()
	}()

	cleanup := prepareApp(ctx, cancel, cfg, server, 0)

	//绑定端口
	l := s.SetBandingPortHandler(app.BandingPort).BandingPortHandler(&cfg.HTTP, cancel)
	defer func() {
		if err1 := l.Close(); err != nil {
			slog.Error("failed to close", err1, "network", "tcp", "address", fmt.Sprintf("%s:%d", cfg.HTTP.Host, cfg.HTTP.Port))
		}
	}()

	err = server.Serve(l)
	if err != nil {
		slog.Error("failed start gRPC server", err, "network", "tcp", "address", fmt.Sprintf("%s:%d", cfg.HTTP.Host, cfg.HTTP.Port))
		cancel()
	}
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	select {
	case v := <-quit:
		cleanup()
		app.Logger.Info("系统退出", zap.Any("signal.Notify", v))
	case done := <-ctx.Done():
		cleanup()
		app.Logger.Info("系统退出", zap.Any("ctx.Done", done))
	}

}

func prepareApp(ctx context.Context, cancel context.CancelFunc, cfg *config.Config, server *grpc.Server, mqCustomerNum int) func() {
	serve, cleanup, err := inject.InitApp(cfg, server, mqCustomerNum)
	if err != nil {
		slog.Error("failed init app", err)
		cancel()
		<-ctx.Done()
	}
	//初始化消费者
	InitMqConsumer(ctx, cancel, cfg, serve, mqCustomerNum)
	return cleanup
}

func InitMqConsumer(ctx context.Context, cancel context.CancelFunc, cfg *config.Config, serve *inject.App, num int) {
	if len(cfg.RabbitMQ.Consumer) > 0 && len(cfg.RabbitMQ.Consumer) == num {
		for i, consu := range cfg.RabbitMQ.Consumer {
			finishOrderConsumer := serve.Consumer[i]
			finishOrderConsumer.Configure(
				consumer.ExchangeName(consu.Body.ExchangeName),
				consumer.BindingKey(consu.Body.BindingKey),
				consumer.QueueName(consu.Body.QueueName),
				consumer.ConsumerTag(consu.Body.ConsumerTag),
			)
			go func(worker consumer.EventConsumer, num int, conp dbconfig.ConsumerProfile) {
				app.Logger.Info("启动消费者", zap.Any(fmt.Sprintf("第[%s]个", strconv.Itoa(num)), conp.Type), zap.String("交换机", conp.Body.ExchangeName))
				err1 := worker.StartConsumer(serve.MqServer.Worker)
				if err1 != nil {
					slog.Error("failed to start Consumer", err1)
					cancel()
					<-ctx.Done()
				}
			}(finishOrderConsumer, i, consu)
		}
	} else {

	}
}
