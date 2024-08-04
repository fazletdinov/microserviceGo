package grpcapp

import (
	"fmt"
	"log/slog"
	"net"
	"posts/config"
	"posts/internal/api/grpc/posts"

	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	postsgrpc "posts/protogen/posts"

	repositoryComment "posts/internal/domain/repository/comment"
	repositoryPost "posts/internal/domain/repository/post"
	commentService "posts/internal/domain/service/comment"
	postService "posts/internal/domain/service/post"

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
	postsgrpc.RegisterGatewayPostsServer(
		gRPCServer,
		&posts.PostsController{
			PostService:    postService.NewPostService(repositoryPost.NewPostRepository(db)),
			CommentService: commentService.NewCommentService(repositoryComment.NewCommentRepository(db)),
			Env:            env,
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
		slog.Int("port", g.Env.GRPC.PostsGRPCPort),
	)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", g.Env.GRPC.PostsGRPCPort))
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
	g.log.With(slog.String("op", op)).Info("Остановка gRPC Server", slog.Int("port", g.Env.GRPC.PostsGRPCPort))
	g.gRPCServer.GracefulStop()
}
