package main

import (
	"context"
	"os/signal"
	"syscall"

	"github.com/SussyaPusya/swiftTalk/chat-service/config"
	"github.com/SussyaPusya/swiftTalk/chat-service/internal/pkg"
	"github.com/SussyaPusya/swiftTalk/chat-service/internal/transport"
)

func main() {
	ctx := context.Background()

	logger, err := pkg.NewDevelopmentLogger()
	if err != nil {
		panic(err)
	}

	config, err := config.New()
	if err != nil {
		panic(err)
	}

	grpcClient := transport.NewClient(&config.GRPC)

	middleware := transport.NewMiddleware(grpcClient.Client)

	server := transport.NewRouter(middleware)

	go server.Run()

	ctx, stop := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)

	defer stop()

	<-ctx.Done()

	grpcClient.Close()

	logger.Info("Service shutdown...")
}
