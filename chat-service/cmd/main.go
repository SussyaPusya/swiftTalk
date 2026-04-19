package main

import (
	"context"
	"os/signal"
	"syscall"

	"github.com/SussyaPusya/swiftTalk/chat-service/internal/pkg"
)

func main() {
	ctx := context.Background()

	logger, err := pkg.NewDevelopmentLogger()
	if err != nil {
		panic(err)
	}

	ctx, stop := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)

	defer stop()

	<-ctx.Done()

	logger.Info("Service shutdown...")
}
