package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/DaniilZ77/authorization/internal/app"
	"github.com/DaniilZ77/authorization/internal/config"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		panic("failed to load config: " + err.Error())
	}

	application := app.New(cfg)

	go func() {
		application.GRPCServer.Run()
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-stop

	application.GRPCServer.Stop()
}
