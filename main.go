package main

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"

	companiesController "github.com/any/companies/internal/api/companies"
	jwtController "github.com/any/companies/internal/api/jwt"
	"github.com/any/companies/internal/config"
	"github.com/any/companies/internal/dataBus"
	"github.com/any/companies/internal/infr/kafka"
	"github.com/any/companies/internal/infr/logger"
	"github.com/any/companies/internal/infr/server"
	"github.com/any/companies/internal/repositories/postgres"
	"github.com/any/companies/internal/services/companyService"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg, err := config.InitConfig("APP")
	if err != nil {
		panic(errors.Wrap(err, "init config error"))
	}

	l, err := logger.InitLogger(cfg.Logger)
	if err != nil {
		panic(errors.Wrap(err, "init logger error"))
	}

	postgresRepository, err := postgres.NewPostgresRepository(cfg.Postgres)
	if err != nil {
		l.Fatal("init database error", zap.Error(err))
	}

	k, err := kafka.InitWriter(cfg.KafkaWriter)
	if err != nil {
		panic(errors.Wrap(err, "init kafka error"))
	}

	companiesService := companyService.NewService(
		postgresRepository,
		dataBus.NewKafkaEventsPublisher(k, l),
	)

	s, err := server.New(
		cfg.Server,
		l,
		jwtController.NewController(l),
		companiesController.NewController(l, companiesService),
	)
	if err != nil {
		l.Fatal("init server error", zap.Error(err))
	}

	serverErrChan := s.Listen()
	exitChan := StartListenForQuit(ctx, l)
	select {
	case err := <-serverErrChan:
		l.Error("received error from server", zap.Error(err))

		break
	case <-exitChan:
		break
	}

	l.Info("start shutdown server")
	if err := s.Shutdown(ctx); err != nil {
		l.Error("error while shutdown server", zap.Error(err))
	}
}

func StartListenForQuit(ctx context.Context, l logger.Logger) <-chan struct{} {
	exitChan := make(chan struct{})
	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)
		select {
		case <-ctx.Done():
			return
		case sig := <-quit:
			l.Info(fmt.Sprintf("OS signal received: %v", sig.String()))
			close(exitChan)

			return
		}
	}()

	return exitChan
}
