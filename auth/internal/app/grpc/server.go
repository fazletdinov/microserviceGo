package grpcapp

import (
	"auth/config"
	"auth/internal/api/grpc/auth"
	"fmt"
	"log/slog"
	"net"

	authgrpc "auth/protogen/auth"

	"auth/internal/domain/repository"
	"auth/internal/domain/service/grpc_service"

	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type GRPC struct {
	log        *slog.Logger
	gRPCServer *grpc.Server
	Env        *config.Config
}

func NewGRPC(log *slog.Logger, env *config.Config, db *gorm.DB) *GRPC {
	gRPCServer := grpc.NewServer()
	authgrpc.RegisterGatewayAuthServer(
		gRPCServer,
		&auth.GRPCController{
			UserGRPCService: grpcservice.NewUserGRPCService(repository.NewUserGRPCRepository(db)),
			Env:             env,
		},
	)

	return &GRPC{
		log:        log,
		gRPCServer: gRPCServer,
		Env:        env,
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
		slog.Int("port", g.Env.GRPC.Port),
	)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", g.Env.GRPC.Port))
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
	g.log.With(slog.String("op", op)).Info("Остановка gRPC Server", slog.Int("port", g.Env.GRPC.Port))
	g.gRPCServer.GracefulStop()
}
