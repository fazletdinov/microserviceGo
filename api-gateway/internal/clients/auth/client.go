package auth

import (
	authgrpc "api-grpc-gateway/protogen/golang/auth"
	"context"

	"api-grpc-gateway/config"
	"github.com/google/uuid"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	oteltrace "go.opentelemetry.io/otel/trace"

	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GRPCClientAuth struct {
	auth authgrpc.GatewayAuthClient
	env  *config.Config
}

func NewGRPCClientAuth(
	addrs string,
	env *config.Config,
) (*GRPCClientAuth, error) {
	cc, err := grpc.NewClient(
		addrs,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithStatsHandler(otelgrpc.NewClientHandler()),
	)
	if err != nil {
		return nil, err
	}
	return &GRPCClientAuth{
		auth: authgrpc.NewGatewayAuthClient(cc),
		env:  env,
	}, nil
}

func (gc *GRPCClientAuth) CreateUser(
	ctx context.Context,
	email string,
	password string,
) (uuid.UUID, error) {
	var tracer = otel.Tracer(gc.env.Jaeger.ServerName)
	traceCtx, span := tracer.Start(
		ctx,
		"CreateUser",
		oteltrace.WithAttributes(attribute.String("Email", email)),
	)
	// traceId := fmt.Sprintf("%s", span.SpanContext().TraceID())
	// traceCtx = metadata.AppendToOutgoingContext(ctx, "x-trace-id", traceId)
	span.AddEvent("Начало gRPC запроса в сервис auth для создания пользователя")
	defer span.End()
	response, err := gc.auth.CreateUser(traceCtx, &authgrpc.CreateUserRequest{
		Email:    email,
		Password: password,
	})

	if err != nil {
		return uuid.Nil, err
	}

	return uuid.MustParse(response.UserId), nil
}

func (gc *GRPCClientAuth) GetUserByID(
	ctx context.Context,
	userID uuid.UUID,
) (*authgrpc.GetUserResponse, error) {
	var tracer = otel.Tracer(gc.env.Jaeger.ServerName)
	traceCtx, span := tracer.Start(
		ctx,
		"GetUserByID",
		oteltrace.WithAttributes(attribute.String("UserID", userID.String())),
	)
	span.AddEvent("Начало gRPC запроса в сервис auth для получения пользователя по ID")
	defer span.End()
	response, err := gc.auth.GetUserByID(traceCtx, &authgrpc.GetUserRequest{
		UserId: userID.String(),
	})

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (gc *GRPCClientAuth) UpdateUser(
	ctx context.Context,
	userID uuid.UUID,
	firstName string,
	lastName string,
) (string, error) {
	var tracer = otel.Tracer(gc.env.Jaeger.ServerName)
	traceCtx, span := tracer.Start(
		ctx,
		"UpdateUser",
		oteltrace.WithAttributes(attribute.String("FirstName", firstName)),
		oteltrace.WithAttributes(attribute.String("LastNameName", lastName)),
	)
	span.AddEvent("Начало gRPC запроса в сервис auth для обновления пользователя")
	defer span.End()
	response, err := gc.auth.UpdateUser(traceCtx, &authgrpc.UpdateUserRequest{
		UserId:    userID.String(),
		FirstName: firstName,
		LastName:  lastName,
	})

	if err != nil {
		return "", err
	}

	return response.SuccessMessage, nil
}

func (gc *GRPCClientAuth) DeleteUser(
	ctx context.Context,
	userID uuid.UUID,
) (string, error) {
	var tracer = otel.Tracer(gc.env.Jaeger.ServerName)
	traceCtx, span := tracer.Start(
		ctx,
		"DeleteUser",
		oteltrace.WithAttributes(attribute.String("UserID", userID.String())),
	)
	span.AddEvent("Начало gRPC запроса в сервис auth для удаления пользователя по ID")
	defer span.End()
	response, err := gc.auth.DeleteUser(traceCtx, &authgrpc.DeleteUserRequest{
		UserId: userID.String(),
	})

	if err != nil {
		return "", err
	}

	return response.SuccessMessage, nil
}

func (gc *GRPCClientAuth) GetUserByEmail(
	ctx context.Context,
	email string,
) (*authgrpc.GetUserResponse, error) {
	var tracer = otel.Tracer(gc.env.Jaeger.ServerName)
	traceCtx, span := tracer.Start(
		ctx,
		"GetUserByEmail",
		oteltrace.WithAttributes(attribute.String("Email", email)),
	)
	span.AddEvent("Начало gRPC запроса в сервис auth для получения пользователя по Email")
	defer span.End()
	response, err := gc.auth.GetUserByEmail(traceCtx, &authgrpc.GetUserByEmailRequest{
		Email: email,
	})

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (gc *GRPCClientAuth) GetUserByEmailIsActive(
	ctx context.Context,
	email string,
) (*authgrpc.GetUserResponse, error) {
	var tracer = otel.Tracer(gc.env.Jaeger.ServerName)
	traceCtx, span := tracer.Start(
		ctx,
		"GetUserByEmailIsActive",
		oteltrace.WithAttributes(attribute.String("Email", email)),
	)
	span.AddEvent("Начало gRPC запроса в сервис auth для получения активного пользователя по Email")
	defer span.End()

	response, err := gc.auth.GetUserByEmailIsActive(traceCtx, &authgrpc.GetUserByEmailRequest{
		Email: email,
	})
	if err != nil {
		return nil, err
	}

	return response, nil
}
