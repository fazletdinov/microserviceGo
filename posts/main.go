package main

import (
	"os"
	"posts/internal/app"
)

func main() {

	app := app.App()
	go app.GRPCServer.MustRun()
	stop := make(chan os.Signal, 1)

	<-stop
	app.GRPCServer.Stop()

}
