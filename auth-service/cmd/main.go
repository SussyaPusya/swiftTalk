package main

import (
	"context"
	"os/signal"
	"syscall"

	"github.com/SussyaPusya/swiftTalk/auth-service/config"
	"github.com/SussyaPusya/swiftTalk/auth-service/internal/pkg"
	"github.com/SussyaPusya/swiftTalk/auth-service/internal/repository"
	"github.com/SussyaPusya/swiftTalk/auth-service/internal/service"
	"github.com/SussyaPusya/swiftTalk/auth-service/internal/transport"
	"go.uber.org/zap"
)

func main() {
	ctx := context.Background()

	logger, err := pkg.NewDevelopmentLogger()
	if err != nil {
		panic(err)
	}

	cfg, err := config.New()
	if err != nil {
		logger.Error("error", zap.Error(err))
		panic(err)
	}

	pg, err := pkg.NewPostgres(ctx, &cfg.Postgres)
	if err != nil {
		logger.Error("error", zap.Error(err))
		panic(err)
	}

	repo := repository.NewRepository(pg)

	srvc := service.NewService(repo, logger.Logger)

	router := transport.NewRouter(srvc)

	go router.Run()
	logger.Info("router started")

	ctx, stop := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)

	defer stop()

	<-ctx.Done()

	logger.Info("Service shutdown...")

}
