package main

import (
	"context"
	"fmt"
	jwtController "github.com/any/companies/internal/api/jwt"
	"github.com/any/companies/internal/config"
	"github.com/any/companies/internal/infr/logger"
	"github.com/any/companies/internal/infr/server"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
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

	cont := jwtController.NewController(l)

	s, err := server.New(cfg.Server, l, cont)
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
