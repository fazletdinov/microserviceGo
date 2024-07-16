package main

import (
	"os"
	"os/signal"
	"posts/internal/app"
	"syscall"
)

func main() {

	app := app.App()
	go app.GRPCServer.MustRun()
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	<-stop
	app.GRPCServer.Stop()

}
