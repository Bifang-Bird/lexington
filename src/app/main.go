package main

import (
	"context"
	"fmt"
	"merchant/src/app/config"
	"merchant/src/app/inject"
	"os"
	"os/signal"
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
	app.Logger.Info("⚡ init app", zap.String("service", cfg.Name), zap.String("version", cfg.App.Version))
	//初始化grpc
	server := s.SetInitGrpcHandler(app.InitGrpcServer).InitGrpcHandler(ctx)
	go func() {
		defer server.GracefulStop()
		<-ctx.Done()
	}()

	cleanup := prepareApp(ctx, cancel, cfg, server)
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
		slog.Info("signal.Notify", v)
	case done := <-ctx.Done():
		cleanup()
		slog.Info("ctx.Done", done)
	}
}

func prepareApp(ctx context.Context, cancel context.CancelFunc, cfg *config.Config, server *grpc.Server) func() {
	_, cleanup, err := inject.InitApp(cfg, &cfg.DataSource)
	if err != nil {
		slog.Error("failed init app", err)
		cancel()
		<-ctx.Done()
	}
	return cleanup
}
