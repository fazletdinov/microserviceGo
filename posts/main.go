package main

import (
	"context"
	"os"
	"os/signal"
	"posts/internal/app"
	"posts/pkg/metric"
	"posts/pkg/tracer"
	"syscall"
)

func main() {

	app := app.App()
	env := app.Env
	go app.GRPCServer.MustRun()
	ctx := context.Background()
	{
		tp, err := tracer.InitTracer(
			ctx,
			env.Jaeger.CollectorUrl,
			env.Jaeger.Application,
		)
		if err != nil {
			panic(err)
		}
		defer tp.Shutdown(ctx)
		mp, err := metric.SetupMetrics(
			ctx,
			env.Jaeger.Application,
		)
		if err != nil {
			panic(err)
		}
		defer mp.Shutdown(ctx)
	}
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	<-stop
	app.GRPCServer.Stop()

}
