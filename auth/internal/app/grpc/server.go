package grpcapp

import (
	"auth/internal/api/grpc/auth"
	"fmt"
	"log/slog"
	"net"

	"auth/internal/domain/service"
	"google.golang.org/grpc"
)

type GRPC struct {
	log        *slog.Logger
	gRPCServer *grpc.Server
	port       int
}

func NewGRPC(log *slog.Logger, port int, pathSecret string) *GRPC {
	gRPCServer := grpc.NewServer()
	auth.Register(gRPCServer, service.NewGRPCService(), pathSecret)

	return &GRPC{
		log:        log,
		gRPCServer: gRPCServer,
		port:       port,
	}
}

func (g *GRPC) MustRun() {
	if err := g.Run(); err != nil {
		panic(err)
	}
}

func (g *GRPC) Run() error {
	const op = "grpcapp.Run"
	log := g.log.With(slog.String("op", op),
		slog.Int("port", g.port),
	)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", g.port))
	if err != nil {
		return fmt.Errorf("%s:%v", op, err)
	}
	log.Info("gRPC Server запущен", slog.String("addr", lis.Addr().String()))

	if err = g.gRPCServer.Serve(lis); err != nil {
		return fmt.Errorf("%s:%v", op, err)
	}
	return nil
}

func (g *GRPC) Stop() {
	const op = "grpcapp.Run"
	g.log.With(slog.String("op", op)).Info("Остановка gRPC Server", slog.Int("port", g.port))
	g.gRPCServer.GracefulStop()
}
