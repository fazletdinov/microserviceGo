package main

import (
	"auth/internal/app"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	app := app.NewApp()
	go app.GRPCServer.MustRun()
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	<-stop
	app.GRPCServer.Stop()
}
