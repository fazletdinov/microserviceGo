package grpcapp

import (
	"fmt"
	"likes/config"
	"likes/internal/api/grpc/likes"
	"log/slog"
	"net"

	likesgrpc "likes/protogen/likes"

	reactionRepository "likes/internal/domain/repository"
	reactionService "likes/internal/domain/service"

	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type GRPC struct {
	log        *slog.Logger
	gRPCServer *grpc.Server
	Env        *config.Config
}

func NewGRPC(log *slog.Logger, env *config.Config, db *gorm.DB) *GRPC {
	gRPCServer := grpc.NewServer(grpc.StatsHandler(otelgrpc.NewServerHandler()))
	likesgrpc.RegisterGatewayLikesServer(
		gRPCServer,
		&likes.ReactionController{
			ReactionService: reactionService.NewReactionService(reactionRepository.NewReactionRepository(db)),
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
		slog.Int("port", g.Env.GRPC.LikesGRPCPort),
	)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", g.Env.GRPC.LikesGRPCPort))
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
	g.log.With(slog.String("op", op)).Info("Остановка gRPC Server", slog.Int("port", g.Env.GRPC.LikesGRPCPort))
	g.gRPCServer.GracefulStop()
}
