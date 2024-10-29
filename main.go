package main

import (
	"context"
	"os/signal"
	"syscall"
	"treads/cmd"
	"treads/infra"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)
	defer stop()

	loadingEnv := infra.NewConfig()
	container := infra.NewContainerDI(loadingEnv)

	cmd.StartAPI(ctx, container)

}
